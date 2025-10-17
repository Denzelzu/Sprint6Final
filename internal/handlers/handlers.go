package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// OutputHTMLHandler обрабатывает GET-запрос и отдает HTML-форму из файла index.html.
func OutputHTMLHandler(res http.ResponseWriter, req *http.Request) {

	http.ServeFile(res, req, "index.html")
}

// UpLoadHandler обрабатывает POST-запрос с загруженным файлом,
// конвертирует его, сохраняет результат локально и возвращает его клиенту.
func UpLoadHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	// Парсинг HTML-формы (лимит 10MB)
	err := req.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Println("Ошибка при парсинге формы:", err)
		http.Error(res, "Ошибка парсинга формы: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Извлечение загруженного файла
	file, header, err := req.FormFile("myFile")
	if err != nil {
		log.Println("Ошибка при получении файла:", err)
		http.Error(res, "Ошибка при получении файла: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Чтение данных из файла
	dataFile, err := io.ReadAll(file)
	if err != nil {
		log.Println("Ошибка при чтении файла:", err)
		http.Error(res, "Ошибка при чтении файла: "+err.Error(), http.StatusInternalServerError)
		return
	}
	fileContent := string(dataFile)

	// Конвертация данных через пакет service
	convertedText, err := service.ConvertData(fileContent)
	if err != nil {
		log.Println("Ошибка конвертации:", err)
		http.Error(res, "Ошибка конвертации: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Создание локального файла (согласно требованию задания)
	ext := filepath.Ext(header.Filename)
	if ext == "" {
		ext = ".txt"
	}
	newFilename := "converted_" + time.Now().UTC().Format("20060102_150405") + ext

	outFile, err := os.Create(newFilename)
	if err != nil {
		log.Println("Ошибка при создании локального файла:", err)
		http.Error(res, "Ошибка при создании файла на сервере: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	// Запись результата конвертации в локальный файл
	_, err = outFile.WriteString(convertedText)
	if err != nil {
		log.Println("Ошибка при записи результата конвертации:", err)
		http.Error(res, "Ошибка при записи в локальный файл: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Формирование сообщения с оригинальным текстом для прохождения теста
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// Включаем оригинальный текст (fileContent) в ответ.
	responseMessage := fmt.Sprintf(
		"Конвертация завершена. Результат сохранен в файл: %s\n\n"+
			"Оригинальный текст:\n%s\n\n"+ // Строка с оригинальным текстом
			"Конвертированный текст:\n%s",
		newFilename,
		fileContent,
		convertedText,
	)

	res.WriteHeader(http.StatusOK)
	res.Write([]byte(responseMessage))
}
