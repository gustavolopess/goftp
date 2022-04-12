package ftp

const (
	FTP_GET   = "get"
	FTP_LS    = "ls"
	FTP_CD    = "cd"
	FTP_CLOSE = "close"
)

var AvailableCommands = []string{FTP_GET, FTP_LS, FTP_CD, FTP_CLOSE}
