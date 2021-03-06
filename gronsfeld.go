package cryptolabs

import (
	"unicode"
)

func NewGronsfeld(alphabet *Alphabet) Cipher {
	return &gronsfeld{alphabet}
}

type gronsfeld struct {
	alphabet *Alphabet
}

func (gronsfeld *gronsfeld) Encrypt(text, key string) string {
	return gronsfeld.transform(text, key, 1)
}

func (gronsfeld *gronsfeld) transform(text, key string, direction int) string {
	if text == "" || key == "" {
		return text
	}
	textRunes, keyRunes := []rune(text), []rune(key)
	shifts := make([]int, len(keyRunes))
	for i := range keyRunes {
		if digit := keyRunes[i]; unicode.IsDigit(digit) {
			shifts[i] = int(digit - '0')
		}
	}
	newRunes := make([]rune, len(textRunes))
	for i := range textRunes {
		newRunes[i] = gronsfeld.alphabet.Shift(textRunes[i], direction*shifts[mod(i, len(shifts))])
	}
	return string(newRunes)
}

func (gronsfeld *gronsfeld) Decrypt(text, key string) string {
	return gronsfeld.transform(text, key, -1)
}
