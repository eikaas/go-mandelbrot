package main

import "os"
import "fmt"
import "log"
import "flag"
import "image/color/palette"

import "code.google.com/p/getopt"

func main() {
	image_width := flag.Int("W", 512, "Image width")
	image_height := flag.Int("H", 512, "Image height")
	filename := flag.String("F", "mandelbrot.png", "Filename to save to")

	iterations := flag.Int("i", 128, "Maximum iterations per step")

	epsilon := flag.Float64("e", 0.005, "Epsilon: Set-iteration step-size")

	bounds_x := flag.Float64("x", -2.0, "Bounds X")
	bounds_y := flag.Float64("y", -2.0, "Bounds Y")
	bounds_w := flag.Float64("w", 2.0, "Bounds Width")
	bounds_h := flag.Float64("h", 2.0, "Bounds Height")

	offset_x := flag.Float64("ox", 512.0, "Render X-offset")
	offset_y := flag.Float64("oy", 512.0, "Render Y-offset")

	zoom_x := flag.Float64("zx", 1024.0, "x-axis zoom")
	zoom_y := flag.Float64("zy", 1024.0, "y-axis zoom")

	initial_c1 := flag.Float64("c1", -1.0, "Initial value of the complex number C, first value")
	initial_c2 := flag.Float64("c2", -0.25, "Initial value of the complex number C, second value")

	help := flag.Bool("help", false, "Print usage")

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	log.Println("[*] Initializing...")
	img := NewImage(*image_width, *image_height, palette.Plan9[0])
	mandelbrot := NewMandelbrot(img)

	mandelbrot.SetBounds(*bounds_x, *bounds_y, *bounds_w, *bounds_h)
	mandelbrot.SetEpsilon(*epsilon)
	mandelbrot.SetMaxIterations(*iterations)
	mandelbrot.SetOffset(*offset_x, *offset_y)
	mandelbrot.SetZoom(*zoom_x, *zoom_y)
	mandelbrot.SetInitialC(*initial_c1, *initial_c2)

	log.Println("[*] Rendering...")
	mandelbrot.Render()

	if err := img.Save(*filename); err != nil {
		fmt.Fprintf(os.Stderr, "There was an error saving the image:", err)
	}
}

func main_old() {
	var image_width int
	var image_height int
	var filename string

	/* TODO: Implement these, Perhaps getopt is a bad choise for argument handling in this case.
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

	// --filename
	if *optFilename != "" {
		filename = *optFilename
	} else {
		filename = "mandelbrot.png"
	}

	// --width
	if *optImageWidth != 0 {
		image_width = *optImageWidth
	} else {
		image_width = 1024
	}

	// --height
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
