package render

import (
	"fmt"
	"strings"
)

type Canvas struct {
	Width  int
	Height int
	Pixels []Color
}

func NewCanvas(width, height int) Canvas {
	return Canvas{
		Width:  width,
		Height: height,
		Pixels: make([]Color, width*height),
	}
}

//--------------------------------------------------------------
// Methods
//--------------------------------------------------------------

func (c Canvas) PixelAt(x, y int) (Color, error) {
	index := x + c.Width*y
	if index < 0 || index >= len(c.Pixels) {
		return NewColor(0, 0, 0), fmt.Errorf("pixel coordinates (%d, %d) out of bounds", x, y)
	}
	return c.Pixels[index], nil
}

func (c Canvas) WritePixel(x, y int, color Color) (*Canvas, error) {
	if x < 0 || x >= c.Width || y < 0 || y >= c.Height {
		return nil, fmt.Errorf("pixel coordinates (%d, %d) out of bounds", x, y)
	}
	c.Pixels[x+c.Width*y] = color
	return &c, nil
}

// --------------------------------------------------------------
// Free functions
// --------------------------------------------------------------
func CanvasToPPM(c Canvas) string {
	output := strings.Builder{}
	output.WriteString("P3\n")
	output.WriteString(fmt.Sprintf("%d %d\n", c.Width, c.Height))
	output.WriteString("255\n")

	for y := 0; y < c.Height; y++ {
		line := strings.Builder{}

		for x := 0; x < c.Width; x++ {
			pixel, _ := c.PixelAt(x, y)
			r := valueToPPM(pixel.Red())
			g := valueToPPM(pixel.Green())
			b := valueToPPM(pixel.Blue())

			line.WriteString(fmt.Sprintf("%d %d %d ", r, g, b))
		}

		lineStr := strings.TrimSpace(line.String())
		runes := []rune(lineStr)

		rest := lineStr[:]

		counter := 0
		for len(rest) > 70 {
			lastSpace := strings.LastIndex(rest[:70], " ")
			counter += lastSpace
			runes[counter] = '\n'
			counter++
			rest = rest[lastSpace+1:]
		}
		lineStr = string(runes)
		output.WriteString(lineStr + "\n")
	}

	return output.String()
}

func valueToPPM(val float64) int {
	if val <= 0 {
		return 0
	} else if val >= 1 {
		return 255
	}
	return int(val * 256)
}
