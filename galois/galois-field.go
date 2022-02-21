package galois // Package galois is the package for finite field (or galois field)

// Mod returns a mod p which is never negative.
func Mod(a, p int) int {
	r := a % p
	if r < 0 {
		r += p
	}
	return r
}

// Inverse finds the inverse of a over the finite field p; assumes p to be a prime.
// Using Extended Euclidean Algorithm.
//
// CAUTION: a must be POSITIVE!
//
// Note that this functions does not check if p is a prime; the result may be unpredictable if p is not a prime!
func Inverse(a, p int) int {
	if a < 0 {
		panic("non-positive a")
	}
	if a == 0 {		// todo: though 0 works here, it should in fact return "none"
		return 0
	}

	if a == 1 {  // the inverse of 1 is always 1
		return 1
	}

	v := 1
	t := 1
	c, u := p % a, p / a
	for c != 1 && t == 1 {
		q := a / c
		a = a % c
		v += q * u
		if a == 1 {
			t = 0
		}
		if t == 1 {
			q, c = c / a, c % a
			u += q * v
		}
	}
	u = v * (1 - t) + t * (p - u)

	return u
}

// Doubling finds the doubling point of (x, y) on the elliptic curve y^2 = x^3 + ax + b
// over the finite field p.		//todo: handling overflow
func Doubling(a, x, y, p int) (int, int) {
	lambda := [2]int{3 * x * x + a, 2 * y}		// nominator, denominator
	x2 := [2]int{lambda[0] * lambda[0] - 2 * x * lambda[1] * lambda[1], lambda[1] * lambda[1]}
	lambdaMulXMinusX2 := [2]int{lambda[0] * (x2[1] * x - x2[0]), lambda[1] * x2[1]}
	y2 := [2]int{lambdaMulXMinusX2[0] - y * lambdaMulXMinusX2[1], lambdaMulXMinusX2[1]}
	return Mod(Mod(x2[0], p) * Inverse(x2[1], p), p), Mod(Mod(y2[0], p) * Inverse(y2[1], p), p)
}

// Add adds (x1, y1) and (x2, y2) on the elliptic curve y^2 = x^3 + ax + b over a finite field p.
func Add(x1, y1, x2, y2, p int) (int, int) {
	// todo: the 2 ifs are in fact wrong: currently I assume that (none, none) is (0, 0); it will cause problem when
	// todo: (0, 0) is actually a point on the curve!
	// todo: change to a point structure using bool to mark if it is none!
	if x1 == 0 && y1 == 0 {
		return x2, y2
	}
	if x2 == 0 && y2 == 0 {
		return x1, y1
	}

	lambda := [2]int{y2 - y1, x2 - x1}
	lambda2 := [2]int{lambda[0] * lambda[0], lambda[1] * lambda[1]}
	x3 := [2]int{lambda2[0] - (x1 + x2) * lambda2[1], lambda2[1]}
	lambdaMulX1MinusX3 := [2]int{lambda[0] * (x3[1] * x1 - x3[0]), lambda[1] * x3[1]}
	y3 := [2]int{lambdaMulX1MinusX3[0] - y1 * lambdaMulX1MinusX3[1], lambdaMulX1MinusX3[1]}

	if x3[1] < 0 {
		x3[0] *= -1
		x3[1] *= -1
	}
	if y3[1] < 0 {
		y3[0] *= -1
		y3[1] *= -1
	}

	return Mod(Mod(x3[0], p) * Inverse(x3[1], p), p), Mod(Mod(y3[0], p) * Inverse(y3[1], p), p)
}

// Multiply returns k * (x, y) on the elliptic curve y^2 = x^3 + ax + b over a finite field p.
//
// using double-add algorithm (recursive). Vulnerable to timing analysis.
func Multiply(a, x, y, k, p int) (int, int) {
	switch {
	case k < 0:
		panic("negative k, not implemented")

	case k == 0:
		return 0, 0

	case k == 1:
		return x, y

	case k % 2 == 1:
		tx, ty := Multiply(a, x, y, k - 1, p)
		return Add(x, y, tx, ty, p)

	default:
		tx, ty := Doubling(a, x, y, p)
		return Multiply(a, tx, ty, k / 2, p)
	}
}

// MultiplyV2 an alternative loop version of double-add algorithm. Vulnerable to timing analysis.
func MultiplyV2(a, x, y, k, p int) (int, int) {
	if k < 0 {
		panic("negative k, not implemented")
	}
	tx1, ty1, tx2, ty2, initialized := x, y, 0, 0, false
	for k != 0 {
		if k & 1 == 1 {
			if initialized {
				tx2, ty2 = Add(tx1, ty1, tx2, ty2, p)
			} else {
				tx2, ty2, initialized = tx1, ty1, true
			}
		}
		tx1, ty1 = Doubling(a, tx1, ty1, p)
		k >>= 1
	}
	return tx2, ty2
}
