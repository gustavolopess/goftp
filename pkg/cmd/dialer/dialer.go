package dialer

import (
	"fmt"
	"io"
	"net"
)

func Dial(addr, command string, cmdHandler func(string, io.ReadWriteCloser), respHandler func(string, []byte) error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cmdHandler(command, conn)

	conn.CloseWrite()

	reply := make([]byte, 1024*10)
	conn.Read(reply)

	respHandler(command, reply)
}
