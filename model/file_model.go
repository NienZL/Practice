package model

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

type File struct {
	Path    string
	Name    string
	SubName string
	Content string
}

func (file *File) Create(input string) {
	filePath := pathConverter(input)
	fileName := path.Base(filePath)

	file.Path = path.Dir(filePath)
	file.Name = strings.Split(fileName, ".")[0]
	file.SubName = strings.Split(fileName, ".")[1]

	err := ioutil.WriteFile(filePath, []byte(file.Content), 0666)
	if err != nil {
		log.Fatal(err)
	}

}

func (file *File) Open(input string) error {
	fileBuffer, readErr := ioutil.ReadFile(pathConverter(input))
	if readErr != nil {
		fmt.Fprintf(os.Stderr, "Read File Error: %s\n", readErr)
		return readErr
	}

	filePath := pathConverter(input)
	fileName := path.Base(filePath)

	file.Path = path.Dir(filePath)
	file.Name = strings.Split(fileName, ".")[0]
	file.SubName = strings.Split(fileName, ".")[1]
	file.Content = string(fileBuffer)

	return nil
}

func (file *File) Write(input string, result string) {
	tempContent := input + " = " + result + "\n"
	file.Content += tempContent
}

func (file *File) Save() error {
	err := ioutil.WriteFile(file.Path+"/"+file.Name+"."+file.SubName, []byte(file.Content), 0666)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("檔案路徑 : %s\n檔案名稱 : %s.\n", file.Path, file.Name+"."+file.SubName)
	return nil
}

func pathConverter(input string) string {
	var replacer = strings.NewReplacer("\\", "/")
	newInput := replacer.Replace(input)
	return newInput
}

// func formatNumber(input string) string {

// }
