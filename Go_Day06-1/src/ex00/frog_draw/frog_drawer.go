package frogdraw

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

const (
	WIDTH  int = 300
	HEIGHT int = 300
)

var upLeft = image.Point{0, 0}
var lowRight = image.Point{WIDTH, HEIGHT}

var colorGreen = color.RGBA{195, 255, 104, 0xff}
var colorPink = color.RGBA{251, 231, 239, 0xff}

func CreateLogoFrog(fileName string) error {
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {
			img.Set(x, y, color.White)
		}
	}
	drawPixelFrog(img)

	f, err := os.Create(fileName)
	png.Encode(f, img)
	return err
}

func drawPixelFrog(img *image.RGBA) {
	// Frog head
	for x := 70; x <= 230; x++ {
		for y := 100; y <= 200; y++ {
			img.Set(x, y, colorGreen)
		}
	}

	// Frog eye base
	for x := 85; x <= 130; x++ {
		for y := 68; y <= 208; y++ {
			img.Set(x, y, colorGreen)
		}
	}
	for x := 170; x <= 215; x++ {
		for y := 68; y <= 208; y++ {
			img.Set(x, y, colorGreen)
		}
	}

	// Frog line between eyes
	for x := 130; x <= 170; x++ {
		for y := 84; y <= 100; y++ {
			img.Set(x, y, colorGreen)
		}
	}

	// Frog eyes
	for x := 100; x <= 115; x++ {
		for y := 84; y <= 116; y++ {
			img.Set(x, y, color.Black)
		}
	}
	for x := 185; x <= 200; x++ {
		for y := 84; y <= 116; y++ {
			img.Set(x, y, color.Black)
		}
	}

	// Frog mouth
	for x := 130; x <= 142; x++ {
		for y := 116; y <= 148; y++ {
			img.Set(x, y, color.Black)
		}
	}
	for x := 158; x <= 170; x++ {
		for y := 116; y <= 148; y++ {
			img.Set(x, y, color.Black)
		}
	}
	for x := 142; x <= 158; x++ {
		for y := 132; y <= 148; y++ {
			img.Set(x, y, color.Black)
		}
	}

	// Frog cheeks
	for x := 85; x <= 115; x++ {
		for y := 116; y <= 132; y++ {
			img.Set(x, y, colorPink)
		}
	}
	for x := 185; x <= 215; x++ {
		for y := 116; y <= 132; y++ {
			img.Set(x, y, colorPink)
		}
	}

}
