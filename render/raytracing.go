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

func RandomUnitVectorInHemisphere(normal linmath.Vector, random *rand.Rand) linmath.Vector {
	var randomVector linmath.Vector

	for {
		randomVector = linmath.Vector{
			X: random.Float64(),
			Y: random.Float64(),
			Z: random.Float64(),
		}
		randomVector = randomVector.MulOnScalar(2.0).Add(linmath.Splat(-1.0))

		if randomVector.Dot(randomVector) > 1.0 {
			continue
		}

		if randomVector.Dot(normal) >= 0.0 {
			return randomVector.Norm()
		}
		return randomVector.Negative().Norm()
	}
}

func TraceRay(primitives []Primitive, ray Ray, environmentMap Image, random *rand.Rand) linmath.Vector {
	primitive, t := FindIntersection(primitives, ray)

	if t == math.MaxFloat64 {
		phi := math.Atan2(ray.Direction.Z, ray.Direction.X)
		omega := math.Sqrt((ray.Direction.X * ray.Direction.X) + (ray.Direction.Z * ray.Direction.Z))
		theta := math.Atan2(ray.Direction.Y, omega)

		return environmentMap.GetPixelBySphericalCoordinates(phi, theta)
	}

	ray = Ray{
		Direction: RandomUnitVectorInHemisphere(primitive.Normal(ray.PointAt(t)), random),
		Origin:    ray.PointAt(t),
	}

	color := primitive.Albedo().Mul(TraceRay(primitives, ray, environmentMap, random))

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
