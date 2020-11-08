package cryptolabs

func NewDSS(p, q, k, x, g int64) Signer {
	return &dss{}
}

type dss struct {

}

func (dss *dss) Encode(m, _ int64) int64 {
	return 0
}

func (dss *dss) Decode(s, _ int64) int64 {
	return 0
}

