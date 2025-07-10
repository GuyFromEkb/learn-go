package urlCredentials

import (
	"demo-app4/scan"
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"strconv"

	"github.com/vinay03/chalk"
)

type UrlCredentials struct {
	url      string
	login    string
	password string
}

func NewUrlCredentials(urlStr, login, password string) (*UrlCredentials, error) {
	var userData UrlCredentials

	_, err := url.ParseRequestURI(urlStr)

	if err != nil {
		fmt.Println("Некорректный ввод url!")
		return nil, errors.New("INVALID_URL")
	}

	userData.login = login
	userData.url = urlStr

	if password == "" {
		userData.generatePassword()
	} else {
		userData.password = password
	}

	return &userData, nil

}

func (credentials *UrlCredentials) generatePassword() {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	passwordLengthStr := scan.ScanPrompt("enter password length: ", false)
	passwordLength, _ := strconv.ParseInt(passwordLengthStr, 10, 32)

	passwordRune := make([]rune, passwordLength)

	for i := range passwordRune {
		rndIdx := rand.IntN(len(letterRunes))
		passwordRune[i] = letterRunes[rndIdx]
	}

	credentials.password = string(passwordRune)
}

func (credentials *UrlCredentials) OutputCredentials() {

	chalk.Yellow().Println(credentials.url)
	chalk.Green().Println(credentials.login)
	chalk.Cyan().Println(credentials.password)

}
