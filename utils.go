package cryptolabs

func mod(a, b int) int {
	return (a%b + b) % b
}

func gcd(a, b int) (int, int, int) {
	gcd, x, y := rawGCD(mod(a, b), b)
	return gcd, mod(x, b), mod(y, b)
}

func rawGCD(a, b int) (int, int, int) {
	if a == 0 {
		return b, 0, 1
	}
	gcd, x, y := rawGCD(b%a, a)
	return gcd, y - (b/a)*x, x
}
