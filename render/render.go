package render

import (
	"github.com/UnTea/Tskidisch/linmath"
	"math"
	"math/rand"
)

const width, height, sampleCount int = 1024, 768, 16
const fov float64 = 90

func Render(primitive []Primitive) {
	framebuffer := NewFramebuffer(width, height)
	aspectRatio := float64(framebuffer.Width) / float64(framebuffer.Height)

	for y := 0; y < framebuffer.Height; y++ {
		for x := 0; x < framebuffer.Width; x++ {
			var sum linmath.Vector

			for i := 0; i < sampleCount; i++ {

				u := 2.0*(float64(x)+rand.Float64())/float64(framebuffer.Width) - 1.0
				v := -(2.0*(float64(y)+rand.Float64())/float64(framebuffer.Height) - 1.0)

				filmU := u * math.Tan(linmath.Radians(fov)/2) * aspectRatio
				filmV := v * math.Tan(linmath.Radians(fov)/2)

				direction := linmath.Vector{X: filmU, Y: filmV, Z: 1.0}.Norm()
				ray := Ray{Direction: direction, Origin: linmath.Vector{}}

				color := TraceRay(primitive, ray)
				sum = linmath.Add(sum, color)
			}
			framebuffer.SetPixel(x, y, linmath.DivOnScalar(sum, float64(sampleCount)))
		}
	}
	framebuffer.SaveImage("image.png")
}
