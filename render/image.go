package render

import (
	"github.com/UnTea/Tskidisch/linmath"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

type Image struct {
	Pixels []linmath.Vector
	Width  int
	Height int
}

func NewImage(width, height int) Image {
	return Image{
		Pixels: make([]linmath.Vector, width*height),
		Width:  width,
		Height: height,
	}
}

func (img Image) SetPixel(x, y int, color linmath.Vector) {
	img.Pixels[x+y*img.Width] = color
}

func (img Image) Save(path string) {

	i := image.NewRGBA(image.Rect(0, 0, img.Width, img.Height))

	for y := 0; y < img.Height; y++ {
		for x := 0; x < img.Width; x++ {
			//filmFramebuffer := ACESFilm(img[x+y*width])

			i.Set(x, y, color.NRGBA{
				R: uint8(255 * img.Pixels[x+y*img.Width].X),
				G: uint8(255 * img.Pixels[x+y*img.Width].Y),
				B: uint8(255 * img.Pixels[x+y*img.Width].Z),
				A: 255,
			})
		}
	}

	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, i); err != nil {
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
