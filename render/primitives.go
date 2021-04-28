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

type Plane struct {
	point  linmath.Vector
	normal linmath.Vector
	albedo linmath.Vector
}

func NewSphere(center linmath.Vector, radius float64, albedo linmath.Vector) Sphere {
	return Sphere{
		center: center,
		albedo: albedo,
		radius: radius,
	}
}

func NewPlane(point, normal, albedo linmath.Vector) Plane {
	return Plane{
		point:  point,
		normal: normal.Norm(),
		albedo: albedo,
	}
}

func (plane Plane) Albedo() linmath.Vector {
	return plane.albedo
}

func (plane Plane) Normal(intersection linmath.Vector) linmath.Vector {
	return plane.normal
}

func (plane Plane) RayIntersect(ray Ray) float64 {
	denominator := linmath.Dot(plane.normal, ray.Direction)

	if math.Abs(denominator) > linmath.Epsilon {
		t := linmath.Dot(linmath.Sub(plane.point, ray.Origin), plane.normal) / denominator

		if t >= linmath.Epsilon {
			return t
		}
	}

	return -1.0
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
