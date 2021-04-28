package render

import (
	"github.com/UnTea/Tskidisch/linmath"
	"math"
)

type Primitive interface {
	Normal(intersection linmath.Vector) linmath.Vector
	RayIntersect(ray Ray) float64
	Albedo() linmath.Vector
}

type Sphere struct {
	center linmath.Vector
	radius float64
	albedo linmath.Vector
}

func NewSphere(center linmath.Vector, radius float64, albedo linmath.Vector) Sphere {
	return Sphere{
		center: center,
		albedo: albedo,
		radius: radius,
	}
}

func (sphere Sphere) Albedo() linmath.Vector {
	return sphere.albedo
}

func (sphere Sphere) Normal(intersection linmath.Vector) linmath.Vector {
	return linmath.Sub(intersection, sphere.center).Norm()
}

func (sphere Sphere) RayIntersect(ray Ray) float64 {
	oc := linmath.Sub(ray.Origin, sphere.center)
	b := linmath.Dot(oc, ray.Direction)
	c := linmath.Dot(oc, oc) - sphere.radius*sphere.radius
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