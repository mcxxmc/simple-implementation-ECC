package test

import (
	"simple-implementation-ECC/ecc"
	"testing"
)

func TestECC(t *testing.T) {
	ep := ecc.SampleElliptic()
	if r1, r2 := ecc.Generate(2, ep); r1 != 2 || r2 != 10 {
		t.Error("wrong generation, expecting (", 2, ", ", 10, "), got (", r1, ", ", r2, ")")
	}
	if r1, r2 := ecc.Generate(3, ep); r1 != 8 || r2 != 3 {
		t.Error("wrong generation, expecting (", 8, ", ", 3, "), got (", r1, ", ", r2, ")")
	}
	if r1, r2 := ecc.Generate(6, ep); r1 != 5 || r2 != 8 {
		t.Error("wrong generation, expecting (", 5, ", ", 8, "), got (", r1, ", ", r2, ")")
	}
	if r1, r2 := ecc.Generate(14, ep); r1 != 12 || r2 != 16 {
		t.Error("wrong generation, expecting (", 12, ", ", 16, "), got (", r1, ", ", r2, ")")
	}
	if r1, r2 := ecc.Generate(15, ep); r1 != 8 || r2 != 14 {
		t.Error("wrong generation, expecting (", 8, ", ", 14, "), got (", r1, ", ", r2, ")")
	}
}
