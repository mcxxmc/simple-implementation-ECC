package test

import (
	"github.com/mcxxmc/simple-implementation-ECC/ecc"
	"testing"
)

func TestECC(t *testing.T) {
	ep := ecc.SampleElliptic()
	ep.SetGeneratorPoint(15, 13)
	if r := ecc.Generate(2, ep); r.X != 2 || r.Y != 10 {
		t.Error("wrong generation, expecting (", 2, ", ", 10, "), got ", r)
	}
	if r := ecc.Generate(3, ep); r.X != 8 || r.Y != 3 {
		t.Error("wrong generation, expecting (", 8, ", ", 3, "), got ", r)
	}
	if r := ecc.Generate(6, ep); r.X != 5 || r.Y != 8 {
		t.Error("wrong generation, expecting (", 5, ", ", 8, "), got ", r)
	}
	if r := ecc.Generate(14, ep); r.X != 12 || r.Y != 16 {
		t.Error("wrong generation, expecting (", 12, ", ", 16, "), got ", r)
	}
	if r := ecc.Generate(15, ep); r.X != 8 || r.Y != 14 {
		t.Error("wrong generation, expecting (", 8, ", ", 14, "), got ", r)
	}
}
