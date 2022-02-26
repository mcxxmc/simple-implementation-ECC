package galois

//todo: handle overflow and enable big int (maybe the "big" standard package)
//todo: possibly introduce prime factorization to reduce the scale of nominators and denominators

// Fraction a way to note a number; it does not check gcd.
type Fraction struct {
	Nominator		int
	Denominator		int		// the divisor
}

func NewFraction(nominator, denominator int) *Fraction {
	return &Fraction{
		Nominator:   nominator,
		Denominator: denominator,
	}
}

func NewFractionFromInt(int int) *Fraction {
	return &Fraction{
		Nominator:   int,
		Denominator: 1,
	}
}

// MulInt returns a new Fraction object equal to this fraction number times integer i.
func (frac *Fraction) MulInt(i int) *Fraction {
	return &Fraction{
		Nominator:   frac.Nominator * i,
		Denominator: frac.Denominator,
	}
}

// MulFrac returns a new Fraction object equal to this fraction number times the other one.
func (frac *Fraction) MulFrac(another *Fraction) *Fraction {
	return &Fraction{
		Nominator:   frac.Nominator * another.Nominator,
		Denominator: frac.Denominator * another.Denominator,
	}
}

// PlusInt returns a new Fraction object equal to this fraction number plus integer i.
func (frac *Fraction) PlusInt(i int) *Fraction {
	return &Fraction{
		Nominator:   frac.Nominator + frac.Denominator * i,
		Denominator: frac.Denominator,
	}
}

// PlusFrac returns a new Fraction object equal to this fraction number plus the other one.
func (frac *Fraction) PlusFrac(another *Fraction) *Fraction {
	return &Fraction{
		Nominator:   frac.Nominator * another.Denominator + frac.Denominator * another.Nominator,
		Denominator: frac.Denominator * another.Denominator,
	}
}

// Switch returns a new Fraction object where both nominator and denominator are multiplied by -1.
func (frac *Fraction) Switch() *Fraction {
	return &Fraction{
		Nominator:   frac.Nominator * -1,
		Denominator: frac.Denominator * -1,
	}
}
