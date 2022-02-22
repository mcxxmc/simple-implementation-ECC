package galois // Package galois is the package for finite field (or galois field)

// Point stands for a point on the 2D finite field.
//
// The attributes of a Point object should not be modified during calculation.
// Therefore, it is always passed by struct instead of by pointer.
type Point struct {
	X 		int
	Y 		int
	IsNone 	bool		// a "none" point stands for infinity
}

// NewPoint returns a pointer to a new Point object that is not none.
func NewPoint(x, y int) Point {
	return Point{X: x, Y: y, IsNone: false}
}

// NonePoint returns a pointer to a new Point object that is none.
func NonePoint() Point {
	return Point{X: 0, Y: 0, IsNone: true}
}

// PointEqual checks if the 2 points are equal.
func PointEqual(point1, point2 Point) bool {
	if point1.IsNone == point2.IsNone && point1.X == point2.X && point1.Y == point2.Y {
		return true
	}
	return false
}

// Copy returns a pointer to a deep copy of the Point object.
func Copy(p Point) Point {
	return Point{X: p.X, Y: p.Y, IsNone: p.IsNone}
}

// Mod returns a mod p which is never negative.
func Mod(a, p int) int {
	r := a % p
	if r < 0 {
		r += p
	}
	return r
}

// Inverse finds the inverse of a over the finite field p; assumes p to be a prime.
// Using Extended Euclidean Algorithm. The second output indicates if the inverse exists.
//
// CAUTION: a must be NON-NEGATIVE!
//
// Note that this functions does not check if p is a prime; the result may be unpredictable if p is not a prime!
func Inverse(a, p int) (int, bool) {
	if a < 0 {
		panic("non-positive a")
	}
	if a == 0 {
		return 0, false
	}

	if a == 1 {  // the inverse of 1 is always 1
		return 1, true
	}

	v, t, c, u := 1, 1, p % a, p / a
	for c != 1 && t == 1 {
		q := a / c
		a, v = a % c, v + q * u
		if a == 1 {
			t = 0
		}
		if t == 1 {
			q, c = c / a, c % a
			u += q * v
		}
	}
	u = v * (1 - t) + t * (p - u)

	return u, true
}

// Doubling finds the doubling point of (x, y) on the elliptic curve y^2 = x^3 + ax + b
// over the finite field p.		//todo: handling overflow
//
// Returns a New Point object.
func Doubling(point Point, a, p int) Point {
	if point.IsNone {
		return Copy(point)
	}
	x, y := point.X, point.Y
	lambda := [2]int{3 * x * x + a, 2 * y}		// nominator, denominator
	x2 := [2]int{lambda[0] * lambda[0] - 2 * x * lambda[1] * lambda[1], lambda[1] * lambda[1]}
	lambdaMulXMinusX2 := [2]int{lambda[0] * (x2[1] * x - x2[0]), lambda[1] * x2[1]}
	y2 := [2]int{lambdaMulXMinusX2[0] - y * lambdaMulXMinusX2[1], lambdaMulXMinusX2[1]}
	inverseX2, exist := Inverse(x2[1], p)
	if !exist {
		return NonePoint()
	}
	inverseY2, exist := Inverse(y2[1], p)
	if !exist {
		return NonePoint()
	}
	x, y = Mod(Mod(x2[0], p) * inverseX2, p), Mod(Mod(y2[0], p) * inverseY2, p)
	return NewPoint(x, y)
}

// Add adds (x1, y1) and (x2, y2) on the elliptic curve y^2 = x^3 + ax + b over a finite field p.
//
// Returns a New Point object.
func Add(point1, point2 Point, a, p int) Point {
	if point1.IsNone {
		return Copy(point2)
	}
	if point2.IsNone {
		return Copy(point1)
	}
	if PointEqual(point1, point2) {		// the algorithm is different when point1 == point2
		return Doubling(point1, a, p)
	}

	x1, y1 := point1.X, point1.Y
	x2, y2 := point2.X, point2.Y

	lambda := [2]int{y2 - y1, x2 - x1}
	lambda2 := [2]int{lambda[0] * lambda[0], lambda[1] * lambda[1]}
	x3 := [2]int{lambda2[0] - (x1 + x2) * lambda2[1], lambda2[1]}
	lambdaMulX1MinusX3 := [2]int{lambda[0] * (x3[1] * x1 - x3[0]), lambda[1] * x3[1]}
	y3 := [2]int{lambdaMulX1MinusX3[0] - y1 * lambdaMulX1MinusX3[1], lambdaMulX1MinusX3[1]}

	// the input to Inverse() should be non-negative
	if x3[1] < 0 {
		x3[0] *= -1
		x3[1] *= -1
	}
	if y3[1] < 0 {
		y3[0] *= -1
		y3[1] *= -1
	}

	inverseX3, exist := Inverse(x3[1], p)
	if !exist {
		return NonePoint()
	}
	inverseY3, exist := Inverse(y3[1], p)
	if !exist {
		return NonePoint()
	}

	x, y := Mod(Mod(x3[0], p) * inverseX3, p), Mod(Mod(y3[0], p) * inverseY3, p)
	return NewPoint(x, y)
}

// Multiply returns k * (x, y) on the elliptic curve y^2 = x^3 + ax + b over a finite field p.
//
// using double-add algorithm (recursive). Vulnerable to timing analysis.
//
// Returns a New Point object.
func Multiply(point Point, a, k, p int) Point {
	if point.IsNone {
		return NonePoint()
	}
	switch {
	case k < 0:
		panic("negative k, not implemented")

	case k == 0:
		return NonePoint()

	case k == 1:
		return Copy(point)

	case k % 2 == 1:
		point2 := Multiply(point, a, k - 1, p)
		return Add(point, point2, a, p)

	default:
		point2 := Doubling(point, a, p)
		return Multiply(point2, a, k / 2, p)
	}
}

// MultiplyV2 an alternative loop version of double-add algorithm. Vulnerable to timing analysis.
//
// Returns a New Point object.
func MultiplyV2(point Point, a, k, p int) Point {
	if point.IsNone {
		return NonePoint()
	}
	if k < 0 {
		panic("negative k, not implemented")
	}
	point1, point2 := Copy(point), NonePoint()
	for k != 0 {
		if k & 1 == 1 {
			point2 = Add(point1, point2, a, p)
		}
		point1 = Doubling(point1, a, p)
		k >>= 1
	}
	return point2
}
