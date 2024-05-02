package fileoperations

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
)

// Read file

func ReadFile(fileName string) []byte {
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Println("error reading file", err)
		return nil
	}

	return file

}

func WriteFile(fileName string, data string) {

	err := os.WriteFile(fileName, []byte(data), fs.ModePerm)
	if err != nil {
		log.Println(err)
		return
	}
}

func UpdateFile(fileName string, data string) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Println("error reading file", err)

	}

	newContent := fmt.Sprintf("%s \n %s", file, data)

	err = os.WriteFile(fileName, []byte(newContent), fs.ModePerm)
	if err != nil {
		log.Println(err)
		return
	}

}

func DeleteFile(path string) {
	err := os.Remove(path)
	if err != nil {
		return
	}

}

func ReadLargeFile(filePath string) {
	open, err := os.Open(filePath)
	if err != nil {
		return
	}

	reader := bufio.NewReader(open)
	for {
		chunk := make([]byte, 1024)
		read, err := reader.Read(chunk)
		if err != nil {
			return
		}

		log.Println("number of bytes read", read)

		log.Println(string(chunk))
	}

}
