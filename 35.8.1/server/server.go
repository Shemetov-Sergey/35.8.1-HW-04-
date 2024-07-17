package main

import (
	"log"
	"math/rand"
	"net"
	"time"
)

const addr = "0.0.0.0:12345"
const proto = "tcp4"

var sayings = []string{
	"Общайтесь не путем обмена памятью, делитесь памятью путем общения.",
	"Параллелизм - это не параллелизм.",
	"Каналы организуются; мьютексы сериализуются.",
	"Чем больше интерфейс, тем слабее абстракция.",
	"Сделайте нулевое значение полезным.",
	"интерфейс {} ничего не говорит.",
	"Стиль Gofmt никому не нравится, но gofmt нравится всем.",
	"Небольшое копирование лучше, чем небольшая зависимость.",
	"Системный вызов всегда должен быть защищен тегами сборки.",
	"Cgo всегда должен быть защищен тегами сборки.",
	"Cgo - это не Go.",
	"С небезопасной упаковкой нет никаких гарантий.",
	"Ясно лучше, чем умно.",
	"Размышления никогда не бывают ясными.",
	"Ошибки - это ценности.",
	"Не просто проверяйте ошибки, обрабатывайте их корректно.",
	"Разработайте архитектуру, назовите компоненты, задокументируйте детали.",
	"Документация предназначена для пользователей.",
	"Не паникуйте.",
}

func main() {
	listener, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			saying := sayings[rand.Intn(len(sayings))]
			_, err := conn.Write([]byte(saying + "/n"))
			if err != nil {
				return // Если есть ошибка, завершаем соединение
			}
		}
	}
}
