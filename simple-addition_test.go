package main

import "testing"

func TestCalculate(t *testing.T) {
	if Calculate(3) != 4 {
		t.Error("expected 4")
	}
}
