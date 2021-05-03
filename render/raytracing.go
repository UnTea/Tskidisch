package render

import (
	"github.com/UnTea/Tskidisch/linmath"
	"math"
	"math/rand"
)

type Ray struct {
	Direction linmath.Vector
	Origin    linmath.Vector
}

func (ray Ray) PointAt(t float64) linmath.Vector {
	return ray.Origin.Add(ray.Direction.MulOnScalar(t))
}

func RandomUnitVectorInHemisphere(normal linmath.Vector) linmath.Vector {
	var random linmath.Vector

	for {
		random = linmath.Vector{
			X: rand.Float64(),
			Y: rand.Float64(),
			Z: rand.Float64(),
		}
		random = random.MulOnScalar(2.0).Add(linmath.Splat(-1.0))

		if random.Dot(random) > 1.0 {
			continue
		}

		if random.Dot(normal) >= 0.0 {
			return random.Norm()
		}
		return random.Negative().Norm()
	}
}

func TraceRay(primitives []Primitive, ray Ray, environmentMap Image) linmath.Vector {
	primitive, t := FindIntersection(primitives, ray)

	if t == math.MaxFloat64 {
		phi := math.Atan2(ray.Direction.Z, ray.Direction.X)
		omega := math.Sqrt(math.Sqrt(ray.Direction.Dot(ray.Direction)))
		theta := math.Atan2(ray.Direction.Y, omega)
		return environmentMap.GetPixelBySphericalCoordinates(phi, theta)
	}

	ray = Ray{
		Direction: RandomUnitVectorInHemisphere(primitive.Normal(ray.PointAt(t))),
		Origin:    ray.PointAt(t),
	}

	color := primitive.Albedo().Mul(TraceRay(primitives, ray, environmentMap))

	return color
}

func FindIntersection(primitives []Primitive, ray Ray) (Primitive, float64) {
	var minT = math.MaxFloat64
	var index int

	for i := 0; i < len(primitives); i++ {
		t := primitives[i].RayIntersect(ray)

		if t == -1.0 {
			continue
		}

		if t < minT {
			minT, index = t, i
		}
	}

	return primitives[index], minT
}
