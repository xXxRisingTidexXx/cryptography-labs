package cryptolabs

type Signer interface {
	Encode(int64, int64) int64
	Decode(int64, int64) int64
	Verify(int64, int64) bool
}
