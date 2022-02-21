package ecc

import "simple-implementation-ECC/galois"

// Elliptic y^2 = x^3 + a x + b over finite field P. Constructor NewElliptic().
type Elliptic struct {
	G 	[2]int
	A 	int
	B 	int
	P  	int
}

// NewElliptic returns a pointer to a new Elliptic curve object y^2 = x^3 + a x + b over finite field P.
func NewElliptic(a, b, p int) *Elliptic {
	return &Elliptic{
		G: [2]int{}, A: a, B: b, P: p,
	}
}

// SetGeneratorPoint sets the generator point G (x, y). Should only be called once.
func (e *Elliptic) SetGeneratorPoint(g [2]int) {
	e.G = g
}

// Generate returns (x`, y`) which is a k multiplication of G (x, y) on the elliptic curve e.
// Should be called after SetGeneratorPoint().
func Generate(k int, e *Elliptic) [2]int {
	x, y := galois.MultiplyV2(e.A, e.G[0], e.G[1], k, e.P)
	return [2]int{x, y}
}

// Calculate returns (x`, y`) which is a k multiplication of (x, y) on the elliptic curve e.
func Calculate(xy [2]int, k int, e *Elliptic) [2]int {
	x, y := galois.MultiplyV2(e.A, xy[0], xy[1], k, e.P)
	return [2]int{x, y}
}

// SampleElliptic returns a pointer to a sample Elliptic with a small p.
//
// Namely, y^2 = x^3 + 0 x + 7 (mod 17), with generator point G (15, 13)
func SampleElliptic() *Elliptic {
	return &Elliptic{
		G: [2]int{}, A: 0, B: 7, P: 17,
	}
}
