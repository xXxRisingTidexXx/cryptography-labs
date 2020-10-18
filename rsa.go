package cryptolabs

func NewRSA(p, q, e int64) Cipher {
	return &rsa{p * q, e, inverse(e, (p-1)*(q-1))}
}

type rsa struct {
	n        int64
	e        int64
	d        int64
}

func (rsa *rsa) Encrypt(text string, _ string) string {
	plain := int64(0)
	for _, char := range []byte(text) {
		plain = 26 * plain + int64(char - 25)
	}
	return text
}

func (rsa *rsa) Decrypt(text string, _ string) string {
	return text
}
