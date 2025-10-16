package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func OutputHTMLHandler(res http.ResponseWriter, req *http.Request) {

	http.ServeFile(res, req, "index.html")

	file, err := os.Open("../index.html") // открываем файл index.html
	if err != nil {                       // в случае возникновения ошибки - выводим сообщение и статус ошибки
		http.Error(res, "не удалось найти index.html:"+err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	io.Copy(res, file) // копиурем содержимое файла указанного выше, и передаем в HTTP-ответ
}
func UpLoadHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	// Парсинг HTML-формы (лимит 10MB)
	err := req.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Println("Ошибка при парсинге формы:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Файлы в форме:", req.MultipartForm.File)

	file, header, err := req.FormFile("myFile")
	if err != nil {
		log.Println("Ошибка при получении файла:", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Чтение данных из файла
	dataFile, err := io.ReadAll(file)
	if err != nil {
		log.Println("Ошибка при чтении файла:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	fileContent := string(dataFile)
	originalText := fileContent // Сохраняем исходный текст

	convertedText, err := service.ConvertData(fileContent) // конвертация данных через пакет service
	if err != nil {
		log.Println("Ошибка конвертации:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	// создание локального файла
	ext := filepath.Ext(header.Filename)
	newFilename := "converted_" + time.Now().UTC().Format("20060102_150405") + ext

	outFile, err := os.Create(newFilename)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	//  запись результата конвертации в локальный файл
	_, err = outFile.WriteString(convertedText)
	if err != nil {
		log.Println("Ошибка при записи результата конвертации:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	//  вернуть результат конвертации строки пользователю
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	res.Write([]byte("Конвертация завершена. Результат сохранен в файл: " + newFilename + "\n\n"))
	res.Write([]byte(originalText))

}
