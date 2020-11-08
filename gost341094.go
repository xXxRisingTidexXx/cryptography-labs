package cryptolabs

func NewGOST341094(p, q, a, k, x int64) Signer {
	return &gost341094{p, q, a, k, x, pow64(a, x, p), mod64(pow64(a, k, p), q)}
}

func NewExplicitGOST341094(p, q, a, y, r int64) Signer {
	return &gost341094{p: p, q: q, a: a, y: y, r: r}
}

type gost341094 struct {
	p int64
	q int64
	a int64
	k int64
	x int64
	y int64
	r int64
}

func (gost *gost341094) Encode(m, _ int64) int64 {
	return mod64(gost.x*gost.r+gost.k*m, gost.q)
}

func (gost *gost341094) Decode(_, _ int64) int64 {
	panic("implement me")
}

func (gost *gost341094) Verify(s, m int64) bool {
	v := pow64(m, gost.q-2, gost.q)
	return gost.r == mod64(
		mod64(
			pow64(gost.a, mod64(s*v, gost.q), gost.p)*
				pow64(gost.y, mod64((gost.q-gost.r)*v, gost.q), gost.p),
			gost.p,
		),
		gost.q,
	)
}
