package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"os/exec"
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

func stripImageBorder(path string) error {
	// Load up the proper image format
	if strings.HasSuffix(path, "png") {
		image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	} else {
		image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	}

	// Load the image file
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

	// Take the result array of our floodFilter and calculate our border depth
	borderDepth := CalculateBorderDepth(floodFiller.FillArray)
	if borderDepth > 0 {
		fmt.Println("Removing borders from file", path)

		newWidth := floodFiller.ImageWidth - (borderDepth * 2)
		newHeight := floodFiller.ImageHeight - (borderDepth * 2)
		cmd := exec.Command("convert", path, "-gravity", "center", "-crop", fmt.Sprintf("%vx%v+0+0", newWidth, newHeight), path)
		err := cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
}
