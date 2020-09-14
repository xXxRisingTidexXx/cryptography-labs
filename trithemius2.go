package cryptolabs

func NewTrithemius2(alphabet *Alphabet) Cipher {
	return &trithemius2{alphabet}
}

type trithemius2 struct {
	alphabet *Alphabet
}

func (trithemius2 *trithemius2) Encrypt(text, _ string) string {
	return text
}

func (trithemius2 *trithemius2) Decrypt(text, _ string) string {
	return text
}
