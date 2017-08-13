package main

import (
	"os"
	"bufio"
	"log"
	"strings"
)

var CurrentFolder string = "/users/aleksandrbogomolov/"

func main() {
	reader := bufio.NewReader(os.Stdin)
	ListFile([]string{CurrentFolder})
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		args := strings.Split(line, " ")
		log.Printf("Command with argument: %s", args)
		switch TrimRight(args[0]) {
		case "ls":
			ListFile(args)
		default:
			log.Print("Wrong command.")
		}
	}
}
