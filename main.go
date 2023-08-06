package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

var proverbs = []string{
	"Don't communicate by sharing memory, share memory by communicating.",
	"Concurrency is not parallelism.",
	"Channels orchestrate; mutexes serialize.",
	"The bigger the interface, the weaker the abstraction.",
	"Make the zero value useful.",
	"interface{} says nothing.",
	"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
	"A little copying is better than a little dependency.",
	"Syscall must always be guarded with build tags.",
	"Cgo must always be guarded with build tags.",
	"Cgo is not Go.",
	"With the unsafe package there are no guarantees.",
	"Clear is better than clever.",
	"Reflection is never clear.",
	"Errors are values.",
	"Don't just check errors, handle them gracefully.",
	"Design the architecture, name the components, document the details.",
	"Documentation is for users.",
	"Don't panic.",
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		proverb := proverbs[rand.Intn(len(proverbs))]
		_, err := conn.Write([]byte(proverb + "\n"))
		if err != nil {
			fmt.Printf("Failed to send data to client: %v\n", err)
			return
		}
		time.Sleep(3 * time.Second)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	listener, err := net.Listen("tcp", ":18080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Server started on :18080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Failed to accept connection: %v\n", err)
			continue
		}

		go handleConnection(conn)
	}
}
