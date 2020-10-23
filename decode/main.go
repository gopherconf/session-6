package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	file, err := os.Open("input.png")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	img, err := png.Decode(file)
	if err != nil {
		panic(err)
	}

	// reverse by opacity
	reversed := reverse(img)

	// save output
	outFile, err := os.Create("out.png")
	if err != nil {
		panic(err)
	}
	png.Encode(outFile, reversed)
	outFile.Close()

	// fmt.Printf("%+v ",img.Bounds())

	//fmt.Printf("%+v ",img.At(250,150))

	/*img,format,err := image.Decode(file)
	if err != nil{
		panic(err)
	}

	fmt.Printf("%+v - f: %s",img.Bounds(),format)
	*/
}

func reverse(img image.Image) image.Image {
	out := image.NewRGBA(img.Bounds())

	for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
			_, _, _, a := img.At(x, y).RGBA()
			if a > 0 {
				out.Set(x, y, color.Transparent)
			} else {
				out.Set(x, y, color.RGBA{
					R: 0,
					G: 0,
					B: 0,
					A: 255,
				})
			}

		}
	}

	return out
}
