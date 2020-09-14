package cryptolabs

import (
	"gonum.org/v1/gonum/mat"
	"math"
)

func NewHill(alphabet *Alphabet) Cipher {
	return &hill{alphabet}
}

type hill struct {
	alphabet *Alphabet
}

func (hill *hill) Encrypt(text, key string) string {
	if text == "" || key == "" {
		return text
	}
	textRunes, keyRunes := []rune(text), []rune(key)
	modulo := int(math.Floor(math.Sqrt(float64(len(keyRunes)))))
	if modulo*modulo != len(keyRunes) || len(textRunes)%modulo != 0 {
		return text
	}
	matrix := mat.NewDense(modulo, modulo, nil)
	for i := range keyRunes {
		matrix.Set(i/modulo, i%modulo, float64(hill.alphabet.indices[keyRunes[i]]))
	}
	if gcd, _, _ := gcd(int(math.Round(mat.Det(matrix))), hill.alphabet.Length()); gcd != 1 {
		return text
	}
	newRunes := make([]rune, len(textRunes))
	for i := 0; i < len(textRunes); i += modulo {
		ngram := mat.NewVecDense(modulo, nil)
		for j := 0; j < modulo; j++ {
			ngram.SetVec(j, float64(hill.alphabet.indices[textRunes[i+j]]))
		}
		ngram.MulVec(matrix, ngram)
		for j := 0; j < modulo; j++ {
			newRunes[i+j] = hill.alphabet.Char(int(math.Round(ngram.AtVec(j))))
		}
	}
	return string(newRunes)
}

func (hill *hill) Decrypt(text, key string) string {
	if text == "" || key == "" {
		return text
	}
	textRunes, keyRunes := []rune(text), []rune(key)
	modulo := int(math.Floor(math.Sqrt(float64(len(keyRunes)))))
	if modulo*modulo != len(keyRunes) || len(textRunes)%modulo != 0 {
		return text
	}
	matrix := mat.NewDense(modulo, modulo, nil)
	for i := range keyRunes {
		matrix.Set(i/modulo, i%modulo, float64(hill.alphabet.indices[keyRunes[i]]))
	}
	det := int(math.Round(mat.Det(matrix)))
	gcd, inverse, _ := gcd(det, len(hill.alphabet.chars))
	if gcd != 1 {
		return text
	}
	_ = matrix.Inverse(matrix)
	matrix.Scale(float64(det*inverse), matrix)
	matrix.Apply(
		func(i, j int, value float64) float64 {
			return float64(mod(int(math.Round(value)), hill.alphabet.Length()))
		},
		matrix,
	)
	newRunes := make([]rune, len(textRunes))
	for i := 0; i < len(textRunes); i += modulo {
		ngram := mat.NewVecDense(modulo, nil)
		for j := 0; j < modulo; j++ {
			ngram.SetVec(j, float64(hill.alphabet.indices[textRunes[i+j]]))
		}
		ngram.MulVec(matrix, ngram)
		for j := 0; j < modulo; j++ {
			newRunes[i+j] = hill.alphabet.Char(int(math.Round(ngram.AtVec(j))))
		}
	}
	return string(newRunes)
}
