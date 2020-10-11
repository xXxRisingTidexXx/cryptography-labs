package cryptolabs

func NewSBox(bytes [8][16]byte) *SBox {
	return &SBox{}
}

type SBox struct {

}

func (sbox *SBox) Get(char, subkey uint32) uint32 {
	return 0
}
