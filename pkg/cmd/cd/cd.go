package cd

import (
	"os"
)

func ChangeDir(dirpath string) (string, error) {
	if err := os.Chdir(dirpath); err != nil {
		return "", err
	}

	return os.Getwd()
}
