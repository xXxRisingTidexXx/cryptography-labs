package cryptolabs

func Magma(text, key rune) rune {
	return text + key
}

func DES(text, key rune) rune {
	return text ^ key
}

func AND(text, key rune) rune {
	return text & key
}
