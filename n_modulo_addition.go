package cryptolabs

func NewNModuloAddition(alphabet *Alphabet) Cipher {
	return &nModuloAddition{alphabet}
}

type nModuloAddition struct {
	alphabet *Alphabet
}

func (addition *nModuloAddition) Encrypt(text, key string) string {
	panic("implement me")
}

func (addition *nModuloAddition) Decrypt(text, key string) string {
	panic("implement me")
}
