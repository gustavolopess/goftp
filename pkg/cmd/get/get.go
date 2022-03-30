package get

import (
	"io/ioutil"
)

func Get(filePath string) (string, error) {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(fileContent), nil
}
