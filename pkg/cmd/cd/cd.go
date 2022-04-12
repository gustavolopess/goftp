package cd

import (
	"fmt"
	"os"
)

func ChangeDir(dirpath string) (string, error) {
	if err := os.Chdir(dirpath); err != nil {
		return "", err
	}

	return os.Getwd()
}

func HandleResponse(response []byte) {
	fmt.Println(response)
}
