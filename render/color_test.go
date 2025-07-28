package render

import (
	"testing"
)

func TestColor(t *testing.T) {
	c1 := NewColor(0.1, 0.2, 0.3)
	c2 := NewColor(0.1, 0.2, 0.3)
	c3 := NewColor(0.4, 0.5, 0.6)

	if !c1.Equals(c2) {
		t.Errorf("Expected colors to be equal, but they are not")
	}

	if c1.Equals(c3) {
		t.Errorf("Expected colors to be different, but they are equal")
	}
}

func TestColorOperations(t *testing.T) {
	a := NewColor(0.9, 0.6, 0.75)
	b := NewColor(0.7, 0.1, 0.25)
	expected := NewColor(1.6, 0.7, 1.0)
	result := a.Add(b)
	if !result.Equals(expected) {
		t.Errorf("Unexpected result during color addition: got %v, want %v", result, expected)
	}

	expected = NewColor(0.2, 0.5, 0.5)
	result = a.Sub(b)
	if !result.Equals(expected) {
		t.Errorf("Unexpected result during color subtraction: got %v, want %v", result, expected)
	}

	a = NewColor(0.2, 0.3, 0.4)
	expected = NewColor(0.4, 0.6, 0.8)
	result = a.Multiply(2)
	if !result.Equals(expected) {
		t.Errorf("Unexpected result during color multiplication: got %v, want %v", result, expected)
	}
}

func TestMultiplyColors(t *testing.T) {
	c1 := NewColor(1, 0.2, 0.4)
	c2 := NewColor(0.9, 1, 0.1)
	expected := NewColor(0.9, 0.2, 0.04)

	result := c1.MultiplyColor(c2)
	if !result.Equals(expected) {
		t.Errorf("Unexpected result during color multiplication: got %v, want %v", result, expected)
	}
}
