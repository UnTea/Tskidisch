package render

import (
	"github.com/UnTea/Tskidisch/linmath"
	"math"
)

type Sphere struct {
	Center linmath.Vector
	Radius float64
}

type Ray struct {
	Direction linmath.Vector
	Origin    linmath.Vector
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
	return -b - h // t is -b -h
}