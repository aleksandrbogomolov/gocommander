package main

import (
	"fmt"
)

/**
Getting the current folder path.
 */
func GetCurrentFolder(path string) string {
	if path != "" {
		return path
	} else {
		return CurrentFolder
	}
}

func GetPath(args []string) string {
	if len(args) > 1 {
		return args[1]
	}
	return ""
}

/**
Returning string flag if file is dir or not.
 */
func IsDirFlag(flag bool) string {
	if flag {
		return "dir"
	}
	return "-"
}

func PrintDelimiter() {
	fmt.Println("=======================================================================================================================")
}
