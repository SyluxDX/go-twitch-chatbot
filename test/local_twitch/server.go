package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func connectChannel(conn net.Conn) (string, error) {
	var pass, channel, nick string
	timeoutDuration := 1 * time.Second
	buffReader := bufio.NewReader(conn)
	log.Println("starting connection")
	for connecting := true; connecting; {
		conn.SetReadDeadline(time.Now().Add(timeoutDuration))
		// read data
		line, err := buffReader.ReadString('\n')
		if err != nil {
			switch {
			case errors.Is(err, io.EOF):
				log.Println("Client closed connection")
				//connecting = false
				return "", err
			case errors.Is(err, os.ErrDeadlineExceeded):
				//log.Print("read timout continue")
				continue
			default:
				log.Printf("ERROR: (%T) %s\n", err, err)
				return "", err
			}

		} else {
			log.Println(strings.TrimSpace(line))
			split := strings.Split(strings.TrimSpace(line), " ")
			if split[0] == "PASS" {
				pass = split[1]
			} else if split[0] == "NICK" {
				nick = split[1]
			} else if split[0] == "JOIN" {
				channel = split[1]
			}

			if pass != "" && nick != "" && channel != "" {
				log.Println("connection done, sending initial msg")
				// send initial lines
				conn.Write([]byte(fmt.Sprintf(":tmi.twitch.tv 001 %s :Welcome, GLHF!\n", nick)))
				conn.Write([]byte(fmt.Sprintf(":tmi.twitch.tv 002 %s :Your host is tmi.twitch.tv\n", nick)))
				conn.Write([]byte(fmt.Sprintf(":tmi.twitch.tv 003 %s :This server is rather new\n", nick)))
				conn.Write([]byte(fmt.Sprintf(":tmi.twitch.tv 004 %s :-\n", nick)))
				conn.Write([]byte(fmt.Sprintf(":tmi.twitch.tv 375 %s :-\n", nick)))
				conn.Write([]byte(fmt.Sprintf(":tmi.twitch.tv 372 %s :You are in a maze of twisty passages, all alike.\n", nick)))
				conn.Write([]byte(fmt.Sprintf(":tmi.twitch.tv 376 %s :>\n", nick)))
				conn.Write([]byte(fmt.Sprintf(":%s!%s@%s.tmi.twitch.tv JOIN #%s\n", nick, nick, nick, channel)))
				conn.Write([]byte(fmt.Sprintf(":%s.tmi.twitch.tv 353 %s = #%s :%s\n", nick, nick, channel, nick)))
				conn.Write([]byte(fmt.Sprintf(":%s.tmi.twitch.tv 366 %s #%s :End of /NAMES list\n", nick, nick, channel)))
				connecting = false
			}
		}
	}
	return channel, nil
}

func clientHandler(conn net.Conn, sendMsg chan string) {
	// generate id based on epoch
	id := time.Now().Unix()
	log.Printf("Handling new connection, id: %d\n", id)

	channel, err := connectChannel(conn)
	if err != nil {
		log.Println(err)
		return
	}

	watchdog := time.NewTicker(10 * time.Second)
	outputTimer := time.NewTicker(2 * time.Second)
	pongTimer := time.NewTimer(time.Second)
	pongTimer.Stop()
	prefixMsg := fmt.Sprintf(":localserver!localserver@localserver.tmi.twitch.tv PRIVMSG #%s", channel)

	defer func() {
		log.Printf("Closing connection %d\n", id)
		conn.Close()
		watchdog.Stop()
		outputTimer.Stop()
		pongTimer.Stop()
	}()
	timeoutDuration := 1 * time.Second
	pongDuration := 5 * time.Second
	buffReader := bufio.NewReader(conn)
	messages := make(chan string)
	waitingPong := false
	running := true
	waitingClose := make(chan struct{}, 1)

	// reader function
	go func() {
		for running {
			//set deadline for reading.
			conn.SetReadDeadline(time.Now().Add(timeoutDuration))
			// read data
			line, err := buffReader.ReadString('\n')
			if err != nil {
				switch {
				case errors.Is(err, io.EOF):
					log.Println("Client closed connection")
					running = false
				case errors.Is(err, os.ErrDeadlineExceeded):
					//log.Print("read timout continue")
					continue
				default:
					log.Printf("ERROR: (%T) %s\n", err, err)
					continue
				}

			} else {
				messages <- strings.TrimSpace(line)
			}
		}
		waitingClose <- struct{}{}
	}()

	for running {
		select {
		case t := <-watchdog.C:
			log.Printf("tick at %s\n", t)
			// buffWriter.WriteString("PING")
			conn.Write([]byte(fmt.Sprintln("PING :tmi.twitch.tv")))
			// wait for PONG
			waitingPong = true
			pongTimer.Reset(pongDuration)
		case <-outputTimer.C:
			// sending message to client
			msg := <-sendMsg
			chatMsg := fmt.Sprintf("%s :%s\n", prefixMsg, msg)
			conn.Write([]byte(chatMsg))
		case msg := <-messages:
			if waitingPong {
				if msg != "PONG :tmi.twitch.tv" {
					log.Println("Missing Pong response")
					running = false
				} else {
					waitingPong = false
					pongTimer.Stop()
					log.Println("Pong received")
				}
			} else {
				log.Printf(">> %s", msg)
			}
		case <-pongTimer.C:
			if waitingPong {
				log.Println("timer: Missing Pong response")
				running = false
			}
		}
	}
	// waiting for buffer reader before
	<-waitingClose
}

func loremGenerator(output chan string) {
	// read file
	loremRaw, err := os.ReadFile("lorem.txt")
	if err != nil {
		panic(err)
	}
	lorem := strings.Split(strings.TrimSpace(string(loremRaw)), "\n")
	// infinite loop with dump
	i := 0
	for {
		output <- lorem[i]
		i += 1
		if i == len(lorem) {
			i = 0
		}
	}
}

func main() {
	var address, port string
	flag.StringVar(&address, "address", "127.0.0.1", "Local Server address")
	flag.StringVar(&port, "port", "8888", "Local Server port")
	flag.Parse()

	log.Printf("Starting Local server at port: %s\n", port)
	// init
	serverAddr := fmt.Sprintf("%s:%s", address, port)
	listener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Panicln(err)
	}
	defer listener.Close()

	text := make(chan string, 1)
	go loremGenerator(text)

	for {
		// listen for an connection
		conn, err := listener.Accept()
		if err != nil {
			log.Panicln(err)
		}
		// create go routine for connection
		go clientHandler(conn, text)
	}
}
