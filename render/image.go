package render

import (
	"github.com/UnTea/Tskidisch/linmath"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
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

func (img Image) GetPixel(x, y int) linmath.Vector {
	return img.Pixels[x+y*img.Width]
}

func (img Image) GetPixelUV(u, v float64) linmath.Vector {
	x := int(float64(img.Width) * (1.0 - u))
	y := int(float64(img.Height) * (1.0 - v))
	return img.GetPixel(x, y)
}

func (img Image) GetPixelBySphericalCoordinates(phi, theta float64) linmath.Vector {
	u := (phi + math.Pi) / (2 * math.Pi)
	v := (theta + math.Pi / 2) / (math.Pi)
	return img.GetPixelUV(u, v)
}

func (img Image) Save(path string) {

	i := image.NewRGBA(image.Rect(0, 0, img.Width, img.Height))

	for y := 0; y < img.Height; y++ {
		for x := 0; x < img.Width; x++ {
			//filmFramebuffer := ACESFilm(img[x+y*width])
			ldrColor := img.Pixels[x+y*img.Width].Pow(1.0/2.2).Clamp(0.0, 1.0)

			i.Set(x, y, color.NRGBA{
				R: uint8(255 * ldrColor.X),
				G: uint8(255 * ldrColor.Y),
				B: uint8(255 * ldrColor.Z),
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
