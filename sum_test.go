package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	actual := Sum(1, 2)
	expected := 3
	if actual != expected {
		t.Errorf("Sum(1, 2): actual %v, expected %v", actual, expected)
	}
}
