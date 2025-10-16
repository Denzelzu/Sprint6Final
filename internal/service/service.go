package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func MorseCode(data string) bool {

	// азбука Морзе может содержать такие символы как ".", "-", "/", " "
	// проводится проверка каждого входящего символа на совпадение символов из азбуки,
	// если выявится другой символ, то это не код Морзе

	for _, char := range data {
		if char != '.' && char != '-' && char != '/' && char != ' ' {
			// данный символ не является частью азбуки Морзе
			return false
		}
	}
	return strings.ContainsAny(data, ".-") // проверяем, чтобы данные не содержали только пробелы или слэши, в случае с проверкой выше
}

func ConvertData(data string) (string, error) {

	data = strings.TrimSpace(data) // Удаляем лишние пробелы
	if data == "" {
		return "", errors.New("входная строка пуста или состоит только из пробельных символов")
	}
	if MorseCode(data) { // проверяем соотвветствует ли содержимое Морзе
		return morse.ToText(data), nil
	}

	return morse.ToMorse(data), nil
}
