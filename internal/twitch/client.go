package twitch

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
)

type chatMsg struct {
	source     string
	command    string
	subcommand string
	message    string
}

func parseMessage(line string) chatMsg {
	parsed := chatMsg{}
	if strings.HasPrefix(line, ":") {
		sline := strings.Split(line, " :")
		// msg
		if len(sline) == 2 {
			parsed.message = sline[1]
		}
		// cmd
		scmd := strings.SplitN(sline[0], " ", 3)
		parsed.source = scmd[0]
		parsed.command = scmd[1]
		parsed.subcommand = scmd[2]

	} else {
		scmd := strings.SplitN(line, " ", 2)
		parsed.command = scmd[0]
		parsed.subcommand = scmd[1]
	}

	return parsed
}

func (conf TwitchConfigs) ReadChat() {
	// init
	conn, err := net.Dial("tcp", conf.TwitchIRL)
	if err != nil {
		log.Panicln("ERROR", err)
	}
	// defer cleanup
	// channel for os signals
	chn := make(chan os.Signal)
	signal.Notify(chn, os.Interrupt)
	// cleanup
	go func(conn net.Conn) {
		<-chn
		log.Println("Clossing connection")
		conn.Close()
	}(conn)

	// join channel
	log.Printf("Joinning Channel %s\n", conf.Channel)
	fmt.Fprintf(conn, "PASS %s\r\nNICK %s\r\nJOIN #%s\r\n", "justinfan6493", "justinfan6493", conf.Channel)

	buffReader := bufio.NewReader(conn)
	for connecting := true; connecting; {
		bytes, _, err := buffReader.ReadLine()
		if err != nil {
			log.Println("ERROR", err)
			connecting = false
		}
		line := string(bytes)
		parsedMsg := parseMessage(line)
		if conf.Debug {
			log.Printf("C:%s %s\n", parsedMsg.command, parsedMsg.message)
		}
		if parsedMsg.command == "366" {
			// End of /Names list
			connecting = false
			if conf.Debug {
				fmt.Println()
			}
		}
	}

	// read block
	log.Println("Echo messages")
	for {
		bytes, _, err := buffReader.ReadLine()
		if err != nil {
			log.Println("ERROR", err)
			break
		}
		line := string(bytes)
		parsedMsg := parseMessage(line)
		switch parsedMsg.command {
		case "PING":
			// respond with PONG
			pong := fmt.Sprintf("PONG %s", parsedMsg.subcommand)
			conn.Write([]byte(pong))
		case "PRIVMSG":
			// get user
			user := parsedMsg.source[1:strings.Index(parsedMsg.source, "!")]
			fmt.Printf("%s:> %s\n", user, parsedMsg.message)
		case "001":
			// Logged in (successfully authenticated).
			fallthrough
		case "002", "003", "004":
			fallthrough
		case "353":
			// Tells you who else is in the chat room you're joining.
			fallthrough
		case "366", "372", "375", "376":
			fmt.Printf("C:%s %s\n", parsedMsg.command, parsedMsg.message)
		}
	}
}
