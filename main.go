package main

import "os"
import "fmt"
import "log"
import "flag"

func main() {
	//image_width := flag.Int("W", 512, "Image width")
	//image_height := flag.Int("H", 512, "Image height")
	filename := flag.String("F", "filename.png", "Filename to save to")

	iterations := flag.Int("i", 128, "Maximum iterations per step")

	epsilon := flag.Float64("e", 0.02, "Epsilon: Set-iteration step-size")

	bounds_x := flag.Float64("Bx", -2.0, "Bounds X")
	bounds_y := flag.Float64("By", -2.0, "Bounds Y")
	bounds_w := flag.Float64("Bw", 2.0, "Bounds Width")
	bounds_h := flag.Float64("Bh", 2.0, "Bounds Height")

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
	mandelbrot := NewMandelbrot()

	mandelbrot.SetBounds(*bounds_x, *bounds_y, *bounds_w, *bounds_h)
	mandelbrot.SetEpsilon(*epsilon)
	mandelbrot.SetMaxIterations(*iterations)
	mandelbrot.SetOffset(*offset_x, *offset_y)
	mandelbrot.SetZoom(*zoom_x, *zoom_y)
	mandelbrot.SetInitialC(*initial_c1, *initial_c2)

	log.Println("[*] Rendering...")
	mandelbrot.Render()

	if err := mandelbrot.Save(*filename); err != nil {
		fmt.Println("Unable to save the file:", err)
	}
}
