package main

import (
	"flag"
	"fmt"
	"github.com/UnTea/Tskidisch/linmath"
	"github.com/UnTea/Tskidisch/loader"
	"github.com/UnTea/Tskidisch/render"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)

		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}

		defer f.Close() // error handling omitted for example

		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}

		defer pprof.StopCPUProfile()
	}

	start := time.Now()
	primitives := []render.Primitive{
		render.NewSphere(
			linmath.Vector{Z: 1.0},
			0.35,
			linmath.Vector{X: 1.0, Y: 1.0, Z: 1.0})}
	//render.NewPlane(
	//	linmath.Vector{Y: -0.35, Z: 1.0},
	//	linmath.Vector{Y: 1.0},
	//	linmath.Vector{X: 0.69, Y: 0.420, Z: 0.228})

	//environmentMap := loader.LoadHDR("wooden_lounge_1k.hdr")
	environmentMap := loader.LoadHDR("comfy_cafe_16k.hdr")
	output := render.Render(primitives, environmentMap)
	output.Save("image.png")
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("Time elapsed: %v\n", elapsed)
}
