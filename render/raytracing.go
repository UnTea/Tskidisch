package render

import (
	"github.com/UnTea/Tskidisch/linmath"
	"math"
	"math/rand"
)

type Sphere struct {
	Center linmath.Vector
	Radius float64
	Albedo linmath.Vector
}

type Ray struct {
	Direction linmath.Vector
	Origin    linmath.Vector
}

func (ray Ray) PointAt(t float64) linmath.Vector {
	return linmath.Add(ray.Origin, linmath.MulOnScalar(ray.Direction, t))
}

func (sphere Sphere) Normal(intersection linmath.Vector) linmath.Vector {
	return linmath.Sub(intersection, sphere.Center).Norm()
}

func (sphere Sphere) RayIntersect(ray Ray) float64 {
	oc := linmath.Sub(ray.Origin, sphere.Center)
	b := linmath.Dot(oc, ray.Direction)
	c := linmath.Dot(oc, oc) - sphere.Radius*sphere.Radius
	h := b*b - c

	if h < 0.0 {
		return -1.0 // no intersection
	}

	h = math.Sqrt(h)

	if -b-h > linmath.Epsilon {
		return -b - h // t is -b -h
	}

	if -b+h > linmath.Epsilon {
		return -b + h
	}

	return -1.0
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

func TraceRay(spheres []Sphere, ray Ray) linmath.Vector {
	sphere, t := FindIntersection(spheres, ray)

	if t == math.MaxFloat64 {
		return linmath.Splat(1)
	}

	ray = Ray{
		Direction: RandomUnitVectorInHemisphere(sphere.Normal(ray.PointAt(t))),
		Origin:    ray.PointAt(t),
	}

	color := linmath.Mul(sphere.Albedo, TraceRay(spheres, ray))

	return color
}
