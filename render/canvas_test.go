package render

import (
	"fmt"
	"strings"
	"testing"
)

func TestCanvas(t *testing.T) {
	c := NewCanvas(10, 20)

	if c.Width != 10 {
		t.Error("Canvas width should be 10")
	}
	if c.Height != 20 {
		t.Error("Canvas height should be 20")
	}

	expectedColor := NewColor(0, 0, 0)
	for _, pixel := range c.Pixels {
		if !pixel.Equals(expectedColor) {
			t.Errorf("Expected pixel to be %v, got %v", expectedColor, pixel)
		}
	}
}

func TestCanvasPixels(t *testing.T) {
	c := NewCanvas(10, 20)
	red := NewColor(1, 0, 0)

	c.WritePixel(2, 3, red)
	if p, _ := c.PixelAt(2, 3); !p.Equals(red) {
		t.Errorf("Expected pixel at (2, 3) to be %v, got %v", red, p)
	}
}

func TestCanvasPPMHeader(t *testing.T) {
	c := NewCanvas(5, 3)
	ppm := CanvasToPPM(c)
	expectedHeader := `P3
5 3
255
`
	header := strings.Split(ppm, "\n")[:3]
	headerJoined := strings.Join(header, "\n") + "\n"
	if !strings.EqualFold(headerJoined, expectedHeader) {
		t.Errorf("Expected PPM header to be %q, got %q", expectedHeader, headerJoined)
	}
}

func TestCanvasPPMContent(t *testing.T) {
	c := NewCanvas(5, 3)
	c.WritePixel(0, 0, NewColor(1.5, 0, 0))
	c.WritePixel(2, 1, NewColor(0, 0.5, 0))
	c.WritePixel(4, 2, NewColor(-0.5, 0, 1))

	expectedBody :=
		`255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 128 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 255
`
	ppm := CanvasToPPM(c)
	fmt.Printf("%s", ppm)
	body := strings.Split(ppm, "\n")[3:]
	bodyJoined := strings.Join(body, "\n")
	if !strings.EqualFold(bodyJoined, expectedBody) {
		t.Errorf("Error in PPM body. Expected %s, received %s", expectedBody, bodyJoined)
	}
}

func TestCanvasPPMContentTooLong(t *testing.T) {
	c := NewCanvas(10, 2)
	for i := range c.Pixels {
		c.Pixels[i] = NewColor(1, 0.8, 0.6)
	}
	ppm := CanvasToPPM(c)
	expectedBody :=
		`255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
`
	body := strings.Split(ppm, "\n")[3:]
	if !strings.EqualFold(strings.Join(body, "\n"), expectedBody) {
		t.Errorf("Error in PPM body. Expected %s, received %s", expectedBody, body)
	}
}

func TestCanvasPPMWithReallyLongLines(t *testing.T) {
	c := NewCanvas(34, 1)
	ppm := CanvasToPPM(c)
	expectedBody :=
		`0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
`
	body := strings.Split(ppm, "\n")[3:]
	if !strings.EqualFold(strings.Join(body, "\n"), expectedBody) {
		t.Errorf("Error in PPM body. Expected %s, received %s", expectedBody, body)
	}
}

func TestCanvasEndsWithNewline(t *testing.T) {
	c := NewCanvas(5, 3)
	ppm := CanvasToPPM(c)
	if !strings.HasSuffix(ppm, "\n") {
		t.Error("PPM output should end with a newline")
	}
}
