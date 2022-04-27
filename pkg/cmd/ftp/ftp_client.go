package ftp

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/gustavolopess/goftp/pkg/cmd/cd"
	"github.com/gustavolopess/goftp/pkg/cmd/ls"
)

type FtpClient interface {
	CommandHandler(string, io.ReadWriteCloser)
	ResponseHandler(string, []byte) error
}

type client struct {
	ResponseHandlers map[string]func([]byte)
}

func NewFtpClient() FtpClient {
	return &client{
		ResponseHandlers: map[string]func([]byte){
			FTP_CD: cd.HandleResponse,
			FTP_LS: ls.HandleResponse,
		},
	}
}

func (c *client) ResponseHandler(command string, response []byte) error {
	args := strings.Split(command, " ")
	verb := strings.TrimRight(args[0], "\n ")

	if handler, ok := c.ResponseHandlers[verb]; ok {
		handler(response)
	} else {
		return fmt.Errorf("no handler for verb '%v'", verb)
	}

	return nil
}

func (c *client) CommandHandler(command string, reader io.ReadWriteCloser) {
	args := strings.Split(command, " ")

	verb := strings.TrimRight(args[0], "\n ")

	for _, cmd := range AvailableCommands {
		if cmd == verb {
			log.Printf("Command: %s", command)

			_, err := reader.Write([]byte(command))
			if err != nil {
				log.Printf("Error writing to connection: %s", err)
				return
			}

			// Tell the server that we're done writing.
			if err != nil {
				fmt.Println(err)
				return
			}
			break
		}
	}

	//response := make([]byte, 1024)
	//reader.Read(response)

	//c.handleResponse(verb, response)
}
