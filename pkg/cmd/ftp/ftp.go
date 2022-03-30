package ftp

import (
	"bufio"
	"io"
	"log"
	"path/filepath"
	"strings"

	"github.com/gustavolopess/goftp/pkg/cmd/get"
)

type FtpServer interface {
	Get(relativePath string) (string, error)
	CommandHandler(io.ReadWriter)
}

type ftp struct {
	dir string
}

func NewFtpServer(dir string) FtpServer {
	return &ftp{dir: dir}
}

func (f *ftp) Get(relativePath string) (string, error) {
	return get.Get(filepath.Join(f.dir, relativePath))
}

func (f *ftp) CommandHandler(reader io.ReadWriter) {
	command, _ := bufio.NewReader(reader).ReadString('\n')

	log.Printf("Command: %s", command)

	args := strings.Split(command, " ")

	log.Println(args, args[0], args[1])

	switch args[0] {
	case FTP_GET:
		fileContent, err := f.Get(args[1])
		if err != nil {
			panic(err)
		}
		io.WriteString(reader, fileContent)
	default:
		return
	}
}
