package main

import (
	"demo-app4/scan"
	"demo-app4/urlCredentials"
)

func main() {
	url := scan.ScanPrompt("enter url: ", false)
	login := scan.ScanPrompt("enter login: ", false)
	password := scan.ScanPrompt("enter password: ", true)

	userData, err := urlCredentials.NewUrlCredentials(url, login, password)

	if err != nil {
		panic(err)
	}

	userData.OutputCredentials()

}
