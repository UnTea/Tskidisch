package render

import (
	"github.com/UnTea/Tskidisch/linmath"
	"math"
	"math/rand"
	"sync"
)

const (
	width       int     = 1024
	height      int     = 780
	sampleCount int     = 32
	fieldOfView float64 = 120
)

func Render(primitive []Primitive, environmentMap Image) Image {
	image := NewImage(width, height)
	var wg sync.WaitGroup

	for lineNumber := 0; lineNumber < image.Height/100; lineNumber++ {
		wg.Add(1)
		go Tile(primitive, environmentMap, &wg, &image, lineNumber)
	}

	wg.Wait()
	return image
}

func Tile(primitive []Primitive, environmentMap Image, wg *sync.WaitGroup, image *Image, lineNumber int) {
	aspectRatio := float64(image.Width) / float64(image.Height)
	random := rand.New(rand.NewSource(69))
	defer wg.Done()

	for line := 0; line < 100; line++ {
		y := lineNumber*100 + line
		for x := 0; x < image.Width; x++ {
			var sum linmath.Vector

			for i := 0; i < sampleCount; i++ {
				u := 2.0*(float64(x)+random.Float64())/float64(image.Width) - 1.0
				v := -(2.0*(float64(y)+random.Float64())/float64(image.Height) - 1.0)

				filmU := u * math.Tan(linmath.Radians(fieldOfView)/2) * aspectRatio
				filmV := v * math.Tan(linmath.Radians(fieldOfView)/2)

				direction := linmath.Vector{X: filmU, Y: filmV, Z: 1.0}.Norm()
				ray := Ray{Direction: direction, Origin: linmath.Vector{}}

				color := TraceRay(primitive, ray, environmentMap, random)
				sum = sum.Add(color)
			}
			image.SetPixel(x, y, sum.DivOnScalar(float64(sampleCount)))
		}
	}
}
