package cryptolabs

import (
	"strconv"
	"strings"
)

func NewMultiplePermutations() Cipher {
	return &multiplePermutations{'_'}
}

type multiplePermutations struct {
	dummy rune
}

func (permutations *multiplePermutations) Encrypt(text, key string) string {
	textRunes := []rune(text)
	keyStrings := strings.Split(key, " ")
	if len(keyStrings) != 2 {
		return text
	}
	rows, err := permutations.split(keyStrings[0])
	if err != nil {
		return text
	}
	columns, err := permutations.split(keyStrings[1])
	if err != nil {
		return text
	}
	dimension := len(rows) * len(columns)
	mappings := make([]int, dimension)
	for i := range rows {
		for j := range columns {
			mappings[i*len(columns)+j] = columns[j] + len(columns)*rows[i]
		}
	}
	length := len(textRunes)
	if modulo := length % dimension; modulo != 0 {
		length += dimension - modulo
	}
	newRunes := make([]rune, length)
	for i := range newRunes {
		newRunes[i] = permutations.dummy
	}
	for i := range textRunes {
		modulo := i % dimension
		newRunes[i-modulo+mappings[modulo]] = textRunes[i]
	}
	return string(newRunes)
}

func (permutations *multiplePermutations) split(text string) ([]int, error) {
	numbers := strings.Split(text, ",")
	indices := make([]int, len(numbers))
	for i := range numbers {
		value, err := strconv.ParseInt(numbers[i], 10, 64)
		if err != nil {
			return nil, err
		}
		indices[i] = int(value) - 1
	}
	return indices, nil
}

func (permutations *multiplePermutations) Decrypt(text, key string) string {
	textRunes := []rune(text)
	keyStrings := strings.Split(key, " ")
	if len(keyStrings) != 2 {
		return text
	}
	rows, err := permutations.split(keyStrings[0])
	if err != nil {
		return text
	}
	columns, err := permutations.split(keyStrings[1])
	if err != nil {
		return text
	}
	dimension := len(rows) * len(columns)
	mappings := make([]int, dimension)
	for i := range rows {
		for j := range columns {
			mappings[columns[j]+len(columns)*rows[i]] = i*len(columns) + j
		}
	}
	newRunes := make([]rune, len(textRunes))
	for i := range textRunes {
		modulo := i % dimension
		newRunes[i-modulo+mappings[modulo]] = textRunes[i]
	}
	return strings.ReplaceAll(string(newRunes), string(permutations.dummy), "")
}
