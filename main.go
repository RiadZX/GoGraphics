package main

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
	"math"
)

type MyGraphicContext struct {
	*draw2dimg.GraphicContext
}

func main() {
	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(0, 0, 300, 300))
	gc := MyGraphicContext{draw2dimg.NewGraphicContext(dest)}

	// Set some properties
	gc.SetFillColor(color.RGBA{R: 0x44, G: 0xff, B: 0x44, A: 0xff})
	gc.SetStrokeColor(color.RGBA{R: 0xff, G: 0, B: 0x44, A: 0xff})
	gc.SetLineWidth(5)

	//p0 := Point{0, 100}
	//p1 := Point{300, 100}
	var t float64 = 0.0
	for t <= 1.0 {
		p := gc.Lerp(Point{0, 100}, Point{300, 100}, t)
		gc.SetPoint(p, 1)
		t += 0.1
	}

	gc.Close()

	// Save to file
	draw2dimg.SaveToPngFile("hello.png", dest)
}

func (gc *MyGraphicContext) SetPoint(point Point, width float64) {
	gc.BeginPath()
	gc.ArcTo(point.x, point.y, width, width, 0, 2*math.Pi)
	gc.Fill()
}

//line between two points
func (gc *MyGraphicContext) SetLine(p0, p1 Point) {
	//line should be drawn through both points, extending beyond them
	gc.BeginPath()
	gc.MoveTo(p0.x, p0.y)
	gc.LineTo(p1.x, p1.y)
	gc.Stroke()
}

func (gc *MyGraphicContext) Lerp(p0, p1 Point, t float64) Point {
	return Point{
		x: p0.x + (p1.x-p0.x)*t,
		y: p0.y + (p1.y-p0.y)*t,
	}
}

type Point struct {
	x, y float64
}
