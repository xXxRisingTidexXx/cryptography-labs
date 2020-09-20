package cryptolabs

func NewMagicSquare() Cipher {
	return &magicSquare{}
}

type magicSquare struct{}

func (square *magicSquare) Encrypt(text string, key string) string {
	return text
}

func (square *magicSquare) Decrypt(text string, key string) string {
	return text
}
