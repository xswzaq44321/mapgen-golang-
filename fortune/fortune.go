package main

import (
	"image"
	"image/png"
	"os"
)

func main() {
	// Create a blank image 100x200 pixels
	myImage := image.NewRGBA(image.Rect(0, 0, 10, 4))

	// outputFile is a File type which satisfies Writer interface
	outputFile, err := os.Create("test.png")
	if err != nil {
		// Handle error
	}
	myImage.Pix[0] = 255 // 1st pixel red
	myImage.Pix[1] = 0   // 1st pixel green
	myImage.Pix[2] = 0   // 1st pixel blue
	myImage.Pix[3] = 255 // 1st pixel alpha

	// Encode takes a writer interface and an image interface
	// We pass it the File and the RGBA
	png.Encode(outputFile, myImage)

	// Don't forget to close files
	outputFile.Close()
}
