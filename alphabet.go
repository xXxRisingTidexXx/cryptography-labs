package cryptolabs

var DefaultAlphabet = NewAlphabet(
	"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyzАБВГҐДЕЄЖЗИІЇЙКЛМНОПРСТУФХЦЧШЩЬЮЯабвгґд" +
		"еєжзиіїйклмнопрстуфхцчшщьюя .,-!?'\"",
)

var EnglishAlphabet = NewAlphabet("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

var UkrainianAlphabet = NewAlphabet("АБВГҐДЕЄЖЗИІЇЙКЛМНОПРСТУФХЦЧШЩЬЮЯ")

func NewAlphabet(letters string) *Alphabet {
	chars := []rune(letters)
	indices := make(map[rune]int, len(chars))
	for i, char := range chars {
		indices[char] = i
	}
	return &Alphabet{chars, indices}
}

type Alphabet struct {
	chars   []rune
	indices map[rune]int
}

func (alphabet *Alphabet) Length() int {
	return len(alphabet.chars)
}

func (alphabet *Alphabet) Shift(char rune, shift int) rune {
	return alphabet.Char(alphabet.indices[char] + shift)
}

func (alphabet *Alphabet) Char(index int) rune {
	return alphabet.chars[mod(index, len(alphabet.chars))]
}

func (alphabet *Alphabet) Modulo(char1, char2 rune, direction int) rune {
	return alphabet.Shift(char1, direction*alphabet.indices[char2])
}
