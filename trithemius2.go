package cryptolabs

func NewTrithemius2(alphabet *Alphabet) Cipher {
	return &trithemius2{alphabet}
}

type trithemius2 struct {
	alphabet *Alphabet
}

func (trithemius2 *trithemius2) Encrypt(text, _ string) string {
	return trithemius2.transform(text, 1)
}

func (trithemius2 *trithemius2) transform(text string, direction int) string {
	runes := []rune(text)
	newRunes := make([]rune, len(runes))
	for i := range runes {
		newRunes[i] = trithemius2.alphabet.Shift(runes[i], direction*i)
	}
	return string(newRunes)
}

func (trithemius2 *trithemius2) Decrypt(text, _ string) string {
	return trithemius2.transform(text, -1)
}
