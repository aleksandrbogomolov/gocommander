package main

import (
	"os"
	"bufio"
	"log"
	"strings"
)

const RootFolder string = "/users/aleksandrbogomolov/"

var CurrentFolder string = ""

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		log.Print(GetCurrentFolder())
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		args := strings.Split(line, " ")
		log.Printf("Command with argument: %s, %s", args[0], args[1])
		if strings.TrimRight(args[0], "\n") == "dir" || strings.TrimRight(args[0], "\n") == "ls" {
			Dir(strings.TrimRight(args[1], "\n"))
		} else {
			log.Print("Wrong command.")
		}
	}
}
