package service

import (
	"errors"
	"strings"
)

var MorseCodeMap = map[rune]string{
	'А': ".-", 'Б': "-...", 'В': ".--", 'Г': "--.", 'Д': "-..",
	'Е': ".", 'Ж': "...-", 'З': "--..", 'И': "..",
	'Й': ".---", 'К': "-.-", 'Л': ".-..", 'М': "--", 'Н': "-.",
	'О': "---", 'П': ".--.", 'Р': ".-.", 'С': "...", 'Т': "-",
	'У': "..-", 'Ф': "..-.", 'Х': "....", 'Ц': "-.-.", 'Ч': "---.",
	'Ш': "----", 'Щ': "--.-", 'Ъ': "-..-", 'Ы': "-.--", 'Ь': "-..-",
	'Э': "..-..", 'Ю': "..--", 'Я': ".-.-",
}

var ReverseMorseCodeMap map[string]rune

func init() {
	ReverseMorseCodeMap = make(map[string]rune)
	for char, code := range MorseCodeMap {
		ReverseMorseCodeMap[code] = char
	}
}

func ConvertMorseOrText(input string) (string, error) {
	if input == "" {
		return "", errors.New("input string is empty")
	}

	isMorse := isMorseCode(input)

	if isMorse {
		return morseToText(input)
	}
	return textToMorse(input)
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

func morseToText(morse string) (string, error) {
	words := strings.Split(strings.TrimSpace(morse), "   ")
	var result strings.Builder

	for i, word := range words {
		if word == "" {
			continue
		}
		chars := strings.Split(word, " ")
		for _, code := range chars {
			if code == "" {
				continue
			}
			char, exists := ReverseMorseCodeMap[code]
			if !exists {
				return "", errors.New("invalid morse code: " + code)
			}
			result.WriteRune(char)
		}
		if i < len(words)-1 {
			result.WriteRune(' ')
		}
	}

	if result.Len() == 0 {
		return "", errors.New("no valid morse code found")
	}
	return result.String(), nil
}

func textToMorse(text string) (string, error) {
	text = strings.ToUpper(strings.TrimSpace(text))
	if text == "" {
		return "", errors.New("input text is empty")
	}

	var result strings.Builder
	words := strings.Fields(text)

	for i, word := range words {
		var wordResult strings.Builder
		for _, char := range word {
			if code, exists := MorseCodeMap[char]; exists {
				wordResult.WriteString(code)
				wordResult.WriteRune(' ')
			} else {
				return "", errors.New("invalid character in text: " + string(char))
			}
		}

		wordMorse := strings.TrimSpace(wordResult.String())
		result.WriteString(wordMorse)
		if i < len(words)-1 {
			result.WriteString("   ")
		}
	}

	return result.String(), nil
}
