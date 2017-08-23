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
		fmt.Printf("| %-90.90s | %15d | %4.4s |\n", file.Name(), file.Size(), IsDir(file.IsDir()))
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

func Move(src string, dst string) {
	Copy(src, dst)
	Remove(src)
}

func Remove(src string) {
	os.Remove(CurrentFolder + "/" + src)
}

func MakeDir(src string) {
	os.Mkdir(CurrentFolder + "/" + src, os.ModePerm)
}
