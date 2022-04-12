package ftp

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"path/filepath"
	"strings"

	"github.com/gustavolopess/goftp/pkg/cmd/cd"
	"github.com/gustavolopess/goftp/pkg/cmd/get"
	"github.com/gustavolopess/goftp/pkg/cmd/ls"
)

type FtpServer interface {
	Get(relativePath string) (string, error)
	ListIt(relativePath string) (string, error)
	ChangeDir(dirPath string) error
	CommandHandler(io.ReadWriteCloser)
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
	return ls.ListIt(filepath.Join(f.dir, relativePath))
}

func (f *ftp) ChangeDir(dirPath string) error {
	newWorkingDir, err := cd.ChangeDir(filepath.Join(f.dir, dirPath))
	if err != nil {
		return err
	}

	f.dir = newWorkingDir

	return nil
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func (f *ftp) CommandHandler(reader io.ReadWriteCloser) {
	command, _ := bufio.NewReader(reader).ReadString('\n')

	log.Printf("Command: %s", command)

	args := strings.Split(command, " ")

	switch strings.TrimRight(args[0], "\n ") {
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
	case FTP_CD:
		dirPath := strings.TrimSpace(args[1])
		err := f.ChangeDir(dirPath)
		check(err)
		io.WriteString(reader, fmt.Sprintf("Current directory: %s", f.dir))
	case FTP_CLOSE:
		io.WriteString(reader, "Closing connection, bye!")
		reader.Close()
	default:
		return
	}
}
