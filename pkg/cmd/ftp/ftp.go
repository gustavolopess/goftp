package ftp

import (
	"bufio"
	"io"
	"log"
	"path/filepath"
	"strings"

	"github.com/gustavolopess/goftp/pkg/cmd/get"
	"github.com/gustavolopess/goftp/pkg/cmd/ls"
)

type FtpServer interface {
	Get(relativePath string) (string, error)
	ListIt(relativePath string) (string, error)
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

func (f *ftp) ListIt(relativePath string) (string, error) {
	return ls.ListIt(relativePath)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func (f *ftp) CommandHandler(reader io.ReadWriter) {
	command, _ := bufio.NewReader(reader).ReadString('\n')

	log.Printf("Command: %s", command)

	args := strings.Split(command, " ")

	log.Println(args, args[0], args[1])

	switch args[0] {
	case FTP_GET:
		filePath := strings.TrimSpace(args[1])
		fileContent, err := f.Get(filePath)
		check(err)
		io.WriteString(reader, fileContent)
	case FTP_LS:
		dirPath := strings.TrimSpace(args[1])
		dirDescription, err := f.ListIt(dirPath)
		check(err)
		io.WriteString(reader, dirDescription)

	default:
		return
	}
}
