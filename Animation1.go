package main

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
	"math"
)

// SimpleRope returns a still frame at time t of a simple rope
func SimpleRope(image *image.RGBA, t float64) *image.RGBA {
	//create the gc
	gc := MyGraphicContext{draw2dimg.NewGraphicContext(image)}
	width := image.Bounds().Size().X
	height := image.Bounds().Size().Y

	//top left pivot.
	p0 := Point{0, 0}
	//top right
	p1 := Point{float64(width), 0}
	//the rope end
	x := 0.0
	for x <= 1.001 {
		//use hue to create rainbow effect
		r, g, b := hueToRGB(x)
		rgb := color.RGBA{R: uint8(r * 255), G: uint8(g * 255), B: uint8(b * 255), A: 0xff}
		//draw the bezier curve
		ropePoint := Point{float64(width)/2.0 + math.Cos(t)*float64(width/2), float64(height)}
		point := gc.BezierCurve3(p0, ropePoint, p1, x)
		gc.SetPoint(point, 2)
		gc.SetFillColor(rgb)
		x += 0.001
	}

	gc.SetPoint(p0, 3)
	gc.SetPoint(p1, 3)
	return image
}
