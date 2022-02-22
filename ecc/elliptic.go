package ecc

import "github.com/mcxxmc/simple-implementation-ECC/galois"

// Elliptic y^2 = x^3 + a x + b over finite field P. Constructor NewElliptic().
type Elliptic struct {
	G 	galois.Point
	A 	int
	B 	int
	P  	int
}

// NewElliptic returns a pointer to a new Elliptic curve object y^2 = x^3 + a x + b over finite field P.
//
// PLEASE CALL SetGeneratorPoint() to set the generator point G!
func NewElliptic(a, b, p int) *Elliptic {
	return &Elliptic{
		A: a, B: b, P: p,
	}
}

// SetGeneratorPoint sets the generator point G (x, y). Should only be called once.
//
// CAUTION: you CANNOT set an infinite point as the generator point!
func (e *Elliptic) SetGeneratorPoint(x, y int) {
	e.G = galois.NewPoint(x, y)
}

// Generate returns (x`, y`) which is a k multiplication of G (x, y) on the elliptic curve e.
// Should be called after SetGeneratorPoint().
func Generate(k int, e *Elliptic) galois.Point {
	return galois.MultiplyV2(e.G, e.A, k, e.P)
}

// Calculate returns (x`, y`) which is a k multiplication of (x, y) on the elliptic curve e.
func Calculate(xy galois.Point, k int, e *Elliptic) galois.Point {
	return galois.MultiplyV2(xy, e.A, k, e.P)
}

// SampleElliptic returns a pointer to a sample Elliptic with a small p.
//
// Namely, y^2 = x^3 + 0 x + 7 (mod 17). A suggested generator point for testing is (15, 13).
//
// PLEASE CALL SetGeneratorPoint() to set the generator point G!
func SampleElliptic() *Elliptic {
	return &Elliptic{
		A: 0, B: 7, P: 17,
	}
}
