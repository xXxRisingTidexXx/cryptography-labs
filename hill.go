package cryptolabs

import (
	"gonum.org/v1/gonum/mat"
)

var Hill = NewHill(DefaultAlphabet)

func NewHill(alphabet *Alphabet) Cipher {
	return &hill{alphabet}
}

type hill struct {
	alphabet *Alphabet
}

func (hill *hill) Encrypt(text, key string) string {
	_ = mat.NewDense(2, 2, []float64{1, 1, 1, 1})
	return text
}

func (hill *hill) Decrypt(text, key string) string {
	return text
}
