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
		log.Printf("Command with argument: %s", args)
		if Trim(args[0]) == "dir" || Trim(args[0]) == "ls" {
			Dir(Trim(args[1]))
		} else {
			log.Print("Wrong command.")
		}
	}
}
