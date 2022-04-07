package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// General script configuration
var (
	floodDistanceMax     = 10
	maxBorderWidth       = 3
	validImageExtensions = []string{"png", "jpg", "jpeg"}
)

func main() {
	// Pull out the specified directory which contains
	// the images we want to convert
	imagesDirectory := os.Args[1]

	// Get the paths to all images in a directory
	images, err := getImagesInDirectory(imagesDirectory)
	if err != nil {
		log.Fatal(err)
	}

	// Go through each image and strip their borders
	for _, image := range images {
		err = stripImageBorder(image)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Returns the paths to all images in a particular directory
func getImagesInDirectory(directory string) ([]string, error) {
	// An array that will contain all of our images
	var imagePaths []string

	// Walk the directory tree, and look for images
	err := filepath.Walk(
		directory,
		func(path string, fileInfo os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Ensure we're dealing with a file
			if !fileInfo.IsDir() {
				isImageFile := false

				// Loop through each extension to see if this file is an image
				for _, ext := range validImageExtensions {
					if strings.HasSuffix(fileInfo.Name(), ext) {
						isImageFile = true
						break
					}
				}

				// If it's an image, add it to our imagePaths array
				if isImageFile {
					imagePaths = append(imagePaths, path)
				}
			}
			return nil
		},
	)
	return imagePaths, err
}

//
func stripImageBorder(path string) error {
	fmt.Println("Strip Image Border:", path)

	// TODO, this is hardcoded to PNG for now
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	// Load the image
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Convert image into pixels array
	pixels, err := getPixels(file)
	if err != nil {
		return err
	}

	// Flood fill the top left pixel to unearth the border
	floodFiller := NewFloodFiller(pixels)
	floodFiller.fill(0, 0)
	borderDepth := CalculateBorderDepth(floodFiller.FillArray)
	fmt.Println(borderDepth)

	return nil
}
