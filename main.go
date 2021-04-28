package main

import (
	"github.com/UnTea/Tskidisch/linmath"
	"github.com/UnTea/Tskidisch/render"
)

func main()  {
	spheres := make([]render.Sphere, 2)
	spheres[0] = render.Sphere{Center: linmath.Vector{Z: 1.0}, Radius: 0.35}
	spheres[1] = render.Sphere{Center: linmath.Vector{Y: -20.4, Z: 1.0}, Radius: 20.0}
	render.Render(spheres)
}
