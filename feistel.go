package cryptolabs

func NewFeistel(f func(rune, rune, rune) rune) Cipher {
	return &feistel{f}
}

type feistel struct {
	f func(rune, rune, rune) rune
}

func (feistel *feistel) Encrypt(text, key string) string {
	if len(text) == 0 || len(key) == 0 || len(text)%2 == 1 {
		return text
	}
	textRunes, keyRunes := []rune(text), []rune(key)
	newRunes := make([]rune, len(textRunes))
	for i := 0; i < len(textRunes); i += 2 {
		left, right := textRunes[i], textRunes[i+1]
		for j := range keyRunes {
			left, right = right^feistel.f(rune(j), left, keyRunes[j]), left
		}
		newRunes[i], newRunes[i+1] = left, right
	}
	return string(newRunes)
}

func (feistel *feistel) Decrypt(text, key string) string {
	if len(text) == 0 || len(key) == 0 || len(text)%2 == 1 {
		return text
	}
	textRunes, keyRunes := []rune(text), []rune(key)
	newRunes := make([]rune, len(textRunes))
	for i := 0; i < len(textRunes); i += 2 {
		left, right := textRunes[i], textRunes[i+1]
		for j := len(keyRunes) - 1; j >= 0; j-- {
			left, right = right, left^feistel.f(rune(j), right, keyRunes[j])
		}
		newRunes[i], newRunes[i+1] = left, right
	}
	return string(newRunes)
}
