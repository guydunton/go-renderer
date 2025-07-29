package render

import (
	"fmt"
	"math"
)

type Tuple struct {
	X, Y, Z, W float64
}

func NewTuple(x, y, z, w float64) Tuple {
	return Tuple{X: x, Y: y, Z: z, W: w}
}

//--------------------------------------------------------------
// Methods
//--------------------------------------------------------------

func (t Tuple) IsPoint() bool {
	return t.W == 1.0
}

func (t Tuple) IsVector() bool {
	return t.W == 0.0
}

func (t Tuple) Equals(other Tuple) bool {
	return FloatEqual(t.X, other.X) &&
		FloatEqual(t.Y, other.Y) &&
		FloatEqual(t.Z, other.Z) &&
		FloatEqual(t.W, other.W)
}

func (t Tuple) String() string {
	return fmt.Sprintf("(%f, %f, %f, %f)", t.X, t.Y, t.Z, t.W)
}

func (t Tuple) Add(other Tuple) (Tuple, error) {
	w := t.W + other.W
	if w < 0 || w > 1 {
		return Tuple{}, fmt.Errorf("invalid W value: %f", w)
	}
	return NewTuple(t.X+other.X, t.Y+other.Y, t.Z+other.Z, w), nil
}

func (t Tuple) Sub(other Tuple) (Tuple, error) {
	w := t.W - other.W
	if w < 0 || w > 1 {
		return Tuple{}, fmt.Errorf("invalid W value: %f", w)
	}
	return NewTuple(t.X-other.X, t.Y-other.Y, t.Z-other.Z, t.W-other.W), nil
}

func (t Tuple) Multiply(scalar float64) Tuple {
	return NewTuple(t.X*scalar, t.Y*scalar, t.Z*scalar, t.W*scalar)
}

func (t Tuple) Divide(scalar float64) Tuple {
	if scalar == 0 {
		panic("division by zero")
	}
	return NewTuple(t.X/scalar, t.Y/scalar, t.Z/scalar, t.W/scalar)
}

func (t Tuple) Magnitude() float64 {
	if t.IsVector() {
		return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z)
	}
	return 0
}

func (t Tuple) Normalize() Tuple {
	m := t.Magnitude()
	return NewTuple(t.X/m, t.Y/m, t.Z/m, t.W/m)
}

func (t Tuple) Dot(other Tuple) float64 {
	return t.X*other.X + t.Y*other.Y + t.Z*other.Z + t.W*other.W
}

func (t Tuple) Cross(other Tuple) Tuple {
	if !t.IsVector() || !other.IsVector() {
		panic("Cross product is only defined for vectors")
	}
	return NewTuple(
		t.Y*other.Z-t.Z*other.Y,
		t.Z*other.X-t.X*other.Z,
		t.X*other.Y-t.Y*other.X,
		0.0, // Cross product results in a vector
	)
}

//--------------------------------------------------------------
// Free functions
//--------------------------------------------------------------

func Point(x, y, z float64) Tuple {
	return NewTuple(x, y, z, 1.0)
}

func Vector(x, y, z float64) Tuple {
	return NewTuple(x, y, z, 0.0)
}

func FloatEqual(a, b float64) bool {
	const epsilon = 1e-10
	return math.Abs(a-b) < epsilon
}

func Equals(a1 Tuple, a2 Tuple) bool {
	return a1.Equals(a2)
}

func Negate(t Tuple) Tuple {
	return NewTuple(-t.X, -t.Y, -t.Z, -t.W)
}
