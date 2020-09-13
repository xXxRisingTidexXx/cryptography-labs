package cryptolabs

var DefaultAlphabet = NewAlphabet(
	"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyzАБВГҐДЕЄЖЗИІЇЙКЛМНОПРСТУФХЦЧШЩЬЮЯабвгґдеєжзиіїйкл" +
		"мнопрстуфхцчшщьюя .,-!?",
)

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

func (alphabet *Alphabet) Shift(char rune, shift int) rune {
	index := (alphabet.indices[char] + shift) % len(alphabet.chars)
	if index < 0 {
		index += len(alphabet.chars)
	}
	return alphabet.chars[index]
}
