package cryptolabs

func NewNModuloAddition(alphabet *Alphabet) Cipher {
	return &nModuloAddition{alphabet}
}

type nModuloAddition struct {
	alphabet *Alphabet
}

func (addition *nModuloAddition) Encrypt(text, key string) string {
	return addition.transform(text, key, 1)
}

func (addition *nModuloAddition) transform(text, key string, direction int) string {
	textRunes, keyRunes := []rune(text), []rune(key)
	newRunes := make([]rune, len(textRunes))
	for i := range textRunes {
		newRunes[i] = addition.alphabet.Modulo(textRunes[i], keyRunes[i%len(keyRunes)], direction)
	}
	return string(newRunes)
}

func (addition *nModuloAddition) Decrypt(text, key string) string {
	return addition.transform(text, key, -1)
}
