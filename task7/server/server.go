package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

// в лекции не было структуры, сделал для хранения имени
type Client struct {
	conn net.Conn
	nick string
}

var (
	clients = make(map[string]*Client)
	mutex   sync.RWMutex
)

func main() {
	listener, err := net.Listen("tcp", ":8081") //8080 использую для других работ
	if err != nil {
		log.Fatal("ошибка", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("ошибка", err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	client := &Client{conn: conn}

	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			removeClient(client.nick)
			return
		}

		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "NICK:") {
			nick := line[5:]
			mutex.Lock()
			if _, exists := clients[nick]; exists {
				conn.Write([]byte("придумайте другое имя, такое есть\n"))
				mutex.Unlock()
				return
			}
			conn.Write([]byte("OK\n"))
			client.nick = nick
			clients[nick] = client
			fmt.Printf("%s зашёл\n", nick)
			mutex.Unlock()
			continue
		}

		if strings.HasPrefix(line, "SEND:") {
			parts := strings.SplitN(line[5:], ":", 2)
			if len(parts) != 2 {
				conn.Write([]byte("ошибка. используйте: send <ник> <сообщение>\n"))
				continue
			}
			recipient := parts[0]
			message := parts[1]

			mutex.RLock()
			targetClient, exists := clients[recipient]
			mutex.RUnlock()

			if !exists {
				conn.Write([]byte("такого юзера нет\n"))
				continue
			}

			targetClient.conn.Write([]byte(fmt.Sprintf("От %s: %s", client.nick, message)))
			conn.Write([]byte("OK\n"))
			continue
		}
	}
}

func removeClient(nick string) {
	mutex.Lock()
	delete(clients, nick)
	fmt.Printf("%s вышел\n", nick)
	mutex.Unlock()
}
