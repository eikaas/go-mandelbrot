package main

import "fmt"
import "math"
import "math/cmplx"
import "image/color/palette"

type Mandelbrot struct {
	img            *Image  // Pointer to the image were drawing to as defined in Image.go, not the std lib "image"
	max_iterations int     // Max iterations.
	epsilon        float64 // Step-size of the simulation. Improves sharpness

	x, y, w, h float64 // The convergance box. Mandelbrot: -2,2 -2,2

	xoffset float64 // Screen/Camera offset, x
	yoffset float64 // Screen/Camera offset, y

	xzoom float64 // x-axis zoom level.
	yzoom float64 // y-axis zoom level

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

func (m *Mandelbrot) SetOffset(x, y float64) {
	m.xoffset = x
	m.yoffset = y
}

func (m *Mandelbrot) SetMaxIterations(iterations int) {
	m.max_iterations = iterations
}

func (m *Mandelbrot) SetInitialC(c1, c2 float64) {
	m.initial_c = complex(c1, c2)
}

func (m *Mandelbrot) Render() {
	var iterations int

	var z complex128
	var c complex128

	// Calculating the width
	pixel_width := (math.Abs(m.x) + m.w) / m.epsilon
	pixel_height := (math.Abs(m.y) + m.h) / m.epsilon

	// Calculating the offset
	xoffset := pixel_width / 2
	yoffset := pixel_height / 2

	// Calculating the scale
	xscale := pixel_width * m.epsilon * 100
	yscale := pixel_height * m.epsilon * 100

	// NOW the image should be created. We know its size.
	m.img = NewImage(int(pixel_width), int(pixel_height), palette.Plan9[0])

	fmt.Printf("\n === Render === \nC=%g, Bounds=([%1.2f, %1.2f], [%1.2f, %1.2f]), Epsilon=%1.5f\n", m.initial_c, m.x, m.y, m.w, m.h, m.epsilon)
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

				xplot := int((x * xscale) + xoffset)
				yplot := int((y * yscale) + yoffset)

				fmt.Printf("\r")
				fmt.Printf("Plot: (%d,%d)", xplot, yplot)

				if iterations == 0 {
					m.img.Plot(xplot, yplot, palette.Plan9[0])
				} else if iterations == 256 {
					m.img.Plot(xplot, yplot, palette.Plan9[0])

				} else {
					m.img.Plot(xplot, yplot, palette.Plan9[255/iterations])
				}
			}
		}
	}
	fmt.Printf("\nFinished!\n\n")
}
