package cryptolabs

type Signer interface {
	Sign(int64, int64) int64
	Verify(int64, int64) bool
}
