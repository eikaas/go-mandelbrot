package main

import "fmt"
import "math/cmplx"
import "image/color/palette"

type Mandelbrot struct {
	// Console output
	verbose bool
	debug   bool

	// We are currently plotting to an Image object. This is a small
	// and perhaps useless wrapper around the golang stdlib image image/png library
	// Image.go simply came along since it already did the saving to png and stuff.
	// It can easily be swapped out to something better, or even a realtime/game library.
	img *Image

	// The pixel size of the object we are rendering to.
	// TODO: Change to int and cast elsewhere.
	render_width  int
	render_height int

	// The offset is half the width or height. If the width or height is not
	// divisible by two we risk a small error to the offset. Most likely unnoticable
	xoffset int
	yoffset int

	// The max number of iterations for each step through the mandelbrot set
	max_iterations int

	xepsilon float64
	yepsilon float64

	// The mandelbrot set is contained betweeb (-2,2), (-2,2)
	// We usualy render a subset within this range. Zooming is done
	// by making these values converge towards eachother: (-1.995, 1.995...)
	bounds_x float64
	bounds_w float64
	bounds_y float64
	bounds_h float64

	// The heart of the object. Changing this will result in radical changes in the way the
	// fractal bends, twists and warps. This contains the starting value.
	initial_c complex128 // Initial value of C. Interesting
}

func NewMandelbrot() *Mandelbrot {
	var m Mandelbrot
	// Some sane defaults that will make it render regardless of any parameters are passed
	m.initial_c = complex(-1.0, -0.25)

	m.render_width = 1024
	m.render_height = 1024
	m.max_iterations = 128

	m.bounds_x = -2.0
	m.bounds_w = 2.0
	m.bounds_y = -2.0
	m.bounds_h = 2.0

	return &m
}

func (m *Mandelbrot) Save(filename string) error {
	if err := m.img.Save(filename); err != nil {
		return err
	}

	return nil
}

func (m *Mandelbrot) SetBounds(x, y, w, h float64) {
	m.bounds_x = x
	m.bounds_y = y
	m.bounds_w = w
	m.bounds_h = h
}

func (m *Mandelbrot) SetRenderHeight(height int) {
	m.render_height = height
}

func (m *Mandelbrot) SetRenderWidth(width int) {
	m.render_width = width
}

func (m *Mandelbrot) SetMaxIterations(iterations int) {
	m.max_iterations = iterations
}

func (m *Mandelbrot) SetInitialC(c1, c2 float64) {
	m.initial_c = complex(c1, c2)
}

func (m *Mandelbrot) Verbose(status bool) {
	if status {
		m.verbose = true
	} else {
		m.verbose = false
	}
}

func (m *Mandelbrot) Debug(status bool) {
	if status {
		m.debug = true
	} else {
		m.debug = false
	}
}

/*
	TODO: Update the render function
	Given the bounds of the set we are going to render and the resolution
	we calculate the offset and epsilon.

	Parameters are passed by the commandline:
	-W1024 -H1024
	--Bounds=-2,2,-2,2 (bx, bw, by, bh)

	Offset:

	Ow = width / 2
	Oh = height / 2

	Epsilon:
	We need to calculate the step-size (epsilon) for both the x- and y-axis in order to support
	non-square resolutions:

	EpsilonX := (bw - bx) / width
	EpsilonY := (bh - by) / width

	Zoom:
	Zooming is accomplished by specifying a smaller subset of the fractal set, or rendering at a higher resolution.
	No special function needed.

	Random thoughts:
	We can extend this further and render it realtime. Capturing mouse-input we can map the on-screen coordinate to fractal-set-subset coordinates
	to enable zooming with a mouse and exploring. Should be absolutely doable. Need Screen-to-fractal-set conversion mathz...
*/
func (m *Mandelbrot) Render() {
	var iterations int

	var z complex128
	var c complex128

	m.xepsilon = (m.bounds_w - m.bounds_x) / float64(m.render_width)
	m.yepsilon = (m.bounds_h - m.bounds_y) / float64(m.render_height)

	m.xoffset = m.render_width / 2
	m.yoffset = m.render_height / 2

	m.img = NewImage(int(m.render_width), int(m.render_height), palette.Plan9[0])

	if m.verbose {
		fmt.Printf("[*] Starting render\n")
		fmt.Printf(" == Render Info == \n")
		fmt.Printf("Initial value of C: %g\n", m.initial_c)
		fmt.Printf("Bounds: [(%1.2f, %1.2f), (%1.2f,%1.2f)]\n", m.bounds_x, m.bounds_w, m.bounds_y, m.bounds_h)
		fmt.Printf("Epsilon: %1.3f,%1.3f\n", m.xepsilon, m.yepsilon)
		fmt.Printf("Render Dimension: %dx%d\n", m.render_width, m.render_height)
		fmt.Printf("Offset: %dx%d\n\n", m.xoffset, m.yoffset)
	}

	for x := m.bounds_x; x <= m.bounds_w; x += m.xepsilon {
		for y := m.bounds_y; y <= m.bounds_h; y += m.yepsilon {

			iterations = 0

			z = complex(x, y)
			c = m.initial_c

			for cmplx.Abs(z) < 2 && iterations < m.max_iterations {
				z = z*z + c
				iterations++

				xplot := int((x / m.yepsilon) + float64(m.xoffset))
				yplot := int((y / m.xepsilon) + float64(m.yoffset))

				if m.verbose {
					fmt.Printf("\r")
					fmt.Printf("Rendering %d/%d", xplot, m.render_width)
				}

				if iterations == 0 {
					//m.Plot(xplot, yplot, iterations) // Plotting to an array
					m.img.Plot(xplot, yplot, palette.Plan9[0]) // Plotting to image
				} else if iterations == 256 {
					//m.Plot(xplot, yplot, iterations)
					m.img.Plot(xplot, yplot, palette.Plan9[0])

				} else {
					//m.Plot(xplot, yplot, iterations)
					m.img.Plot(xplot, yplot, palette.Plan9[255/iterations])
				}

			}
		}
	}
	if m.verbose {
		fmt.Printf("\nFinished!\n\n")
	}
}
