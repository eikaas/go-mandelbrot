package main

import "fmt"
import "math/cmplx"
import "image/color/palette"

type Mandelbrot struct {
	// We are currently plotting to an Image object. This is a small
	// and perhaps useless wrapper around the golang stdlib image image/png library.
	// It can easily be swapped out to something better, or even a realtime/game library.
	img *Image

	// The pixel size of the object we are rendering to.
	// TODO: Change to int and cast elsewhere.
	image_width  float64
	image_height float64

	// The offset is calculated by dividing the width/height by two.
	xoffset float64
	yoffset float64

	// The max number of iterations for each step through the mandelbrot set
	max_iterations int

	epsilon_x float64
	epsilon_y float64

	// The mandelbrot set is contained betweeb (-2,2), (-2,2)
	// We usualy render a subset within this range. Zooming is done
	// by making these values converge towards eachother: (-1.995, 1.995...)
	bounds_x float64
	bounds_w float64
	bounds_y float64
	bounds_h float64

	epsilon    float64 // Deprecated
	x, y, w, h float64 // Deprecated

	// These are nonsensical. Deprecated
	xzoom float64 // x-axis zoom level.
	yzoom float64 // y-axis zoom level

	// The heart of the object. Changing this will result in radical changes in the way the
	// fractal bends, twists and warps. This contains the starting value.
	initial_c complex128 // Initial value of C. Interesting
}

func NewMandelbrot() *Mandelbrot {
	var m Mandelbrot

	// Set some defaults which produce a useable output
	m.initial_c = complex(-1.0, -0.25)
	m.max_iterations = 10
	m.epsilon = 0.0005
	m.x = -2.0
	m.w = 2.0
	m.y = -2.0
	m.h = 2.0

	return &m
}

func (m *Mandelbrot) Save(filename string) error {
	if err := m.img.Save(filename); err != nil {
		return err
	}

	return nil
}

func (m *Mandelbrot) SetEpsilon(epsilon float64) {
	m.epsilon = epsilon
}

func (m *Mandelbrot) SetBounds(x, y, w, h float64) {
	m.x = x
	m.y = y
	m.w = w
	m.h = h
}

func (m *Mandelbrot) SetZoom(x, y float64) {
	m.xzoom = x
	m.yzoom = y
}

func (m *Mandelbrot) SetMaxIterations(iterations int) {
	m.max_iterations = iterations
}

func (m *Mandelbrot) SetInitialC(c1, c2 float64) {
	m.initial_c = complex(c1, c2)
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

	// Correct for m.w > m.x
	m.image_width = (m.w - m.x) / m.epsilon

	// Correct for m.h > m.y
	m.image_height = (m.h - m.y) / m.epsilon

	// Correct
	m.xoffset = m.image_width / 2
	m.yoffset = m.image_height / 2

	// Testing with array instead. No diff....
	m.img = NewImage(int(m.image_width), int(m.image_height), palette.Plan9[0])

	// faster pixel storage?
	fmt.Println("arr:", int(m.image_width*m.image_height))

	fmt.Printf("\n === Render === \nC=%g, Bounds: mx:%1.2f, mw:%1.2f], my:%1.2f, mh:%1.2f, Epsilon=%1.5f\n", m.initial_c, m.x, m.w, m.y, m.h, m.epsilon)
	// Loop from x=-2.0 to x=2.0

	for x := m.x; x <= m.w; x += m.epsilon {
		// Loop from y=-2.0 to y=2.0
		for y := m.y; y <= m.h; y += m.epsilon {

			iterations = 0

			z = complex(x, y)
			c = m.initial_c

			for cmplx.Abs(z) < 2 && iterations < m.max_iterations {
				z = z*z + c
				iterations++

				xplot := int((x / m.epsilon) + m.xoffset)
				yplot := int((y / m.epsilon) + m.yoffset)

				fmt.Printf("\r")
				fmt.Printf("Plot: (%d,%d), offset: %f, %f", xplot, yplot, m.xoffset, m.yoffset)

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
	fmt.Printf("\nFinished!\n\n")
	//m.SaveToImage() Not used.
}
