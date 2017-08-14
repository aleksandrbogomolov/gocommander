package main

import (
	"os"
	"bufio"
	"log"
	"strings"
)

const RootFolder string = "/users/aleksandrbogomolov"

var CurrentFolder string = RootFolder

func main() {
	reader := bufio.NewReader(os.Stdin)
	ListFile([]string{CurrentFolder})
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		args := strings.Split(strings.TrimRight(line, "\n"), " ")
		log.Printf("Command with argument: %s", args)
		switch args[0] {
		case "ls":
			ListFile(args)
		case "cd":
			ChangeDir(args)
		default:
			log.Print("Wrong command.")
		}
	}
}
