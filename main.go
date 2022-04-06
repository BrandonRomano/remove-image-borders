package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// General script configuration
var validImageExtensions = []string{"png", "jpg", "jpeg"}

func main() {
	// Pull out the specified directory which contains
	// the images we want to convert
	imagesDirectory := os.Args[1]

	// Get the paths to all images in a directory
	images, err := getImagesInDirectory(imagesDirectory)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(images)
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
