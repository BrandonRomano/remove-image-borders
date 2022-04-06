package main

import (
	"image"
	"io"
)

type Pixel struct {
	R int
	G int
	B int
	A int
}

// Calculates the "distance" on how similarly-colored two pixels are
func (this Pixel) distanceFrom(other Pixel) int {
	rDist := Abs(this.R - other.R)
	gDist := Abs(this.G - other.G)
	bDist := Abs(this.B - other.B)
	aDist := Abs(this.A - other.A)
	return rDist + gDist + bDist + aDist
}

// Calculates the absolute value of two integers.
// Go StdLib offers this for floats, but I need it for ints
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Get the bi-dimensional pixel array
// https://stackoverflow.com/questions/33186783/get-a-pixel-array-from-from-golang-image-image
func getPixels(file io.Reader) ([][]Pixel, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	var pixels [][]Pixel
	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}
	return pixels, nil
}

// img.At(x, y).RGBA() returns four uint32 values; we want a Pixel
func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}
