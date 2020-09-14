package cryptolabs

var DefaultAlphabet = NewAlphabet(
	"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyzАБВГҐДЕЄЖЗИІЇЙКЛМНОПРСТУФХЦЧШЩЬЮЯабвгґдеєжзиіїйкл" +
		"мнопрстуфхцчшщьюя .,-!?",
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
	return alphabet.chars[mod((alphabet.indices[char]+shift)%len(alphabet.chars), len(alphabet.chars))]
}

func (alphabet *Alphabet) Char(index int) rune {
	return alphabet.chars[mod(index, len(alphabet.chars))]
}
