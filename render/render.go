package render

import (
	"github.com/UnTea/Tskidisch/linmath"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
)

const width, height, sampleCount int = 1024, 768, 16
const fov float64 = 90

func ACESFilm(x linmath.Vector) linmath.Vector {
	a := 2.51
	b := linmath.Splat(0.03)
	c := 2.43
	d := linmath.Splat(0.59)
	e := linmath.Splat(0.14)

	nominator := linmath.Mul(x, linmath.Add(linmath.MulOnScalar(x, a), b))
	denominator := linmath.Add(linmath.Mul(x, linmath.Add(linmath.MulOnScalar(x, c), d)), e)

	return linmath.Div(nominator, denominator).Clamp(0.0, 1.0)
}

func Render(spheres []Sphere) {
	framebuffer := make([]linmath.Vector, width*height)
	aspectRatio := float64(width) / float64(height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			u := 2.0*(float64(x)+rand.Float64())/float64(width) - 1.0
			v := -(2.0*(float64(y)+rand.Float64())/float64(height) - 1.0)
			filmU := u * math.Tan(linmath.Radians(fov)/2) * aspectRatio
			filmV := v * math.Tan(linmath.Radians(fov)/2)
			direction := linmath.Vector{X: filmU, Y: filmV, Z: 1.0}.Norm()
			ray := Ray{Direction: direction, Origin: linmath.Vector{}}

			var minT float64 = math.MaxFloat64
			var index int

			for i := 0; i < len(spheres); i++ {
				t := spheres[i].RayIntersect(ray)

				if t == -1.0 {
					continue
				}

				if t < minT {
					minT, index = t, i
				}
			}

			var color linmath.Vector

			if minT == math.MaxFloat64 {
				color = linmath.Splat(1)
			} else {
				color = spheres[index].Normal(ray.PointAt(minT))
				color = linmath.Add(linmath.MulOnScalar(color, 0.5), linmath.Splat(0.5))
			}

			framebuffer[x+y*width] = color
		}
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			//filmFramebuffer := ACESFilm(framebuffer[x+y*width])
			filmFramebuffer := framebuffer[x+y*width]

			img.Set(x, y, color.NRGBA{
				R: uint8(255 * filmFramebuffer.X),
				G: uint8(255 * filmFramebuffer.Y),
				B: uint8(255 * filmFramebuffer.Z),
				A: 255,
			})
		}
	}

	f, err := os.Create("image.png")
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
