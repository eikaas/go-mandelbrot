package main

import "os"
import "fmt"
import "log"
import "flag"
import "strings"
import "strconv"

func main() {
	var err error
	var bounds_x float64
	var bounds_w float64
	var bounds_y float64
	var bounds_h float64

	var render_height int
	var render_width int

	filename := flag.String("F", "filename.png", "Filename to save to")

	iterations := flag.Int("i", 128, "Maximum iterations per step")

	bounds := flag.String("bounds", "", "Set bounds. x,w,y,h ie. \"-2.0,2.0,-2.0,2.0\"")

	initial_c1 := flag.Float64("c1", -1.0, "Initial value of the complex number C, first value")
	initial_c2 := flag.Float64("c2", -0.25, "Initial value of the complex number C, second value")

	help := flag.Bool("help", false, "Print usage")

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	// -bounds supplied
	if *bounds != "" {
		s := strings.Split(*bounds, ",")
		if len(s) != 4 {
			log.Println("Error: Need exactly four numbers in the form: x,w,y,h")
			os.Exit(1)
		}

		if s[0] > s[1] {
			log.Println("x cannot be larger than width: x,w,y,h")
			os.Exit(1)
		}

		if s[2] > s[3] {
			log.Println("y cannot be larger than the height: x,w,y,h")
			os.Exit(1)
		}

		bounds_x, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			fmt.Println("Error parsing bounds number x")
			os.Exit(2)
		}

		bounds_w, err = strconv.ParseFloat(s[1], 64)
		if err != nil {
			fmt.Println("Error parsing bounds number w")
			os.Exit(2)
		}

		bounds_y, err = strconv.ParseFloat(s[2], 64)
		if err != nil {
			fmt.Println("Error parsing bounds number y")
			os.Exit(2)
		}

		bounds_h, err = strconv.ParseFloat(s[3], 64)
		if err != nil {
			fmt.Println("Error parsing bounds number h")
			os.Exit(2)
		}

	} else {
		// Default
		bounds_x = -2.0
		bounds_w = 2.0
		bounds_y = -2.0
		bounds_h = 2.0
	}

	log.Println("[*] Initializing...")
	mandelbrot := NewMandelbrot()

	// Configure the mandelbrot either by defaults or by params
	mandelbrot.SetRenderHeight(render_height)
	mandelbrot.SetRenderWidth(render_width)
	mandelbrot.SetBounds(bounds_x, bounds_y, bounds_w, bounds_h)
	mandelbrot.SetMaxIterations(*iterations)
	mandelbrot.SetInitialC(*initial_c1, *initial_c2)

	log.Println("[*] Rendering...")
	mandelbrot.Render()

	if err := mandelbrot.Save(*filename); err != nil {
		fmt.Println("Unable to save the file:", err)
	}
}
