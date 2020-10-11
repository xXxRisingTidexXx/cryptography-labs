package cryptolabs

import (
	"encoding/binary"
	log "github.com/sirupsen/logrus"
)

func NewGOST2814789(sbox *SBox) Cipher {
	return &gost2814789{sbox}
}

type gost2814789 struct {
	sbox *SBox
}

func (gost *gost2814789) Encrypt(text, key string) string {
	if len(text)/8 == 0 || len(text)%8 != 0 || len(key) != 32 {
		log.Debugf("cryptolabs: gost2814789 got insufficient input, returned no digest")
		return text
	}
	message, subkeys, output := []byte(text), gost.subkeys(key), make([]byte, len(text))
	for i := 0; i < len(text); i += 8 {
		left := binary.LittleEndian.Uint32(message[i : i+4])
		right := binary.LittleEndian.Uint32(message[i+4 : i+8])
		for j := 0; j < 32; j++ {
			left, right = right^gost.sbox.Get(left, subkeys[j]), left
			gost.debugf("encryption", i, j, left, right)
		}
		binary.LittleEndian.PutUint32(output[i:i+4], left)
		binary.LittleEndian.PutUint32(output[i+4:i+8], right)
	}
	return string(output)
}

func (gost *gost2814789) subkeys(key string) [32]uint32 {
	secret, subkeys := []byte(key), [32]uint32{}
	for i := 0; i < 32; i += 4 {
		j := i / 4
		subkeys[j] = binary.LittleEndian.Uint32(secret[i : i+4])
		subkeys[j+8] = subkeys[j]
		subkeys[j+16] = subkeys[j]
		subkeys[31-j] = subkeys[j]
	}
	return subkeys
}

func (gost *gost2814789) debugf(process string, i, j int, left, right uint32) {
	log.Debugf(
		"cryptolabs: gost2814789 %s, block %d, round %d, left %d / %x, right %d / %x",
		process,
		i/8,
		j,
		left,
		left,
		right,
		right,
	)
}

func (gost *gost2814789) Decrypt(text, key string) string {
	if len(text)/8 == 0 || len(text)%8 != 0 || len(key) != 32 {
		log.Debugf("cryptolabs: gost2814789 got insufficient input, returned no message")
		return text
	}
	message, subkeys, output := []byte(text), gost.subkeys(key), make([]byte, len(text))
	for i := 0; i < len(text); i += 8 {
		left := binary.LittleEndian.Uint32(message[i : i+4])
		right := binary.LittleEndian.Uint32(message[i+4 : i+8])
		for j := 31; j >= 0; j-- {
			left, right = right, left^gost.sbox.Get(right, subkeys[j])
			gost.debugf("decryption", i, j, left, right)
		}
		binary.LittleEndian.PutUint32(output[i:i+4], left)
		binary.LittleEndian.PutUint32(output[i+4:i+8], right)
	}
	return string(output)
}
