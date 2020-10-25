package numtheo

import "errors"

// ExtGCD solve an equation: "a*x + b*y = Gcd(a, b)"
// (x, y) satisfies that
// "|x|+|y| is minimized" and "if there are some such ones, x <= y".
// g is Gcd(a, b)
func ExtGCD(a, b int) (g, x, y int) {
	if b == 0 {
		return a, 1, 0
	}

	g, s, t := ExtGCD(b, a%b)

	return g, t, s - (a/b)*t
}

// Gcd returns the Greatest Common Divisor of two natural numbers.
// Gcd only accepts two natural numbers (a, b >= 0).
// Negative number causes panic.
// Gcd uses the Euclidean Algorithm.
func Gcd(a, b int) int {
	if a < 0 || b < 0 {
		panic(errors.New("[argument error]: Gcd only accepts two NATURAL numbers"))
	}

	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

// Lcm returns the Least Common Multiple of two natural numbers.
// Lcd only accepts two natural numbers (a, b >= 0).
// Negative number causes panic.
// Lcd uses the Euclidean Algorithm indirectly.
func Lcm(a, b int) int {
	if a < 0 || b < 0 {
		panic(errors.New("[argument error]: Gcd only accepts two NATURAL numbers"))
	}

	if a == 0 || b == 0 {
		return 0
	}

	// a = a'*gcd, b = b'*gcd, a*b = a'*b'*gcd^2
	// a' and b' are relatively prime numbers
	// gcd consists of prime numbers, that are included in a and b
	g := Gcd(a, b)

	// not (a * b / gcd), because of reducing a probability of overflow
	return (a / g) * b
}