package main

import (
	"strings"
	"fmt"
)

/**
Getting the current folder path.
 */
func GetCurrentFolder(path string) string {
	if path != "" {
		return TrimRight(path)
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

/**
Trim the right delimiter from getting strings.
 */
func TrimRight(s string) string {
	return strings.TrimRight(s, "\n")
}

func PrintDelimiter() {
	fmt.Println("==============================================================")
}
