package main

import (
	"fmt"
	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
)

func main() {
	img, err := imgio.Open("input.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}

	inverted := effect.Invert(img)
	resized := transform.Resize(inverted, 800, 800, transform.Linear)
	rotated := transform.Rotate(resized, 45, nil)

	if err := imgio.Save("output.png", rotated, imgio.PNGEncoder()); err != nil {
		fmt.Println(err)
		return
	}
}