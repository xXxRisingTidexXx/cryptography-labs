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

func gcd64(a, b int64) (int64, int64) {
	gcd, x, _ := solve64(mod64(a, b), b)
	return gcd, mod64(x, b)
}

func mod64(a, b int64) int64 {
	return (a%b + b) % b
}

func solve64(a, b int64) (int64, int64, int64) {
	if a == 0 {
		return b, 0, 1
	}
	gcd, x, y := solve64(b%a, a)
	return gcd, y - (b/a)*x, x
}

func pow64(a, b, c int64) int64 {
	if c == 1 {
		return 0
	}
	x := int64(1)
	for i := int64(0); i < b; i++ {
		x = mod64(x*a, c)
	}
	return x
}

func gcdRune(a, b rune) (rune, rune) {
	gcd, x, _ := solveRune(modRune(a, b), b)
	return gcd, modRune(x, b)
}

func modRune(a, b rune) rune {
	return (a%b + b) % b
}

func solveRune(a, b rune) (rune, rune, rune) {
	if a == 0 {
		return b, 0, 1
	}
	gcd, x, y := solveRune(b%a, a)
	return gcd, y - (b/a)*x, x
}

func powRune(a, b, c rune) rune {
	if c == 1 {
		return 0
	}
	x := rune(1)
	for i := rune(0); i < b; i++ {
		x = modRune(x*a, c)
	}
	return x
}
