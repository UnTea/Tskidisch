package main

import (
	"github.com/UnTea/Tskidisch/linmath"
	"github.com/UnTea/Tskidisch/loader"
	"github.com/UnTea/Tskidisch/render"
)

func main() {
	primitives := []render.Primitive{
	render.NewSphere(
		linmath.Vector{Z: 1.0},
		0.35,
		linmath.Vector{X: 1.0, Y: 1.0, Z: 1.0})}
	//primitives[1] = render.NewPlane(
	//	linmath.Vector{Y: -0.35, Z: 1.0},
	//	linmath.Vector{Y: 1.0},
	//	linmath.Vector{X: 0.69, Y: 0.420, Z: 0.228})

	//environmentMap := loader.LoadHDR("wooden_lounge_1k.hdr")
	environmentMap := loader.LoadHDR("comfy_cafe_16k.hdr")
	output := render.Render(primitives, environmentMap)
	output.Save("image.png")
}
