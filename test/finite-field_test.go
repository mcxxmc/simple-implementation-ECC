package test

import (
	"simple-implementation-ECC/galois"
	"testing"
)

func TestInverseFiniteFieldP(t *testing.T) {
	if r := galois.Inverse(322, 701); r != 455 {
		t.Error("wrong inverse; expecting ", 455, ", got ", r)
	}
}

func TestDoubling(t *testing.T) {
	if r1, r2 := galois.Doubling(0, 15, 13, 17); r1 != 2 || r2 != 10 {
		t.Error("wrong doubling, expecting (", 2, ", ", 10, "), got (", r1, ", ", r2, ")")
	}
	if r1, r2 := galois.Doubling(0, 2, 10, 17); r1 != 12 || r2 != 1 {
		t.Error("wrong doubling, expecting (", 12, ", ", 1, "), got (", r1, ", ", r2, ")")
	}
	if r1, r2 := galois.Doubling(0, 8, 3, 17); r1 != 5 || r2 != 8 {
		t.Error("wrong doubling, expecting (", 5, ", ", 8, "), got (", r1, ", ", r2, ")")
	}
	if r1, r2 := galois.Doubling(0, 1, 12, 17); r1 != 2 || r2 != 7 {
		t.Error("wrong doubling, expecting (", 2, ", ", 7, "), got (", r1, ", ", r2, ")")
	}
}

func TestAdd(t *testing.T) {
	if r1, r2 := galois.Add(15, 13, 2, 10, 17); r1 != 8 || r2 != 3 {
		t.Error("wrong addition, expecting (", 8, ", ", 3, "), got (", r1, ", ", r2, ")")
	}
	if r1, r2 := galois.Add(5, 8, 10, 15, 17); r1 != 6 || r2 != 11 {
		t.Error("wrong addition, expecting (", 6, ", ", 11, "), got (", r1, ", ", r2, ")")
	}
}

func TestMultiple(t *testing.T) {
	if r1, r2 := galois.Multiply(0,15, 13, 2, 17); r1 != 2 || r2 != 10 {
		t.Error("wrong multiplying, expecting (", 2, ", ", 10, "), got (", r1, ", ", r2, ")")
	}
	if r1, r2 := galois.Multiply(0,15, 13, 3, 17); r1 != 8 || r2 != 3 {
		t.Error("wrong multiplying, expecting (", 8, ", ", 3, "), got (", r1, ", ", r2, ")")
	}
	if r1, r2 := galois.Multiply(0,15, 13, 6, 17); r1 != 5 || r2 != 8 {
		t.Error("wrong multiplying, expecting (", 5, ", ", 8, "), got (", r1, ", ", r2, ")")
	}
	if r1, r2 := galois.Multiply(0,5, 8, 2, 17); r1 != 5 || r2 != 9 {
		t.Error("wrong multiplying, expecting (", 5, ", ", 9, "), got (", r1, ", ", r2, ")")
	}
	if r1, r2 := galois.Multiply(0,8, 3, 5, 17); r1 != 8 || r2 != 14 {
		t.Error("wrong multiplying, expecting (", 8, ", ", 14, "), got (", r1, ", ", r2, ")")
	}

	// testing with v2
	if r1, r2 := galois.MultiplyV2(0,15, 13, 2, 17); r1 != 2 || r2 != 10 {
		t.Error("wrong multiplying, expecting (", 2, ", ", 10, "), got (", r1, ", ", r2, ")")
	}
	if r1, r2 := galois.MultiplyV2(0,15, 13, 3, 17); r1 != 8 || r2 != 3 {
		t.Error("wrong multiplying, expecting (", 8, ", ", 3, "), got (", r1, ", ", r2, ")")
	}
	if r1, r2 := galois.MultiplyV2(0,15, 13, 6, 17); r1 != 5 || r2 != 8 {
		t.Error("wrong multiplying, expecting (", 5, ", ", 8, "), got (", r1, ", ", r2, ")")
	}
	if r1, r2 := galois.MultiplyV2(0,5, 8, 2, 17); r1 != 5 || r2 != 9 {
		t.Error("wrong multiplying, expecting (", 5, ", ", 9, "), got (", r1, ", ", r2, ")")
	}
	if r1, r2 := galois.MultiplyV2(0,8, 3, 5, 17); r1 != 8 || r2 != 14 {
		t.Error("wrong multiplying, expecting (", 8, ", ", 14, "), got (", r1, ", ", r2, ")")
	}
}
