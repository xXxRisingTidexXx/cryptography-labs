package cryptolabs

func NewElgamal(p, q, x, k rune) Cipher {
	return &elgamal{p, q, x, powRune(q, x, p), k}
}

type elgamal struct {
	p rune
	q rune
	x rune
	y rune
	k rune
}

func (elgamal *elgamal) Encrypt(text, _ string) string {
	runes := []rune(text)
	output := make([]rune, 2*len(runes))
	for i := range runes {
		output[i] = powRune(elgamal.q, elgamal.k, elgamal.p)
		output[i+len(runes)] = (powRune(elgamal.y, elgamal.k, elgamal.p) *
			(runes[i] % elgamal.p)) % elgamal.p
	}
	return string(output)
}

func (elgamal *elgamal) Decrypt(text, _ string) string {
	runes := []rune(text)
	output := make([]rune, len(runes)/2)
	for i := range output {
		output[i] = ((runes[i+len(output)] % elgamal.p) *
			powRune(runes[i], elgamal.p-1-elgamal.x, elgamal.p)) % elgamal.p
	}
	return string(output)
}
