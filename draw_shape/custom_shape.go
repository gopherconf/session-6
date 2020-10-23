package main

import (
	"image"
	"image/color"
)

type CustomShape struct {
	Bound image.Rectangle
}

func (c *CustomShape) ColorModel() color.Model {
	return color.Alpha16Model
}

func (c *CustomShape) Bounds() image.Rectangle {
	return c.Bound
}

func (c *CustomShape) At(x, y int) color.Color {
	if x%6 == 0 && y%6 == 0 {
		return color.Alpha{A: 0}
	}
	return color.Alpha{A: 255}
}
