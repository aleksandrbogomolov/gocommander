package main

import "strings"

/**
Getting the current folder path.
 */
func GetCurrentFolder() string {
	if CurrentFolder == "" {
		return RootFolder
	} else {
		return CurrentFolder
	}
}

/**
Trim the right delimiter from getting strings.
 */
func Trim(s string) string {
	return strings.TrimRight(s, "\n")
}
