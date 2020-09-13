package cryptolabs

var DefaultAlphabet = NewAlphabet(
	"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyzАБВГҐДЕЄЖЗИІЇЙКЛМНОПРСТУФХЦЧШЩЬЮЯабвгґдеєжзиіїйкл" +
		"мнопрстуфхцчшщьюя .,-!?",
)

func NewAlphabet(chars string) *Alphabet {
	runes := []rune(chars)
	indices := make(map[rune]int, len(runes))
	for i, char := range runes {
		indices[char] = i
	}
	return &Alphabet{runes, indices}
}

type Alphabet struct {
	runes   []rune
	indices map[rune]int
}

func (alphabet *Alphabet) Get(char rune, shift int) rune {
	return alphabet.runes[(alphabet.indices[char]+shift)%len(alphabet.runes)]
}
