package ls

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type DirFileDescription struct {
	Name  string `json:"name"`
	IsDir bool   `json:"isDir"`
	Size  int64  `json:"size"`
}

func ListIt(dirPath string) (string, error) {
	dirFiles, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return "", err
	}

	dirFilesDescriptions := make([]DirFileDescription, len(dirFiles))
	for i, dirFile := range dirFiles {
		dirFilesDescriptions[i] = DirFileDescription{
			Name:  dirFile.Name(),
			IsDir: dirFile.IsDir(),
			Size:  dirFile.Size(),
		}
	}

	dirFilesDescriptionsJson, err := json.Marshal(dirFilesDescriptions)
	if err != nil {
		return "", err
	}

	return string(dirFilesDescriptionsJson), nil
}

func HandleResponse(response []byte) {
	var dirs []DirFileDescription

	err := json.Unmarshal(response, &dirs)
	if err != nil {
		log.Fatalf("Error unmarshalling response: %s\n", err)
	}

	fmt.Println(dirs)
}
