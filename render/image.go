package render

import (
	"github.com/UnTea/Tskidisch/linmath"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

type Framebuffer struct {
	Pixels []linmath.Vector
	Width  int
	Height int
}

func NewFramebuffer(width, height int) Framebuffer {
	return Framebuffer{
		Pixels: make([]linmath.Vector, width*height),
		Width:  width,
		Height: height,
	}
}

func (framebuffer Framebuffer) SetPixel(x, y int, color linmath.Vector) {
	framebuffer.Pixels[x+y*framebuffer.Width] = color
}

func (framebuffer Framebuffer) SaveImage(path string) {

	img := image.NewRGBA(image.Rect(0, 0, framebuffer.Width, framebuffer.Height))

	for y := 0; y < framebuffer.Height; y++ {
		for x := 0; x < framebuffer.Width; x++ {
			//filmFramebuffer := ACESFilm(framebuffer[x+y*width])

			img.Set(x, y, color.NRGBA{
				R: uint8(255 * framebuffer.Pixels[x+y*framebuffer.Width].X),
				G: uint8(255 * framebuffer.Pixels[x+y*framebuffer.Width].Y),
				B: uint8(255 * framebuffer.Pixels[x+y*framebuffer.Width].Z),
				A: 255,
			})
		}
	}

	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		err := f.Close()
		if err != nil {
			return
		}
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
