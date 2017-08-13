package main

import (
	"io/ioutil"
	"log"
)

func Dir(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		log.Println(file.Name())
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
