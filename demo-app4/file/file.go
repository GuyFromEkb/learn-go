package file

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func ReadFile(filePath string) {

	file, err := os.Open(filePath)

	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatalf("unable to close file: %v", err)
		}
	}()

	contentBytes, err := io.ReadAll(file)

	var data map[string]any

	json.Unmarshal(contentBytes, &data)

	credentials, ok := data["credentials"]
	if !ok {
		log.Fatalf("credentials field not found")
	}

	resultBytes, err := json.MarshalIndent(credentials, "", "  ")
	if err != nil {
		log.Fatalf("unable to marshal credentials: %v", err)
	}

	fmt.Print(string(resultBytes))

}

type UrlCredentials struct {
	Url      string `json:"url"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type CredentialsFile struct {
	Credentials []UrlCredentials `json:"credentials"`
}

func WriteFile(pathToFile string) {
	// Открываем файл
	file, err := os.OpenFile(pathToFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}
	defer file.Close()

	// Читаем содержимое
	contentBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var credFile CredentialsFile

	// Если файл не пустой, парсим
	if len(contentBytes) > 0 {
		err = json.Unmarshal(contentBytes, &credFile)
		if err != nil {
			log.Fatalf("unable to unmarshal json: %v", err)
		}
	} else {
		credFile.Credentials = []UrlCredentials{}
	}

	// Создаём новые данные
	newCredential := UrlCredentials{
		Url:      "https://google.com",
		Login:    "Ragnarek9000",
		Password: "goGOOOLANG",
	}

	// Добавляем в срез
	credFile.Credentials = append(credFile.Credentials, newCredential)

	// Перемещаемся в начало файла, чтобы перезаписать
	_, err = file.Seek(0, 0)
	if err != nil {
		log.Fatalf("unable to seek file: %v", err)
	}

	// Очищаем файл
	err = file.Truncate(0)
	if err != nil {
		log.Fatalf("unable to truncate file: %v", err)
	}

	// Сохраняем новый JSON
	resultBytes, err := json.MarshalIndent(credFile, "", "  ")
	if err != nil {
		log.Fatalf("unable to marshal json: %v", err)
	}

	_, err = file.Write(resultBytes)
	if err != nil {
		log.Fatalf("unable to write to file: %v", err)
	}
}
