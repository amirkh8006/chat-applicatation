package main

import (
	"bufio"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
)

func main()  {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()

	slog.Info("Server started on port 8080")

	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	go recieveMessages(conn)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Server> ")
		scanner.Scan()
		message := scanner.Text()
		if message == "/quit" {
			break
		}
		sendMessage(conn, message)
	}

	slog.Info("Server shutting down")
}

func sendMessage(conn net.Conn, message string) {
	_, err := conn.Write([]byte(message + "\n"))
	if err != nil {
		slog.Error("Error sending message", "error", err)
		return
	}
}


func recieveMessages(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Printf("Client> %s" , message)
	}
}