package math

import (
	"math"
	"testing"
)

func TestTuple(t *testing.T) {
	a := NewTuple(4.3, -4.2, 3.1, 1.0)
	if a.X != 4.3 || a.Y != -4.2 || a.Z != 3.1 || a.W != 1.0 {
		t.Errorf("Expected Tuple(4.3, -4.2, 3.1, 1.0), got (%f, %f, %f, %f)", a.X, a.Y, a.Z, a.W)
	}
	if !a.IsPoint() {
		t.Error("Expected Tuple with W = 1.0 to be point")
	}
	if a.IsVector() {
		t.Error("Expected Tuple with W = 0.0 to not be a vector")
	}

	b := NewTuple(4.3, -4.2, 3.1, 0.0)
	if b.IsPoint() {
		t.Error("Tuple with W=0.0 should not be a point")
	}
	if !b.IsVector() {
		t.Error("Tuple with W=0.0 should be a vector")
	}

	c1 := NewTuple(1.00000000001, 2.0, 3.0, 1.0)
	c2 := NewTuple(1.0, 2.0, 3.0, 1.0)

	if !c1.Equals(c2) {
		t.Error("Expected Tuples to be equal, but they are not")
	}
}

func TestPoint(t *testing.T) {
	p := Point(4.0, -4.0, 3.0)

	if p.X != 4.0 || p.Y != -4.0 || p.Z != 3.0 || p.W != 1.0 {
		t.Errorf("Expected Point(4.0, -4.0, 3.0), got (%f, %f, %f, %f)", p.X, p.Y, p.Z, p.W)
	}
}

func TestVector(t *testing.T) {
	v := Vector(4, -4, 3)

	if v.X != 4.0 || v.Y != -4.0 || v.Z != 3.0 || v.W != 0.0 {
		t.Errorf("Expected Vector(4.0, -4.0, 3.0), got (%f, %f, %f, %f)", v.X, v.Y, v.Z, v.W)
	}
}

func TestAdd(t *testing.T) {
	t1 := NewTuple(3, -2, 5, 1)
	t2 := NewTuple(-2, 3, 1, 0)

	result, _ := t1.Add(t2)
	if !Equals(result, NewTuple(1, 1, 6, 1)) {
		t.Errorf("Expected result of addition to be (1, 1, 6, 1), got (%f, %f, %f, %f)", result.X, result.Y, result.Z, result.W)
	}

	p1 := Point(1, 2, 3)
	p2 := Point(3, 4, 5)

	result, err := p1.Add(p2)
	if err == nil {
		t.Error("Expected error when adding two points, but got none")
	}
}

func TestSub(t *testing.T) {
	p1 := Point(3, 2, 1)
	p2 := Point(5, 6, 7)

	result, _ := p1.Sub(p2)
	if !result.Equals(Vector(-2, -4, -6)) {
		t.Errorf("Expected result of subtraction to be (-2, -4, -6), got (%f, %f, %f)", result.X, result.Y, result.Z)
	}
}

func TestSub2(t *testing.T) {
	p := Point(3, 2, 1)
	v := Vector(5, 6, 7)

	result, _ := p.Sub(v)
	if !result.Equals(Point(-2, -4, -6)) {
		t.Errorf("Expected result of subtraction to be (-2, -4, -6), got (%f, %f, %f)", result.X, result.Y, result.Z)
	}
}

func TestSub3(t *testing.T) {
	v1 := Vector(3, 2, 1)
	v2 := Vector(5, 6, 7)

	result, _ := v1.Sub(v2)
	if !result.Equals(Vector(-2, -4, -6)) {
		t.Errorf("Expected result of subtraction to be (-2, -4, -6), got (%f, %f, %f)", result.X, result.Y, result.Z)
	}
}

func TestSubErr(t *testing.T) {
	v := Vector(1, 2, 3)
	p := Point(3, 4, 5)

	_, err := v.Sub(p)
	if err == nil {
		t.Error("Expected error when subtracting a point from a vector")
	}
}

func TestVecFromZero(t *testing.T) {
	zero := Vector(0, 0, 0)
	v := Vector(1, -2, 3)

	result, _ := zero.Sub(v)
	if !result.Equals(Vector(-1, 2, -3)) {
		t.Error("Expected result of subtraction to be (-1, 2, -3)")
	}
}

func TestTupleNegation(t *testing.T) {
	a := NewTuple(1, -2, 3, -4)
	result := Negate(a)
	if !result.Equals(NewTuple(-1, 2, -3, 4)) {
		t.Errorf("Expected negation to be (-1, 2, -3, 4), got (%f, %f, %f, %f)", result.X, result.Y, result.Z, result.W)
	}
}

func MultiplyTupleByScalar(t *testing.T) {
	a := NewTuple(1, -2, 3, -4)
	result := a.Multiply(3.5)
	if !Equals(result, NewTuple(3.5, -7, 10.5, -14)) {
		t.Errorf("Expected multiplication by scalar to be (3.5, -7, 10.5, -14), got (%f, %f, %f, %f)", result.X, result.Y, result.Z, result.W)
	}

	a = NewTuple(1, -2, 3, -4)
	result = a.Multiply(0.5)
	if !Equals(result, NewTuple(0.5, -1, 1.5, -2)) {
		t.Errorf("Expected multiplication by scalar to be (0.5, -1, 1.5, -2), got (%f, %f, %f, %f)", result.X, result.Y, result.Z, result.W)
	}
}

func TestDivideByScalar(t *testing.T) {
	a := NewTuple(1, -2, 3, -4)
	result := a.Divide(2)
	if !Equals(result, NewTuple(0.5, -1, 1.5, -2)) {
		t.Errorf("Expected division by scalar to be (0.5, -1, 1.5, -2), got (%f, %f, %f, %f)", result.X, result.Y, result.Z, result.W)
	}
}

func TestMagnitude(t *testing.T) {
	v := Vector(1, 0, 0)
	r := v.Magnitude()
	if !FloatEqual(r, 1) {
		t.Errorf("Expected magnitude of (1, 0, 0) to be 1.0, got %f", r)
	}

	v = Vector(0, 1, 0)
	r = v.Magnitude()
	if !FloatEqual(r, 1.0) {
		t.Errorf("Expected magnitude of (0, 1, 0) to be 1.0, got %f", r)
	}

	v = Vector(0, 0, 1)
	r = v.Magnitude()
	if !FloatEqual(r, 1.0) {
		t.Errorf("Expected magnitude of (0, 0, 1) to be 1.0, got %f", r)
	}

	v = Vector(1, 2, 3)
	r = v.Magnitude()
	if !FloatEqual(r, math.Sqrt(14)) {
		t.Errorf("Expected magnitude of (1, 2, 3) to be sqrt(14), got %f", r)
	}

	v = Vector(-1, -2, -3)
	r = v.Magnitude()
	if !FloatEqual(r, math.Sqrt(14)) {
		t.Errorf("Expected magnitude of (-1, -2, -3) to be sqrt(14), got %f", r)
	}

	p := Point(1, 2, 3)
	if !FloatEqual(p.Magnitude(), 0) {
		t.Errorf("Expected magnitude of point (1, 2, 3) to be 0, got %f", p.Magnitude())
	}
}

func TestNormalization(t *testing.T) {
	v := Vector(4, 0, 0)
	r := v.Normalize()
	if !r.Equals(Vector(1, 0, 0)) {
		t.Errorf("Expected normalization of %s to be (1, 0, 0), got (%f, %f, %f)", v, r.X, r.Y, r.Z)
	}

	v = Vector(1, 2, 3)
	r = v.Normalize()
	if !r.Equals(Vector(1/math.Sqrt(14), 2/math.Sqrt(14), 3/math.Sqrt(14))) {
		t.Errorf("Expected normalization of (1, 2, 3) to be (1/sqrt(14), 2/sqrt(14), 3/sqrt(14)), got (%f, %f, %f)", r.X, r.Y, r.Z)
	}

	v = Vector(1, 2, 3)
	norm := v.Normalize()
	if !FloatEqual(norm.Magnitude(), 1.0) {
		t.Errorf("Expected normalized vector to have magnitude 1, got %f", norm.Magnitude())
	}
}

func TestDotProduct(t *testing.T) {
	v1 := Vector(1, 2, 3)
	v2 := Vector(2, 3, 4)

	dot := v1.Dot(v2)
	if !FloatEqual(dot, 20) {
		t.Errorf("Expected dot product of (%s) and (%s) to be 20, got %f", v1, v2, dot)
	}
}

func TestCrossProduct(t *testing.T) {
	v1 := Vector(1, 2, 3)
	v2 := Vector(2, 3, 4)
	expected := Vector(-1, 2, -1)

	r := v1.Cross(v2)
	if !r.Equals(expected) {
		t.Errorf("Expected cross product of (%s) and (%s) to be %s, got %s", v1, v2, expected, r)
	}

	expected = Vector(1, -2, 1)
	r = v2.Cross(v1)
	if !r.Equals(expected) {
		t.Errorf("Expected cross product of (%s) and (%s) to be %s, got %s", v2, v1, expected, r)
	}
}
