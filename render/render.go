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

func Render(sphere Sphere) {
	framebuffer := make([]linmath.Vector, width*height)
	aspectRatio := float64(width) / float64(height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			var summa linmath.Vector

			for i := 0; i < sampleCount; i++ {
				u := 2.0*(float64(x)+rand.Float64())/float64(width) - 1.0
				v := -(2.0*(float64(y)+rand.Float64())/float64(height) - 1.0)

				filmU := u * math.Tan(linmath.Radians(fov)/2) * aspectRatio
				filmV := v * math.Tan(linmath.Radians(fov)/2)

				direction := linmath.Vector{X: filmU, Y: filmV, Z: 1.0}.Norm()
				ray := Ray{Direction: direction, Origin: linmath.Vector{}}
				t := sphere.RayIntersect(ray)
				var color linmath.Vector

				if t == -1.0 {
					color = linmath.Splat(1)
				} else {
					color = sphere.Normal(ray.PointAt(t))
					color = linmath.Add(linmath.MulOnScalar(color, 0.5), linmath.Splat(0.5))
				}

				summa = linmath.Add(summa, color)
			}
			summa = linmath.DivOnScalar(summa, float64(sampleCount))
			framebuffer[x+y*width] = summa
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
