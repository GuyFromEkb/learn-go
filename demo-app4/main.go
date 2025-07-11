package main

import (
	// "demo-app4/scan"
	// "demo-app4/urlCredentials"
	"demo-app4/file"
)

func main() {
	// file.ReadFile("./urlCredentials.json")
	file.WriteFile("./urlCredentials.json")
	// url := scan.ScanPrompt("enter url: ", false)
	// login := scan.ScanPrompt("enter login: ", false)
	// password := scan.ScanPrompt("enter password: ", true)

	// userData, err := urlCredentials.NewUrlCredentials(url, login, password)

	// if err != nil {
	// 	panic(err)
	// }

	// userData.OutputCredentials()

}
