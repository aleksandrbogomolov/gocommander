package main

import (
	"io/ioutil"
	"log"
	"fmt"
	"strings"
	"os"
	"io"
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
		if CurrentFolder != rootFolder {
			CurrentFolder = CurrentFolder[0:strings.LastIndex(CurrentFolder, "/")]
		} else {
			CurrentFolder = rootFolder
		}
	case "":
		CurrentFolder = rootFolder
	default:
		CurrentFolder = CurrentFolder + "/" + newPath
	}
}

func Copy(src string, dst string) {
	in, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied file %s to %s.", in, out)
}

func Move(from string, to string) {

}

func Remove(path string) {

}

func MakeFile(path string) {

}

func MakeDir(path string) {

}
