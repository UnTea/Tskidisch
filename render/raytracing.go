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
	return linmath.Add(ray.Origin, linmath.MulOnScalar(ray.Direction, t))
}

func RandomUnitVectorInHemisphere(normal linmath.Vector) linmath.Vector {
	var random linmath.Vector

	for {
		random = linmath.Vector{
			X: rand.Float64(),
			Y: rand.Float64(),
			Z: rand.Float64(),
		}
		random = linmath.Add(linmath.MulOnScalar(random, 2), linmath.Splat(-1))

		if linmath.Dot(random, random) > 1 {
			continue
		}

		if linmath.Dot(random, normal) >= 0 {
			return random.Norm()
		}
		return random.Negative().Norm()
	}
}

func TraceRay(primitives []Primitive, ray Ray) linmath.Vector {
	primitive, t := FindIntersection(primitives, ray)

	if t == math.MaxFloat64 {
		return linmath.Splat(1)
	}

	ray = Ray{
		Direction: RandomUnitVectorInHemisphere(primitive.Normal(ray.PointAt(t))),
		Origin:    ray.PointAt(t),
	}

	color := linmath.Mul(primitive.Albedo(), TraceRay(primitives, ray))

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
