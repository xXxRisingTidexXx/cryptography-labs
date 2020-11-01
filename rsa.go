package cryptolabs

func NewRSA(p, q, e int64) Cipher {
	_, d := gcd64(e, (p-1)*(q-1))
	return &rsa{p * q, e, d}
}

func NewSignerRSA(n int64) Signer {
	return &rsa{n: n}
}

type rsa struct {
	n int64
	e int64
	d int64
}

func (rsa *rsa) Encrypt(text, _ string) string {
	return rsa.transform(text, rsa.e)
}

func (rsa *rsa) transform(text string, key int64) string {
	plain, bytes := int64(0), []byte(text)
	for i := len(bytes) - 1; i >= 0; i-- {
		plain = 26*plain + mod64(int64(bytes[i]-65), 26)
	}
	cipher, output := pow64(plain, key, rsa.n), make([]byte, 0)
	for ; cipher != 0; {
		output = append(output, byte(cipher%26+65))
		cipher /= 26
	}
	return string(output)
}

func (rsa *rsa) Decrypt(text, _ string) string {
	return rsa.transform(text, rsa.d)
}

func (rsa *rsa) Sign(m int64, d int64) int64 {
	return pow64(m, d, rsa.n)
}

func (rsa *rsa) Verify(s int64, e int64) bool {
	return s == pow64(s, e, rsa.n)
}
