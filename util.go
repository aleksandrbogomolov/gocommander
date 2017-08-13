package main

func GetCurrentFolder() string {
	if CurrentFolder == "" {
		return RootFolder
	} else {
		return CurrentFolder
	}
}
