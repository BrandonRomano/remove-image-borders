package main

func NewFloodFiller(image [][]Pixel) *FloodFiller {
	width := len(image[0])
	height := len(image)
	fillArray := make([][]bool, height)
	for i := range fillArray {
		fillArray[i] = make([]bool, width)
	}
	return &FloodFiller{
		Image:       image,
		ImageWidth:  width,
		ImageHeight: height,
		FillArray:   fillArray,
	}
}

type FloodFiller struct {
	Image       [][]Pixel
	ImageWidth  int
	ImageHeight int
	FillArray   [][]bool
}

func (f *FloodFiller) fill(y int, x int) {
	// Fill the selected pixel
	f.FillArray[y][x] = true

	// Fill Above, if there is a pixel above, it hasn't been visited yet, and if it is within our fill tolerance
	if y > 0 && f.FillArray[y-1][x] != true && f.Image[y][x].distanceFrom(f.Image[y-1][x]) <= floodDistanceMax {
		f.fill(y-1, x)
	}

	// Fill Below if there is a pixel below, it hasn't been visited yet, and if it is within our fill tolerance
	if y < f.ImageHeight-1 && f.FillArray[y+1][x] != true && f.Image[y][x].distanceFrom(f.Image[y+1][x]) <= floodDistanceMax {
		f.fill(y+1, x)
	}

	// Fill Left if there is a pixel to the left, it hasn't been visited yet, and if it is within our fill tolerance
	if x > 0 && f.FillArray[y][x-1] != true && f.Image[y][x].distanceFrom(f.Image[y][x-1]) <= floodDistanceMax {
		f.fill(y, x-1)
	}

	// Fill Right if there is a pixel to the right, it hasn't been visited yet, and if it is within our fill tolerance
	if x < f.ImageWidth-1 && f.FillArray[y][x+1] != true && f.Image[y][x].distanceFrom(f.Image[y][x+1]) <= floodDistanceMax {
		f.fill(y, x+1)
	}
}
