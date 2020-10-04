package cryptolabs

import "fmt"

func NewFeistel(f func(rune, rune, rune) rune) Cipher {
	return &feistel{f}
}

type feistel struct {
	f func(rune, rune, rune) rune
}

func (feistel *feistel) Encrypt(text, key string) string {
	if len(text) == 0 || len(key) == 0 {
		return text
	}
	textRunes, keyRunes := []rune(text), []rune(key)
	newRunes := make([]rune, len(textRunes))
	for i := 0; i < len(textRunes); i += 2 {
		left, right := textRunes[i], textRunes[0]
		isEven := i+1 != len(textRunes)
		if isEven {
			right = textRunes[i+1]
		}
		for j := range keyRunes {
			left, right = right^feistel.f(rune(j), left, keyRunes[j]), left
		}
		newRunes[i] = left
		if isEven {
			newRunes[i+1] = right
		}
	}
	fmt.Printf("%q\n", newRunes)
	return string(newRunes)
}

func (feistel *feistel) Decrypt(text, key string) string {
	return text
}
