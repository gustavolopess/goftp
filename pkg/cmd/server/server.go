package main

import (
	"os"

	"github.com/gustavolopess/goftp/pkg/cmd/ftp"
	"github.com/gustavolopess/goftp/pkg/cmd/listener"
)

func main() {
	homeDir := "/Users/gustavo.lopes/"

	os.Chdir(homeDir)

	ftpServer := ftp.NewFtpServer(homeDir)

	listener.Listen(8080, ftpServer.CommandHandler)
}
