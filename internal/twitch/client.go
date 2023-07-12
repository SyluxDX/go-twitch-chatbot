package twitch

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"regexp"
	"strings"
)

// func joinChannel(conn net.Conn) error {

// 	return nil
// }

func (conf TwitchConfigs) ReadChat() {
	// init
	conn, err := net.Dial("tcp", conf.TwitchIRL)
	if err != nil {
		log.Panicln("ERROR", err)
	}
	defer conn.Close()
	// join channel
	log.Printf("Joinning Channel %s\n", conf.Channel)
	loginMsg := fmt.Sprintf("PASS %s\r\nNICK %s\r\nJOIN #%s\r\n", "justinfan6493", "justinfan6493", conf.Channel)
	conn.Write([]byte(loginMsg))

	buffReader := bufio.NewReader(conn)
	// var data []byte
	// Process init list of names
	connected := false
	for {
		bytes, _, err := buffReader.ReadLine()
		if err != nil {
			log.Println("ERROR", err)
			connected = true

		}
		line := string(bytes)
		if strings.HasSuffix(line, "End of /NAMES list") {
			connected = true
		}
		// log.Println(line)
		if connected {
			break
		}
	}

	// read block
	log.Println("Echo messages")
	msgPattern := regexp.MustCompile(`:(?P<user>.+)!\S+ PRIVMSG #\S+ :(?P<msg>.*)`)
	// msgPattern := regexp.MustCompile(`:(.+)!\S+ PRIVMSG #\S+ :(.*)`)
	for {
		bytes, _, err := buffReader.ReadLine()
		if err != nil {
			log.Println("ERROR", err)
			break
		}
		line := string(bytes)
		// write to output file

		match := msgPattern.FindStringSubmatch(line)
		if len(match) == 0 {
			log.Println(line)
		} else {
			fmt.Printf("%s:> %s\n", match[msgPattern.SubexpIndex("user")], match[msgPattern.SubexpIndex("msg")])
		}
		// fmt.Println(string(bytes), "->", match)
		// fmt.Printf("%s:> %s\n", match[0], match[1])
		// log.Println(string(bytes))
	}
}
