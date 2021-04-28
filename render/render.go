package render

import (
	"github.com/UnTea/Tskidisch/linmath"
	"math"
	"math/rand"
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

func FindIntersection(spheres []Sphere, ray Ray) (Sphere, float64) {
	var minT = math.MaxFloat64
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

	return spheres[index], minT
}

func Render(spheres []Sphere) {
	framebuffer := NewFramebuffer(width, height)
	aspectRatio := float64(framebuffer.Width) / float64(framebuffer.Height)

	for y := 0; y < framebuffer.Height; y++ {
		for x := 0; x < framebuffer.Width; x++ {

			u := 2.0*(float64(x)+rand.Float64())/float64(framebuffer.Width) - 1.0
			v := -(2.0*(float64(y)+rand.Float64())/float64(framebuffer.Height) - 1.0)

			filmU := u * math.Tan(linmath.Radians(fov)/2) * aspectRatio
			filmV := v * math.Tan(linmath.Radians(fov)/2)

			direction := linmath.Vector{X: filmU, Y: filmV, Z: 1.0}.Norm()
			ray := Ray{Direction: direction, Origin: linmath.Vector{}}

			color := TraceRay(spheres, ray)

			framebuffer.SetPixel(x, y, color)
		}
	}
	framebuffer.SaveImage("image.png")
}
