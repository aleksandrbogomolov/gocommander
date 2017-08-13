package main

import (
	"io/ioutil"
	"log"
	"fmt"
)

func ListFile(args []string) {
	path := ""
	if len(args) > 1 {
		path = args[1]
	}
	files, err := ioutil.ReadDir(GetCurrentFolder(path))
	if err != nil {
		log.Fatal(err)
	}
	PrintDelimiter()
	for _, file := range files {
		fmt.Printf("|%-40s | %10d | %4s|\n", file.Name(), file.Size(), IsDirFlag(file.IsDir()))
	}
	PrintDelimiter()
}

func Copy(from string, to string) {

}

func Move(from string, to string) {

}

func Remove(path string) {

}

func MakeFile(path string) {

}

func MakeDir(path string) {

}
