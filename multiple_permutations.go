package cryptolabs

func NewMultiplePermutations() Cipher {
	return &multiplePermutations{'_'}
}

type multiplePermutations struct {
	dummy rune
}

func (permutations *multiplePermutations) Encrypt(text, key string) string {
	panic("implement me")
}

func (permutations *multiplePermutations) Decrypt(text, key string) string {
	panic("implement me")
}
