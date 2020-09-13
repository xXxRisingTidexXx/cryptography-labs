package cryptolabs

var Hill = NewHill(DefaultAlphabet)

func NewHill(alphabet *Alphabet) Cipher {
	return &hill{alphabet}
}

type hill struct {
	alphabet *Alphabet
}

func (hill *hill) Encrypt(text, key string) string {
	return text
}

func (hill *hill) Decrypt(text, key string) string {
	return text
}
