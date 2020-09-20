package cryptolabs

import (
	"math"
)

func mod(a, b int) int {
	return (a%b + b) % b
}

func gcd(a, b int) (int, int) {
	gcd, x, _ := rawGCD(mod(a, b), b)
	return gcd, mod(x, b)
}

func rawGCD(a, b int) (int, int, int) {
	if a == 0 {
		return b, 0, 1
	}
	gcd, x, y := rawGCD(b%a, a)
	return gcd, y - (b/a)*x, x
}

func sqrt(a int) (int, bool) {
	b := int(math.Floor(math.Sqrt(float64(a))))
	return b, b*b == a
}
