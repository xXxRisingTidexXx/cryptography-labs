package cryptolabs

func NewMerkleHellman(w [8]rune, q, r rune) Cipher {
	b := [8]rune{}
	for i := range w {
		b[i] = (w[i] * r) % q
	}
	_, i := gcdRune(r, q)
	return &merkleHellman{w, q, r, b, i}
}

type merkleHellman struct {
	w [8]rune
	q rune
	r rune
	b [8]rune
	i rune
}

func (merkleHellman *merkleHellman) Encrypt(text, _ string) string {
	bytes := []byte(text)
	output := make([]rune, len(bytes))
	for i := range bytes {
		bits := bytes[i]
		for j := 7; j >= 0; j-- {
			output[i] += merkleHellman.b[j] * rune(bits%2)
			bits >>= 1
		}
	}
	return string(output)
}

func (merkleHellman *merkleHellman) Decrypt(text, _ string) string {
	runes := []rune(text)
	output := make([]byte, len(runes))
	for i := range runes {
		sum := (runes[i] * merkleHellman.i) % merkleHellman.q
		for j := 7; j >= 0; j-- {
			if sum >= merkleHellman.w[j] {
				output[i] += 1 << (7 - j)
				sum -= merkleHellman.w[j]
			}
		}
	}
	return string(output)
}
