package cryptolabs

import (
	"strconv"
	"strings"
)

func NewMagicSquare() Cipher {
	return &magicSquare{'_'}
}

type magicSquare struct {
	dummy rune
}

func (square *magicSquare) Encrypt(text, key string) string {
	textRunes := []rune(text)
	keyStrings := strings.Split(key, ",")
	if _, ok := sqrt(len(keyStrings)); !ok {
		return text
	}
	mappings := make([]int, len(keyStrings))
	for i := range keyStrings {
		value, err := strconv.ParseInt(keyStrings[i], 10, 64)
		if err != nil || value < 1 || int(value) > len(keyStrings) {
			return text
		}
		mappings[value-1] = i
	}
	length := len(textRunes)
	if modulo := length % len(keyStrings); modulo != 0 {
		length += len(keyStrings) - modulo
	}
	newRunes := make([]rune, length)
	for i := range newRunes {
		newRunes[i] = square.dummy
	}
	for i := range textRunes {
		modulo := i % len(keyStrings)
		newRunes[i-modulo+mappings[modulo]] = textRunes[i]
	}
	return string(newRunes)
}

func (square *magicSquare) Decrypt(text, key string) string {
	textRunes := []rune(text)
	keyStrings := strings.Split(key, ",")
	if _, ok := sqrt(len(keyStrings)); !ok {
		return text
	}
	mappings := make([]int, len(keyStrings))
	for i := range keyStrings {
		value, err := strconv.ParseInt(keyStrings[i], 10, 64)
		if err != nil || value < 1 || int(value) > len(keyStrings) {
			return text
		}
		mappings[i] = int(value) - 1
	}
	newRunes := make([]rune, len(textRunes))
	for i := range textRunes {
		modulo := i % len(keyStrings)
		newRunes[i-modulo+mappings[modulo]] = textRunes[i]
	}
	return strings.ReplaceAll(string(newRunes), string(square.dummy), "")
}
