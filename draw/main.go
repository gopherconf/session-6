package main

import (
	"image"
	"image/draw"
	"image/png"
	"os"
)

func main()  {
	img1 := readPNGFile("input.png")
	img2 := readPNGFile("input2.png")

	dist := image.NewRGBA(img1.Bounds())

	draw.Draw(dist,img1.Bounds(),img1,image.ZP,0)

	draw.Draw(dist,image.Rectangle{
		Min: image.Point{X:100,Y:100},
		Max: image.Point{X:200,Y:200},
	},img2,image.ZP,0)

	/*draw.Draw(dist,image.Rectangle{
		Min: image.Point{X:100,Y:100},
		Max: image.Point{X:200,Y:200},
	},img2,image.ZP,1)*/


	// save output
	outFile, err := os.Create("out.png")
	if err != nil {
		panic(err)
	}
	png.Encode(outFile, dist)
	outFile.Close()

}

func readPNGFile(path string)image.Image  {
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
