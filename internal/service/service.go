package service

import (
	"errors"
	"regexp"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

var morsePattern = regexp.MustCompile(`^[.\-/\s]*$`)

func MorseCode(data string) bool {
	//  Проверяем, содержит ли строка ТОЛЬКО символы, которые могут быть в Морзе
	if !morsePattern.MatchString(data) {
		return false // Содержит недопустимые символы
	}

	// Убеждаемся, что это не пустая строка и содержит хотя бы один "сигнал"
	cleaned := strings.TrimSpace(data)
	if cleaned == "" {
		return false
	}

	//Должен содержать хотя бы точку или тире, иначе это просто пробелы/слэши
	return strings.ContainsAny(cleaned, ".-")
}

func ConvertData(data string) (string, error) {

	data = strings.TrimSpace(data)
	if data == "" {
		return "", errors.New("входная строка пуста или состоит только из пробельных символов")
	}

	if MorseCode(data) {
		return morse.ToText(data), nil
	}

	return morse.ToMorse(data), nil
}
