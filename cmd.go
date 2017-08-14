package main

import (
	"io/ioutil"
	"log"
	"fmt"
	"strings"
)

func ListFile(args []string) {
	path := GetPath(args)
	files, err := ioutil.ReadDir(GetCurrentFolder(path))
	if err != nil {
		log.Fatal(err)
	}
	PrintDelimiter()
	for _, file := range files {
		fmt.Printf("| %-40s | %10d | %4s |\n", file.Name(), file.Size(), IsDirFlag(file.IsDir()))
	}
	PrintDelimiter()
}

func ChangeDir(path []string) {
	newPath := GetPath(path)
	switch newPath {
	case "..":
		if CurrentFolder != RootFolder {
			CurrentFolder = CurrentFolder[0:strings.LastIndex(CurrentFolder, "/")]
		} else {
			CurrentFolder = RootFolder
		}
	case "":
		CurrentFolder = RootFolder
	default:
		CurrentFolder = CurrentFolder + "/" + newPath
	}
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
