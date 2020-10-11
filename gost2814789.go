package cryptolabs

func NewGOST2814789(sbox SBox) Cipher {
	return &gost2814789{sbox}
}

type gost2814789 struct {
	sbox SBox
}

func (gost *gost2814789) Encrypt(text, key string) string {
	panic("implement me")
}

func (gost *gost2814789) Decrypt(text, key string) string {
	panic("implement me")
}
