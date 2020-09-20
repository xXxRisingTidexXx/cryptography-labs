package cryptolabs

func NewXOR() Cipher {
	return &xor{}
}

type xor struct{}

func (xor *xor) Encrypt(text, key string) string {
	textRunes, keyRunes := []rune(text), []rune(key)
	newRunes := make([]rune, len(textRunes))
	for i := range textRunes {
		newRunes[i] = textRunes[i] ^ keyRunes[i%len(keyRunes)]
	}
	return string(newRunes)
}

func (xor *xor) Decrypt(text, key string) string {
	return xor.Encrypt(text, key)
}
