package render

import (
	"github.com/UnTea/Tskidisch/linmath"
	"math"
	"math/rand"
)

const width, height, sampleCount int = 1024, 768, 16
const fov float64 = 120

func Render(primitive []Primitive, environmentMap Image) Image {
	image := NewImage(width, height)
	aspectRatio := float64(image.Width) / float64(image.Height)

	for y := 0; y < image.Height; y++ {
		for x := 0; x < image.Width; x++ {
			var sum linmath.Vector

			for i := 0; i < sampleCount; i++ {

				u := 2.0*(float64(x)+rand.Float64())/float64(image.Width) - 1.0
				v := -(2.0*(float64(y)+rand.Float64())/float64(image.Height) - 1.0)

				filmU := u * math.Tan(linmath.Radians(fov)/2) * aspectRatio
				filmV := v * math.Tan(linmath.Radians(fov)/2)

				direction := linmath.Vector{X: filmU, Y: filmV, Z: 1.0}.Norm()
				ray := Ray{Direction: direction, Origin: linmath.Vector{}}

				color := TraceRay(primitive, ray, environmentMap)
				sum = sum.Add(color)
			}
			image.SetPixel(x, y, sum.DivOnScalar(float64(sampleCount)))
		}
	}
	return image
}
