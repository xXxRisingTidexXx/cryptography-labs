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
	newRunes := make([]rune, len(text))
	for i := 0; i < len(text); i += 2 {
		left, right := textRunes[i], textRunes[0]
		isEven := i+1 != len(textRunes)
		if isEven {
			right = textRunes[i+1]
		}
		for j := range key {
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

func (feistel *feistel) transform(text, key []rune) string {
	if len(text) == 0 || len(key) == 0 {
		return string(text)
	}
	result := make([]rune, len(text))
	for i := 0; i < len(text); i += 2 {
		left, right := text[i], text[0]
		isEven := i+1 != len(text)
		if isEven {
			right = text[i+1]
		}
		for j := range key {
			left, right = right^feistel.f(rune(j), left, key[j]), left
		}
		result[i] = left
		if isEven {
			result[i+1] = right
		}
	}
	fmt.Printf("%q\n", result)
	return string(result)
}

func (feistel *feistel) Decrypt(text, key string) string {
	return text
}
