package main

import (
	"image"
	"image/draw"
	"image/png"
	"os"
)


func main() {
	img1 := readPNGFile("input.png")

	dst := image.NewRGBA(img1.Bounds())

	/*mask := &circle{r: 100, p: image.Point{
		X: 400,
		Y: 400,
	}}*/

	mask := &CustomShape{}
	mask.Bound = img1.Bounds()

	draw.DrawMask(dst, dst.Bounds(), img1, image.ZP, mask, image.ZP, 1)

	// save output
	outFile, err := os.Create("out.png")
	if err != nil {
		panic(err)
	}
	png.Encode(outFile, dst)
	outFile.Close()

}

func readPNGFile(path string) image.Image {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	img, err := png.Decode(file)
	if err != nil {
		panic(err)
	}
	return img
}
