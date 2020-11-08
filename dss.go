package cryptolabs

func NewDSS(p, q, k, x, g int64) Signer {
	return &dss{p, q, k, x, g, pow64(g, k, p), mod64(pow64(g, x, p), q)}
}

func NewExplicitDSS(p, q, g, y, r int64) Signer {
	return &dss{p: p, q: q, g: g, y: y, r: r}
}

type dss struct {
	p int64
	q int64
	k int64
	x int64
	g int64
	y int64
	r int64
}

func (dss *dss) Encode(m, _ int64) int64 {
	_, i := gcd64(dss.x, dss.q)
	return mod64(i*mod64(m+dss.k*dss.r, dss.q), dss.q)
}

func (dss *dss) Decode(_, _ int64) int64 {
	panic("implement me")
}

func (dss *dss) Verify(s, m int64) bool {
	_, w := gcd64(s, dss.q)
	return dss.r == mod64(
		mod64(
			pow64(dss.g, mod64(m*w, dss.q), dss.p)*pow64(dss.y, mod64(dss.r*w, dss.q), dss.p),
			dss.p,
		),
		dss.q,
	)
}
