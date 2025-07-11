package service

import (
	"errors"
	"strings"

	"github.com/ElenaMask/go1fl-sprint6-final/pkg/morse"
)

func ConvertMorseOrText(input string) (string, error) {
	if input == "" {
		return "", errors.New("input string is empty")
	}

	isMorse := isMorseCode(input)

	if isMorse {
		return morse.ToText(input), nil
	}
	return morse.ToMorse(input), nil
}

func isMorseCode(input string) bool {
	trimmed := strings.TrimSpace(input)
	for _, char := range trimmed {
		if char != '.' && char != '-' && char != ' ' {
			return false
		}
	}
	return strings.Contains(trimmed, ".") || strings.Contains(trimmed, "-")
}
