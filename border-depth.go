package main

// Takes a FloodFiller.FillArray and returns the border depth
func CalculateBorderDepth(fillArray [][]bool) int {
	// Calculate the Width / Height of the 2d array for later usage
	width := len(fillArray[0])
	height := len(fillArray)

	// Check left depth
	leftBorderDepth := 0
leftDepthLoop:
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if !fillArray[y][x] {
				break leftDepthLoop
			}
		}
		leftBorderDepth++
	}

	// Check right depth
	rightBorderDepth := 0
rightDepthLoop:
	for x := width - 1; x >= 0; x-- {
		for y := 0; y < height; y++ {
			if !fillArray[y][x] {
				break rightDepthLoop
			}
		}
		rightBorderDepth++
	}

	// Check top depth
	topBorderDepth := 0
topDepthLoop:
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if !fillArray[y][x] {
				break topDepthLoop
			}
		}
		topBorderDepth++
	}

	// Check bottom depth
	bottomBorderDepth := 0
bottomDepthLoop:
	for y := height - 1; y >= 0; y-- {
		for x := 0; x < width; x++ {
			if !fillArray[y][x] {
				break bottomDepthLoop
			}
		}
		bottomBorderDepth++
	}

	// Take the smallest of all of our borders, and return that value if it's below our maxBorderWidth
	// specified in the config in the mainfile.
	normalizedDepth := Min(Min(Min(leftBorderDepth, rightBorderDepth), bottomBorderDepth), topBorderDepth)
	if normalizedDepth > maxBorderWidth {
		return 0
	} else {
		return normalizedDepth
	}
}

// Returns the min of two numbers. Math.min
// takes a float but I need it for an int
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
