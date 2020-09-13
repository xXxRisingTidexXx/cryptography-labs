package cryptolabs

import (
	"unicode"
)

var Gronsfeld = NewGronsfeld(DefaultAlphabet)

func NewGronsfeld(alphabet *Alphabet) Cipher {
	return &gronsfeld{alphabet}
}

type gronsfeld struct {
	alphabet *Alphabet
}

func (gronsfeld *gronsfeld) Encrypt(message, key string) string {
	return gronsfeld.transform(message, key, 1)
}

func (gronsfeld *gronsfeld) transform(message, key string, direction int) string {
	if message == "" || key == "" {
		return ""
	}
	messageRunes, keyRunes := []rune(message), []rune(key)
	shifts := make([]int, len(keyRunes))
	for i := range keyRunes {
		if digit := keyRunes[i]; unicode.IsDigit(digit) {
			shifts[i] = int(digit - '0')
		}
	}
	newRunes := make([]rune, len(messageRunes))
	for i := range messageRunes {
		newRunes[i] = gronsfeld.alphabet.Get(messageRunes[i], direction*shifts[i%len(shifts)])
	}
	return string(newRunes)
}

func (gronsfeld *gronsfeld) Decrypt(message, key string) string {
	return gronsfeld.transform(message, key, -1)
}
