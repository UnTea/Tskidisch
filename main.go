package main

import (
	"github.com/UnTea/Tskidisch/linmath"
	"github.com/UnTea/Tskidisch/render"
)

func main()  {
	render.Render(render.Sphere{Center: linmath.Vector{Z: -1.0}, Radius: 0.35})
}
