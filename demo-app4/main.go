package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"os"
	"strconv"
)

type urlCredentials struct {
	url      string
	login    string
	password string
}

func newUrlCredentials(urlStr, login, password string) (*urlCredentials, error) {
	var userData urlCredentials

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

func (credentials *urlCredentials) generatePassword() {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	passwordLengthStr := scanPrompt("enter password length: ", false)
	passwordLength, _ := strconv.ParseInt(passwordLengthStr, 10, 32)

	passwordRune := make([]rune, passwordLength)

	for i := range passwordRune {
		rndIdx := rand.IntN(len(letterRunes))
		passwordRune[i] = letterRunes[rndIdx]
	}

	credentials.password = string(passwordRune)
}

func (credentials *urlCredentials) outputCredentials() {
	fmt.Println(credentials)
}

func main() {
	url := scanPrompt("enter url: ", false)
	login := scanPrompt("enter login: ", false)
	password := scanPrompt("enter password: ", true)

	userData, err := newUrlCredentials(url, login, password)

	if err != nil {
		panic(err)
	}

	userData.outputCredentials()

}

func scanPrompt(placeholder string, canEmpty bool) string {

	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print(placeholder)
		scanner.Scan()
		input := scanner.Text()

		if input == "" && !canEmpty {
			fmt.Println("Некорректный ввод. Значение не может быть пустым, повторите попытку.")
			continue
		}

		return input
	}

}
