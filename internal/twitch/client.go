package twitch

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func Client(url, channel string) {
	log.Println(url)
	// init
	conn, err := net.Dial("tcp", "irc.chat.twitch.tv:6667")
	if err != nil {
		log.Panicln("ERROR", err)
	}
	defer conn.Close()
	// join channel
	loginMsg := fmt.Sprintf("PASS %s\r\nNICK %s\r\nJOIN #%s\r\n", "justinfan6493", "justinfan6493", channel)
	fmt.Println(loginMsg)
	conn.Write([]byte(loginMsg))

	buffReader := bufio.NewReader(conn)
	// var data []byte
	connected := false
	for {
		bytes, _, err := buffReader.ReadLine()
		if err != nil {
			log.Println("ERROR", err)
			connected = true

		}
		log.Println(string(bytes))

		if connected {
			break
		}
	}
}
