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

		// 2 versions of the string so we can manipulate it
		lineStr := strings.TrimSpace(line.String())
		output.WriteString(chunkLine(lineStr) + "\n")
	}

	return output.String()
}

func chunkLine(line string) string {
	// Create a version of the string that can be manipulated
	runes := []rune(line)

	// View on the string
	lineView := line[:]

	// Keep a newlineIndex to track the position of the newlines
	newlineIndex := 0
	for len(lineView) > 70 {
		newlineIndex += strings.LastIndex(lineView[:70], " ")

		// Insert the newline character & bump on the index
		runes[newlineIndex] = '\n'
		newlineIndex++

		// Reset the view to the remaining part of the string
		lineView = line[newlineIndex:]
	}
	return string(runes)
}

func valueToPPM(val float64) int {
	if val <= 0 {
		return 0
	} else if val >= 1 {
		return 255
	}
	return int(val * 256)
}
