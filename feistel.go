package cryptolabs

func NewFeistel() Cipher {
	return &feistel{}
}

type feistel struct{}

func (feistel *feistel) Encrypt(text, key string) string {
	return text
}

func (feistel *feistel) Decrypt(text, key string) string {
	return text
}
