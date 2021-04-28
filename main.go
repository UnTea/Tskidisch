package main

import (
	"github.com/UnTea/Tskidisch/linmath"
	"github.com/UnTea/Tskidisch/render"
)

func main() {
	primitives := make([]render.Primitive, 2)
	primitives[0] = render.NewSphere(
		linmath.Vector{Z: 1.0},
		0.35,
		linmath.Vector{X: 0.769, Y: 0.54, Z: 0.11})
	primitives[1] = render.NewSphere(
		linmath.Vector{Y: -20.4, Z: 1.0},
		20.0,
		linmath.Vector{X: 0.409, Y: 0.24, Z: 0.81})

	render.Render(primitives)
}
