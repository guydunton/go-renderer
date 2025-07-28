package main

import (
	"fmt"
	"guydunton/go-renderer/render"
)

func main() {
	color := render.NewColor(0.5, 0.5, 0.5)
	fmt.Printf("Color: %v\n", color)
}
