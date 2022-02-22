package test

import (
	"github.com/mcxxmc/simple-implementation-ECC/galois"
	"testing"
)

func TestInverseFiniteFieldP(t *testing.T) {
	if r, _ := galois.Inverse(322, 701); r != 455 {
		t.Error("wrong inverse; expecting ", 455, ", got ", r)
	}
}

func TestDoubling(t *testing.T) {
	if r := galois.Doubling(galois.NewPoint(15, 13), 0, 17); r.X != 2 || r.Y != 10 {
		t.Error("wrong doubling, expecting (", 2, ", ", 10, "), got (", r.X, ", ", r.Y, ")")
	}
	if r := galois.Doubling(galois.NewPoint(2, 10), 0, 17); r.X != 12 || r.Y != 1 {
		t.Error("wrong doubling, expecting (", 12, ", ", 1, "), got (", r.X, ", ", r.Y, ")")
	}
	if r := galois.Doubling(galois.NewPoint(8, 3), 0, 17); r.X != 5 || r.Y != 8 {
		t.Error("wrong doubling, expecting (", 5, ", ", 8, "), got (", r.X, ", ", r.Y, ")")
	}
	if r := galois.Doubling(galois.NewPoint(1, 12), 0, 17); r.X != 2 || r.Y != 7 {
		t.Error("wrong doubling, expecting (", 2, ", ", 7, "), got (", r.X, ", ", r.Y, ")")
	}
	if r := galois.Doubling(galois.NewPoint(5, 9), 0, 17); r.X != 5 || r.Y != 8 {
		t.Error("wrong doubling, expecting (", 5, ", ", 8, "), got (", r.X, ", ", r.Y, ")")
	}
	if r := galois.Doubling(galois.NewPoint(8, 14), 0, 17); r.X != 5 || r.Y != 9 {
		t.Error("wrong doubling, expecting (", 5, ", ", 9, "), got (", r.X, ", ", r.Y, ")")
	}
}

func TestAdd(t *testing.T) {
	if r := galois.Add(galois.NewPoint(15, 13), galois.NewPoint(2, 10), 0, 17); r.X != 8 || r.Y != 3 {
		t.Error("wrong addition, expecting (", 8, ", ", 3, "), got (", r.X, ", ", r.Y, ")")
	}
	if r := galois.Add(galois.NewPoint(5, 8), galois.NewPoint(10, 15), 0, 17); r.X != 6 || r.Y != 11 {
		t.Error("wrong addition, expecting (", 6, ", ", 11, "), got (", r.X, ", ", r.Y, ")")
	}
	if r := galois.Add(galois.NewPoint(5, 8), galois.NewPoint(5, 9), 0, 17); r.X != 0 || r.Y != 0 || !r.IsNone {
		t.Error("wrong addition, expecting (", 0, ", ", 0, ", ", true, "), got (", r.X, ", ", r.Y, ", ", r.IsNone, ")")
	}
	if r := galois.Add(galois.NewPoint(5, 8), galois.NewPoint(8, 14), 0, 17); r.X != 8 || r.Y != 3 {
		t.Error("wrong addition, expecting (", 8, ", ", 3, "), got (", r.X, ", ", r.Y, ")")
	}
	if r := galois.Add(galois.NonePoint(), galois.NewPoint(5, 9), 0, 17); r.X != 5 || r.Y != 9 || r.IsNone {
		t.Error("wrong addition, expecting (", 5, ", ", 9, ", ", false, "), got (", r.X, ", ", r.Y, ", ", r.IsNone, ")")
	}
	if r := galois.Add(galois.NewPoint(5, 9), galois.NewPoint(5, 9), 0, 17); r.X != 5 || r.Y != 8 {
		t.Error("wrong addition, expecting (", 5, ", ", 8, "), got (", r.X, ", ", r.Y, ")")
	}
}

func TestMultiple(t *testing.T) {
	if r := galois.Multiply(galois.NewPoint(15, 13), 0, 2, 17); r.X != 2 || r.Y != 10 {
		t.Error("wrong multiplying, expecting (", 2, ", ", 10, "), got (", r.X, ", ", r.Y, ")")
	}
	if r := galois.Multiply(galois.NewPoint(15, 13), 0, 3, 17); r.X != 8 || r.Y != 3 {
		t.Error("wrong multiplying, expecting (", 8, ", ", 3, "), got (", r.X, ", ", r.Y, ")")
	}
	if r := galois.Multiply(galois.NewPoint(15, 13), 0, 6, 17); r.X != 5 || r.Y != 8 {
		t.Error("wrong multiplying, expecting (", 5, ", ", 8, "), got (", r.X, ", ", r.Y, ")")
	}
	if r := galois.Multiply(galois.NewPoint(5, 8), 0, 2, 17); r.X != 5 || r.Y != 9 {
		t.Error("wrong multiplying, expecting (", 5, ", ", 9, "), got (", r.X, ", ", r.Y, ")")
	}
	if r := galois.Multiply(galois.NewPoint(8, 3), 0, 5, 17); r.X != 8 || r.Y != 14 {
		t.Error("wrong multiplying, expecting (", 8, ", ", 14, "), got (", r.X, ", ", r.Y, ")")
	}
	if r := galois.Multiply(galois.NewPoint(15, 13), 0, 20, 17); r.X != 2 || r.Y != 10 {
		t.Error("wrong multiplying, expecting (", 2, ", ", 10, "), got (", r.X, ", ", r.Y, ")")
	}
	if r := galois.Multiply(galois.NewPoint(15, 13), 0, 21, 17); r.X != 8 || r.Y != 3 {
		t.Error("wrong multiplying, expecting (", 8, ", ", 3, "), got (", r.X, ", ", r.Y, ")")
	}
	if r := galois.Multiply(galois.NewPoint(5, 9), 0, 7, 17); r.X != 5 || r.Y != 9 {
		t.Error("wrong multiplying, expecting (", 5, ", ", 9, "), got (", r.X, ", ", r.Y, ")")
	}
	if r := galois.Multiply(galois.NewPoint(5, 9), 0, 5, 17); r.X != 5 || r.Y != 8 {
		t.Error("wrong multiplying, expecting (", 5, ", ", 8, "), got (", r.X, ", ", r.Y, ")")
	}

	// testing with v2
	if r := galois.MultiplyV2(galois.NewPoint(15, 13), 0, 2, 17); r.X != 2 || r.Y != 10 {
		t.Error("wrong multiplying, expecting (", 2, ", ", 10, "), got (", r.X, ", ", r.Y, ")")
	}
	if r := galois.MultiplyV2(galois.NewPoint(15, 13), 0, 3, 17); r.X != 8 || r.Y != 3 {
		t.Error("wrong multiplying, expecting (", 8, ", ", 3, "), got (", r.X, ", ", r.Y, ")")
	}
	if r := galois.MultiplyV2(galois.NewPoint(15, 13), 0, 6, 17); r.X != 5 || r.Y != 8 {
		t.Error("wrong multiplying, expecting (", 5, ", ", 8, "), got (", r.X, ", ", r.Y, ")")
	}
	if r := galois.MultiplyV2(galois.NewPoint(5, 8), 0, 2, 17); r.X != 5 || r.Y != 9 {
		t.Error("wrong multiplying, expecting (", 5, ", ", 9, "), got (", r.X, ", ", r.Y, ")")
	}
	if r := galois.MultiplyV2(galois.NewPoint(8, 3), 0, 5, 17); r.X != 8 || r.Y != 14 {
		t.Error("wrong multiplying, expecting (", 8, ", ", 14, "), got (", r.X, ", ", r.Y, ")")
	}
	if r := galois.MultiplyV2(galois.NewPoint(15, 13), 0, 20, 17); r.X != 2 || r.Y != 10 {
		t.Error("wrong multiplying, expecting (", 2, ", ", 10, "), got (", r.X, ", ", r.Y, ")")
	}
	if r := galois.MultiplyV2(galois.NewPoint(15, 13), 0, 21, 17); r.X != 8 || r.Y != 3 {
		t.Error("wrong multiplying, expecting (", 8, ", ", 3, "), got (", r.X, ", ", r.Y, ")")
	}
	if r := galois.MultiplyV2(galois.NewPoint(5, 9), 0, 7, 17); r.X != 5 || r.Y != 9 {
		t.Error("wrong multiplying, expecting (", 5, ", ", 9, "), got (", r.X, ", ", r.Y, ")")
	}
	if r := galois.MultiplyV2(galois.NewPoint(5, 9), 0, 5, 17); r.X != 5 || r.Y != 8 {
		t.Error("wrong multiplying, expecting (", 5, ", ", 8, "), got (", r.X, ", ", r.Y, ")")
	}
}
