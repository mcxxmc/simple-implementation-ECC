package ecc

import "simple-implementation-ECC/galois"

// Elliptic y^2 = x^3 + a x + b over finite field P. Constructor NewElliptic().
type Elliptic struct {
	A int
	B int
	P int
	X int
	Y int
}

// NewElliptic returns a pointer to a new Elliptic curve object y^2 = x^3 + a x + b over finite field P.
func NewElliptic(a, b, p int) *Elliptic {
	return &Elliptic{
		A: a,
		B: b,
		P: p,
	}
}

// sets the generator point G (x, y). Should only be called once.
func (e *Elliptic) setGeneratorPoint(x, y int) {
	e.X, e.Y = x, y
}

// Generate returns (x`, y`) which is a k multiplication of G (x, y) on the elliptic curve e.
// Should be called after setGeneratorPoint() and k <= p.
func Generate(k int, e *Elliptic) (int, int) {
	if k > e.P {
		panic("k is too big")
	}
	return galois.MultiplyV2(e.A, e.X, e.Y, k, e.P)
}

// Calculate returns (x`, y`) which is a k multiplication of (x, y) on the elliptic curve e, k <= p.
func Calculate(x, y, k int, e *Elliptic) (int, int) {
	if k > e.P {
		panic("k is too big")
	}
	return galois.MultiplyV2(e.A, x, y, k, e.P)
}

// SampleElliptic returns a pointer to a sample Elliptic with a small p.
//
// Namely, y^2 = x^3 + 0 x + 7 (mod 17), with generator point G (15, 13)
func SampleElliptic() *Elliptic {
	return &Elliptic{
		A: 0,
		B: 7,
		P: 17,
		X: 15,
		Y: 13,
	}
}
