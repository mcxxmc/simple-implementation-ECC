package test

import "testing"

func TestAll(t *testing.T) {
	TestInverseFiniteFieldP(t)
	TestDoubling(t)
	TestAdd(t)
	TestMultiple(t)
	TestECC(t)
	TestECDH(t)
}
