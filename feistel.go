package cryptolabs

import (
	log "github.com/sirupsen/logrus"
)

func NewFeistel(f func(rune, rune) rune) Cipher {
	return &feistel{f}
}

type feistel struct {
	f func(rune, rune) rune
}

func (feistel *feistel) Encrypt(text, key string) string {
	textRunes, keyRunes := []rune(text), []rune(key)
	if len(textRunes) == 0 || len(keyRunes) == 0 || len(textRunes)%2 == 1 {
		log.Debugf("cryptolabs: feistel got insufficient input, returned no digest")
		return text
	}
	newRunes := make([]rune, len(textRunes))
	for i := 0; i < len(textRunes); i += 2 {
		left, right := textRunes[i], textRunes[i+1]
		for j := range keyRunes {
			left, right = right^feistel.f(left, keyRunes[j]), left
			feistel.debugf("encryption", i, j, left, right)
		}
		newRunes[i], newRunes[i+1] = left, right
	}
	return string(newRunes)
}

func (feistel *feistel) debugf(process string, i, j int, left, right rune) {
	log.Debugf(
		"cryptolabs: feistel %s, block %d, round %d, left %q / %x, right %q / %x",
		process,
		i/2,
		j,
		left,
		left,
		right,
		right,
	)
}

func (feistel *feistel) Decrypt(text, key string) string {
	textRunes, keyRunes := []rune(text), []rune(key)
	if len(textRunes) == 0 || len(keyRunes) == 0 || len(textRunes)%2 == 1 {
		log.Debugf("cryptolabs: feistel got insufficient input, returned no message")
		return text
	}
	newRunes := make([]rune, len(textRunes))
	for i := 0; i < len(textRunes); i += 2 {
		left, right := textRunes[i], textRunes[i+1]
		for j := len(keyRunes) - 1; j >= 0; j-- {
			left, right = right, left^feistel.f(right, keyRunes[j])
			feistel.debugf("decryption", i, j, left, right)
		}
		newRunes[i], newRunes[i+1] = left, right
	}
	return string(newRunes)
}
