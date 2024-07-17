package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const addr = "127.0.0.1:12345" // адрес сервера

func main() {
	// Подключаемся к серверу
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Connected to server", addr)

	// Читаем входящие сообщения от сервера и выводим их на стандартный вывод
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Println("Received from server:", message)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading: %v '/n'", err)
	}
}
