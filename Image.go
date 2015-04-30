package main

import "os"
import "image"
import "image/color"
import "image/png"

type Image struct {
	data   *image.RGBA
	width  int
	height int
}

func (i *Image) GetHeight() int {
	return i.height
}

func (i *Image) GetWidth() int {
	return i.width
}

func NewImage(width int, height int, bg_color color.Color) *Image {
	var img Image

	img.width = width
	img.height = height

	img.data = image.NewRGBA(image.Rect(0, 0, img.width, img.height))

	// Initialize the image to the background color
	for y := img.data.Rect.Min.Y; y < img.data.Rect.Max.Y; y++ {
		for x := img.data.Rect.Min.X; x < img.data.Rect.Max.X; x++ {
			img.data.Set(x, y, bg_color)
		}
	}

	return &img
}

func (i *Image) Save(filename string) error {
	file, err := os.Create(filename)

	if err != nil {
		return err
	} else {
		defer file.Close()
	}

	//jpeg.Encode(file, i.data, &jpeg.Options{80})
	png.Encode(file, i.data)

	return nil
}

func (i *Image) Plot(x int, y int, c color.Color) {
	i.data.Set(x, y, c)
}
