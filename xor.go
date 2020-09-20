package cryptolabs

func NewXOR() Cipher {
	return &xor{}
}

type xor struct{}

func (xor *xor) Encrypt(text, key string) string {
	return text
}

func (xor *xor) Decrypt(text, key string) string {
	return text
}
