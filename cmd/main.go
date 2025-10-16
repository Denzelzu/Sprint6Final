package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "[MORSE_SERVER]", log.LstdFlags)

	srv := server.MyServer(logger)

	logger.Println("Сервер запускается на http://localhost:8080")
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
