package listener

import (
	"io"
	"log"
	"net"
	"strconv"
)

func Listen(port int, handler func(io.ReadWriteCloser)) error {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		panic(err)
	}

	log.Printf("Listening on port %d", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		log.Printf("New connection from %s", conn.RemoteAddr())

		go handler(conn)
	}
}
