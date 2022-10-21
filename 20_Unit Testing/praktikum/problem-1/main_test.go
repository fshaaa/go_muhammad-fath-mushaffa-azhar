package main

import (
	"testing"
)

func TestAddition(t *testing.T) {
	if Addition(2, 3) != 5 {
		t.Errorf("expected 2 + 3 = 5")
	}
	if Addition(-5, 3) != -2 {
		t.Errorf("expected -5 + 3 = 2")
	}
	if Addition(2, -7) != -5 {
		t.Errorf("expected 2 + (-7) = (-5)")
	}
	if Addition(-6, -2) != -8 {
		t.Errorf("expected (-6) + (-2) = (-8)")
	}
}

func TestSubtraction(t *testing.T) {
	if Subtraction(3, 1) != 2 {
		t.Errorf("expected 3 - 1 = 2")
	}
	if Subtraction(2, -6) != 8 {
		t.Errorf("expected 2 - (-6) = 8")
	}
	if Subtraction(-4, 8) != -12 {
		t.Errorf("expected (-4) - (-8) = (-12)")
	}
	if Subtraction(-2, -6) != 4 {
		t.Errorf("expected (-2) - (-6) = 4")
	}
}

func TestDivision(t *testing.T) {
	if Division(4, 2) != 2 {
		t.Errorf("expected 4 / 2 = 2")
	}
	if Division(4, -2) != -2 {
		t.Errorf("expected 4 / (-2) = (-2)")
	}
	if Division(-6, 3) != -2 {
		t.Errorf("expected (-6) / 3 = (-2)")
	}
	if Division(-9, -3) != 3 {
		t.Errorf("expected (-9) / (-3) = 3")
	}
}

func TestMultiplication(t *testing.T) {
	if Multiplication(2, 3) != 6 {
		t.Errorf("expected 2 * 3 = 6")
	}
	if Multiplication(3, -4) != -12 {
		t.Errorf("expected 3 * (-4) = (-12)")
	}
	if Multiplication(-2, 2) != -4 {
		t.Errorf("expected (-2) * 2 = (-4)")
	}
	if Multiplication(-3, -1) != 3 {
		t.Errorf("expected (-3) * (-1) = 3")
	}
}
