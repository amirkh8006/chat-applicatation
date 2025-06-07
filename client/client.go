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
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	go receiveMessages(conn)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Client> ")
		message, _ := reader.ReadString('\n')
		conn.Write([]byte(message))
	}
}

func receiveMessages(conn net.Conn) {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			slog.Error(err.Error())
		}
		fmt.Print("Server: " + message)
	}
}