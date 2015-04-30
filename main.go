package main

import "os"
import "fmt"
import "image/color/palette"

import "code.google.com/p/getopt"

func main() {
	var image_width int
	var image_height int
	var filename string

	/* TODO: Implement these
	optMaxIter := getopt.StringLong("max-iter", 'm', "", "Max number of iterations per value in the set")
	optEpsilon := getopt.StringLong("epsilon", 'e', "", "Epsilon, or step-size when iterating though the bounds of the set")

	optX := getopt.StringLong("dim-x", 'x', "", "Dimensional-X value (-2.0 -> 2.0)")
	optY := getopt.StringLong("dim-y", 'y', "", "Dimensional-Y value (-2.0 -> 2.0)")
	optW := getopt.StringLong("dim-w", 'w', "", "Dimensional-width value (-2.0 -> 2.0)")
	optH := getopt.StringLong("dim-h", 'h', "", "Dimensional-height value (-2.0 -> 2.0)")

	optXOffset := getopt.StringLong("xoffset", 'f', "", "Render translation X-offset")
	optYOffset := getopt.StringLong("yoffset", 'F', "", "Render translation Y-offset")

	optXZoom := getopt.StringLong("xzoom", 'z', "", "X-Axis zoom")
	optYZoom := getopt.StringLong("yzoom", 'Z', "", "Y-Axis zoom")

	optInitialCA := getopt.StringLong("c1", 'c', "", "First term in C")
	optInitialCB := getopt.StringLong("c2", 'C', "", "Second term in C")
	*/

	optFilename := getopt.StringLong("filename", 'O', "", "Output filename")
	optImageWidth := getopt.IntLong("width", 'W', 0, "Image Width")
	optImageHeight := getopt.IntLong("height", 'H', 0, "Image Height")

	getopt.Parse()

	if *optFilename != "" {
		filename = *optFilename
	} else {
		filename = "mandelbrot.png"
	}

	if *optImageWidth != 0 {
		image_width = *optImageWidth
	} else {
		image_width = 1024
	}

	if *optImageHeight != 0 {
		image_height = *optImageHeight
	} else {
		image_height = 1024
	}

	img := NewImage(image_width, image_height, palette.Plan9[0])

	mandelbrot := NewMandelbrot(img)
	mandelbrot.Render()

	if err := img.Save(filename); err != nil {
		fmt.Fprintf(os.Stderr, "There was an error saving the image:", err)
	}

}
