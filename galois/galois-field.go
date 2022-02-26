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

	lambda := NewFraction(3 * point.X * point.X + a, 2 * point.Y)
	fx := lambda.MulFrac(lambda).PlusInt(-2 * point.X)
	fy := lambda.MulFrac(fx.PlusInt(-1 * point.X)).MulInt(-1).PlusInt(-1 * point.Y)
	inverseX, exist := Inverse(fx.Denominator, p)
	if !exist {
		return NonePoint()
	}
	inverseY, exist := Inverse(fy.Denominator, p)
	if !exist {
		return NonePoint()
	}
	x, y := Mod(Mod(fx.Nominator, p) * inverseX, p), Mod(Mod(fy.Nominator, p) * inverseY, p)
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

	lambda := NewFraction(point2.Y - point1.Y, point2.X - point1.X)
	fx := lambda.MulFrac(lambda).PlusInt(-1 * (point1.X + point2.X))
	fy := lambda.MulFrac(fx.PlusInt(-1 * point1.X)).MulInt(-1).PlusInt(-1 * point1.Y)
	if fx.Denominator < 0 {
		fx = fx.Switch()
	}
	if fy.Denominator < 0 {
		fy = fy.Switch()
	}
	inverseX, exist := Inverse(fx.Denominator, p)
	if !exist {
		return NonePoint()
	}
	inverseY, exist := Inverse(fy.Denominator, p)
	if !exist {
		return NonePoint()
	}
	x, y := Mod(Mod(fx.Nominator, p) * inverseX, p), Mod(Mod(fy.Nominator, p) * inverseY, p)
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
