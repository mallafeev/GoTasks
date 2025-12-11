package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("пишите: go run client.go localhost:8081")
		return
	}
	serverAddr := os.Args[1]

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatal("ошибка", err)
	}
	defer conn.Close()

	fmt.Print("придумайте имя ")
	reader := bufio.NewReader(os.Stdin)
	nick, _ := reader.ReadString('\n')
	nick = strings.TrimSpace(nick)

	_, err = conn.Write([]byte("NICK:" + nick + "\n"))
	if err != nil {
		log.Fatal("ошибка", err)
	}
	fmt.Println("Отправлен ник:", nick)

	response, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("Получен ответ:", response)
	response = strings.TrimSpace(response)

	if response == "такое имя занято, увы" {
		fmt.Println("ошибка", response[7:])
		return
	}

	fmt.Println("успешный вход, для просмотра комманд введите help")

	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			msg := scanner.Text()
			fmt.Println("\n[СООБЩЕНИЕ] " + msg)
			fmt.Print("введите команду (help для справки): ")
		}
	}()

	for {
		fmt.Print("введите команду (help для справки): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "help" {
			fmt.Println("команды:")
			fmt.Println("send <ник_получателя> <сообщение> - отправить сообщение пользователю")
			fmt.Println("exit - выйти из программы")
			continue
		}

		if input == "exit" {
			break
		}

		if strings.HasPrefix(input, "send ") {
			parts := strings.SplitN(input[5:], " ", 2)
			if len(parts) != 2 {
				fmt.Println("используйте: send <ник> <сообщение>")
				continue
			}
			recipient := parts[0]
			message := parts[1]

			sendMsg := fmt.Sprintf("SEND:%s:%s\n", recipient, message)
			_, err := conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("ошибка", err)
				continue
			}

			response, _ := bufio.NewReader(conn).ReadString('\n')
			response = strings.TrimSpace(response)
			if strings.HasPrefix(response, "ERROR:") {
				fmt.Println("ошибка", response[6:])
			} else {
				fmt.Println("отправлено!")
			}
			continue
		}

		fmt.Println("такой команды нет. help - справка")
	}
}
