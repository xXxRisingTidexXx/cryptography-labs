package cryptolabs

type Cipher interface {
	Encrypt(string, string) string
	Decrypt(string, string) string
}
