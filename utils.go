package cryptolabs

import (
	"math"
)

func mod(a, b int) int {
	return (a%b + b) % b
}

func gcd(a, b int) (int, int) {
	gcd, x, _ := solve(mod(a, b), b)
	return gcd, mod(x, b)
}

func solve(a, b int) (int, int, int) {
	if a == 0 {
		return b, 0, 1
	}
	gcd, x, y := solve(b%a, a)
	return gcd, y - (b/a)*x, x
}

func sqrt(a int) (int, bool) {
	b := int(math.Floor(math.Sqrt(float64(a))))
	return b, b*b == a
}

func inverse(a, b int64) int64 {
	_, x, _ := euclidean((a%b+b)%b, b)
	return (x%b + b) % b
}

func euclidean(a, b int64) (int64, int64, int64) {
	if a == 0 {
		return b, 0, 1
	}
	gcd, x, y := euclidean(b%a, a)
	return gcd, y - (b/a)*x, x
}
