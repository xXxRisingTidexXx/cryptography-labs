package cryptolabs

func NewSBox(table [8][16]byte) *SBox {
	sbox := &SBox{}
	for i := 0; i < 256; i++ {
		sbox.k87[i] = table[7][i>>4]<<4 | table[6][i&15]
		sbox.k65[i] = table[5][i>>4]<<4 | table[4][i&15]
		sbox.k43[i] = table[3][i>>4]<<4 | table[2][i&15]
		sbox.k21[i] = table[1][i>>4]<<4 | table[0][i&15]
	}
	return sbox
}

type SBox struct {
	k87 [256]byte
	k65 [256]byte
	k43 [256]byte
	k21 [256]byte
}

func (sbox *SBox) Get(char, subkey uint32) uint32 {
	x := char + subkey
	x = uint32(sbox.k87[x>>24&255])<<24 |
		uint32(sbox.k65[x>>16&255])<<16 |
		uint32(sbox.k43[x>>8&255])<<8 |
		uint32(sbox.k21[x&255])
	return x<<11 | x>>(32-11)
}
