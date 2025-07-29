package main

import (
	"guydunton/go-renderer/render"
	"log"
	"os"
)

func main() {
	canvas := render.NewCanvas(900, 550)
	start := render.Point(0, 1, 0)
	velocity := render.Vector(1, 1.8, 0).Normalize().Multiply(11.25)
	gravity := render.Vector(0, -0.1, 0)

	p := start

	for counter := 0; counter < 200; counter++ {
		canvas.WritePixel(int(p.X), int(550-p.Y), render.NewColor(1, 1, 1))
		p, _ = p.Add(velocity)
		velocity, _ = velocity.Add(gravity)
	}

	file, err := os.OpenFile("output.ppm", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	ppm := render.CanvasToPPM(canvas)
	if _, err := file.WriteString(ppm); err != nil {
		log.Fatal(err)
	}
}
