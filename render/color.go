package render

type Color struct {
	data Tuple
}

func NewColor(r, g, b float64) Color {
	return Color{data: NewTuple(r, g, b, 0.0)}
}

func (c Color) Red() float64 {
	return c.data.X
}

func (c Color) Green() float64 {
	return c.data.Y
}

func (c Color) Blue() float64 {
	return c.data.Z
}

func (c Color) Equals(other Color) bool {
	return c.data.Equals(other.data)
}

func (c Color) String() string {
	return c.data.String()
}

func (c Color) Add(other Color) Color {
	result, _ := c.data.Add(other.data)
	return Color{data: result}
}

func (c Color) Sub(other Color) Color {
	result, _ := c.data.Sub(other.data)
	return Color{data: result}
}

func (c Color) Multiply(scalar float64) Color {
	result := c.data.Multiply(scalar)
	return Color{data: result}
}

func (c Color) MultiplyColor(other Color) Color {
	return NewColor(c.Red()*other.Red(), c.Green()*other.Green(), c.Blue()*other.Blue())
}
