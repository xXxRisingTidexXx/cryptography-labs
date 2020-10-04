package cryptolabs

func Magma(text, key rune) rune {
	return text + key
}

func DES(text, key rune) rune {
	return text ^ key
}

func Hypot(text, key rune) rune {
	return text*text + key*key
}
