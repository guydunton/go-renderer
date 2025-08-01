package math

type Matrix struct {
	values [4][4]float64
	size   int
}

func NewMat4(values [4][4]float64) Matrix {
	return Matrix{
		values: values,
		size:   4,
	}
}

func NewMat3(values [3][3]float64) Matrix {
	return Matrix{
		values: [4][4]float64{
			{values[0][0], values[0][1], values[0][2], 0},
			{values[1][0], values[1][1], values[1][2], 0},
			{values[2][0], values[2][1], values[2][2], 0},
			{0, 0, 0, 1},
		},
		size: 3,
	}
}

func NewMat2(values [2][2]float64) Matrix {
	return Matrix{
		values: [4][4]float64{
			{values[0][0], values[0][1], 0, 0},
			{values[1][0], values[1][1], 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		},
		size: 2,
	}
}

func Identity() Matrix {
	return NewMat4([4][4]float64{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	})
}

//--------------------------------------------------------------
// Methods
//--------------------------------------------------------------

func (m Matrix) At(row, col int) float64 {
	return m.values[row][col]
}

func (m Matrix) Equal(m2 Matrix) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if !FloatEqual(m.At(i, j), m2.At(i, j)) {
				return false
			}
		}
	}
	return true
}

func (a Matrix) Multiply(b Matrix) Matrix {
	var result = [4][4]float64{}

	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			result[row][col] =
				a.At(row, 0)*b.At(0, col) +
					a.At(row, 1)*b.At(1, col) +
					a.At(row, 2)*b.At(2, col) +
					a.At(row, 3)*b.At(3, col)
		}
	}

	return NewMat4(result)
}

func (m Matrix) MultiplyTuple(t Tuple) Tuple {
	return NewTuple(
		m.At(0, 0)*t.X+m.At(0, 1)*t.Y+m.At(0, 2)*t.Z+m.At(0, 3)*t.W,
		m.At(1, 0)*t.X+m.At(1, 1)*t.Y+m.At(1, 2)*t.Z+m.At(1, 3)*t.W,
		m.At(2, 0)*t.X+m.At(2, 1)*t.Y+m.At(2, 2)*t.Z+m.At(2, 3)*t.W,
		m.At(3, 0)*t.X+m.At(3, 1)*t.Y+m.At(3, 2)*t.Z+m.At(3, 3)*t.W,
	)
}

func (m Matrix) Transpose() Matrix {
	var result [4][4]float64
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result[j][i] = m.At(i, j)
		}
	}
	return NewMat4(result)
}

func (m Matrix) Determinant() float64 {
	if m.size == 2 {
		return m.values[0][0]*m.values[1][1] - m.values[0][1]*m.values[1][0]
	}
	panic("unimplemented for non 2x2 matrices")
}

func (m Matrix) Submatrix(row, col int) Matrix {
	vals := [4][4]float64{}

	var rowCounter, colCounter int
	for i := 0; i < m.size; i++ {
		if i == row {
			continue
		}
		for j := 0; j < m.size; j++ {
			if j == col {
				continue
			}

			vals[rowCounter][colCounter] = m.At(i, j)
			colCounter++
		}
		rowCounter++
		colCounter = 0
	}

	for i := m.size - 1; i < 4; i++ {
		vals[i][i] = 1
	}
	return NewMat4(vals)
}
