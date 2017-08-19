package main

import (
	"os"
	"bufio"
	"log"
	"strings"
)

const rootFolder = "/users/aleksandrbogomolov"

const userName = "aleksandrbogomolov"

var CurrentFolder string = rootFolder

func main() {
	reader := bufio.NewReader(os.Stdin)
	printGreeting()
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
		case "cp":
			Copy(args[1], args[2])
		case "exit":
			os.Exit(0)
		default:
			log.Print("Wrong command.")
		}
		printGreeting()
	}
}

func printGreeting() {
	log.Printf("%s in %s", userName, CurrentFolder)
}
