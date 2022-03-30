package main

import (
	"fmt"
	"os"

	"github.com/gustavolopess/goftp/pkg/cmd/ftp"
	"github.com/gustavolopess/goftp/pkg/cmd/listener"
)

func main() {
	os.Chdir("/home/gustavo/")

	fmt.Println(os.Getwd())

	ftpServer := ftp.NewFtpServer("/tmp/")

	listener.Listen(8080, ftpServer.CommandHandler)
}
