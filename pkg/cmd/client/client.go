package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gustavolopess/goftp/pkg/cmd/dialer"
	"github.com/gustavolopess/goftp/pkg/cmd/ftp"
)

func main() {
	client := ftp.NewFtpClient()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("ftp> ")
		scanner.Scan()

		command := scanner.Text()

		dialer.Dial("localhost:8080", command, client.CommandHandler, client.ResponseHandler)
	}
}
