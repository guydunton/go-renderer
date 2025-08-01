package math

import (
	"testing"
)

func TestNewMat4(t *testing.T) {
	m := NewMat4([4][4]float64{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5},
	})
	if !FloatEqual(m.At(0, 0), 1) {
		t.Error("Error")
	}
	if !FloatEqual(m.At(0, 3), 4) {
		t.Error("Error")
	}
	if !FloatEqual(m.At(1, 0), 5.5) {
		t.Error("Error")
	}
	if !FloatEqual(m.At(1, 2), 7.5) {
		t.Error("Error")
	}
	if !FloatEqual(m.At(2, 2), 11) {
		t.Error("Error")
	}
	if !FloatEqual(m.At(3, 0), 13.5) {
		t.Error("Error")
	}
	if !FloatEqual(m.At(3, 3), 16.5) {
		t.Error("Error")
	}
}

func TestMatrix3x3(t *testing.T) {
	m := NewMat3([3][3]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
	if !FloatEqual(m.At(0, 0), 1) {
		t.Error("Error")
	}
	if !FloatEqual(m.At(0, 2), 3) {
		t.Error("Error")
	}
	if !FloatEqual(m.At(1, 0), 4) {
		t.Error("Error")
	}
	if !FloatEqual(m.At(1, 1), 5) {
		t.Error("Error")
	}
	if !FloatEqual(m.At(2, 2), 9) {
		t.Error("Error")
	}
	if !FloatEqual(m.At(3, 3), 1) {
		t.Error("Error")
	}
}

func TestMatrix2x2(t *testing.T) {
	m := NewMat2([2][2]float64{
		{1, 2},
		{3, 4},
	})
	if !FloatEqual(m.At(0, 0), 1) {
		t.Error("Error")
	}
	if !FloatEqual(m.At(0, 1), 2) {
		t.Error("Error")
	}
	if !FloatEqual(m.At(1, 0), 3) {
		t.Error("Error")
	}
	if !FloatEqual(m.At(1, 1), 4) {
		t.Error("Error")
	}
}

func TestEqual(t *testing.T) {
	m1 := NewMat4([4][4]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	})
	m2 := NewMat4([4][4]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	})
	m3 := NewMat4([4][4]float64{
		{2, 3, 4, 5},
		{6, 7, 8, 9},
		{10, 11, 12, 13},
		{14, 15, 16, 17},
	})
	if !m1.Equal(m2) {
		t.Error("Error")
	}
	if m1.Equal(m3) {
		t.Error("Error")
	}
}

func TestMultiply(t *testing.T) {
	m1 := NewMat4([4][4]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	})
	m2 := NewMat4([4][4]float64{
		{-2, 1, 2, 3},
		{3, 2, 1, -1},
		{4, 3, 6, 5},
		{1, 2, 7, 8},
	})

	expected := NewMat4([4][4]float64{
		{20, 22, 50, 48},
		{44, 54, 114, 108},
		{40, 58, 110, 102},
		{16, 26, 46, 42},
	})

	if !m1.Multiply(m2).Equal(expected) {
		t.Error("Error")
	}
}

func TestMatrixVecMultiply(t *testing.T) {
	m := NewMat4([4][4]float64{
		{1, 2, 3, 4},
		{2, 4, 4, 2},
		{8, 6, 4, 1},
		{0, 0, 0, 1},
	})
	tuple := NewTuple(1, 2, 3, 1)
	expected := NewTuple(18, 24, 33, 1)
	result := m.MultiplyTuple(tuple)
	if !result.Equal(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestIdentityMatrix(t *testing.T) {
	m := NewMat4([4][4]float64{
		{0, 1, 2, 3},
		{1, 2, 4, 8},
		{2, 4, 4, 8},
		{4, 8, 12, 24},
	})
	identity := Identity()
	result := m.Multiply(identity)
	if !result.Equal(m) {
		t.Errorf("expected %v, got %v", m, result)
	}
	tuple := NewTuple(1, 2, 3, 4)
	resultTuple := identity.MultiplyTuple(tuple)
	if !resultTuple.Equal(tuple) {
		t.Errorf("expected %v, got %v", tuple, result)
	}
}

func TestMatrixTranspose(t *testing.T) {
	m := NewMat4([4][4]float64{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
		{12, 13, 14, 15},
	})
	expected := NewMat4([4][4]float64{
		{0, 4, 8, 12},
		{1, 5, 9, 13},
		{2, 6, 10, 14},
		{3, 7, 11, 15},
	})
	result := m.Transpose()
	if !result.Equal(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	identity := Identity()
	result = identity.Transpose()
	if !result.Equal(identity) {
		t.Errorf("Expected identity matrix, got %v", result)
	}
}

func TestDeterminant(t *testing.T) {
	m1 := NewMat2([2][2]float64{
		{1, 5},
		{-3, 2},
	})

	result := m1.Determinant()
	if !FloatEqual(result, 17) {
		t.Errorf("Expected 17, got %v", result)
	}
}

func TestSubMatrix3x3(t *testing.T) {
	m3 := NewMat3([3][3]float64{
		{1, 5, 0},
		{-3, 2, 7},
		{0, 6, -3},
	})

	expected := NewMat2([2][2]float64{
		{-3, 2},
		{0, 6},
	})
	result := m3.Submatrix(0, 2)
	if !result.Equal(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSubmatrix4x4(t *testing.T) {
	m4 := NewMat4([4][4]float64{
		{-6, 1, 1, 6},
		{-8, 5, 8, 6},
		{-1, 0, 8, 2},
		{-7, 1, -1, 1},
	})

	expected := NewMat3([3][3]float64{
		{-6, 1, 6},
		{-8, 8, 6},
		{-7, -1, 1},
	})
	result := m4.Submatrix(2, 1)
	if !result.Equal(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
