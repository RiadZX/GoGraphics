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
	gc.SetLineWidth(1)

	//p0 := Point{0, 100}
	//p1 := Point{300, 100}
	var t = 0.0
	p0 := Point{0, 100}
	p2 := Point{300, 100}
	p1 := Point{150, 300}
	for t <= 1.001 {
		p3 := gc.Lerp(p0, p1, t)
		p4 := gc.Lerp(p1, p2, t)
		//p5 := gc.Lerp(p3, p4, t
		//use hue to create rainbow effect
		r, g, b := hueToRGB(t)
		rgb := color.RGBA{R: uint8(r * 255), G: uint8(g * 255), B: uint8(b * 255), A: 0xff}

		gc.SetStrokeColor(rgb)
		gc.BeginPath()
		gc.MoveTo(p3.x, p3.y)
		gc.LineTo(p4.x, p4.y)
		gc.Stroke()

		t += 0.05
	}
	gc.SetPoint(p2, 1)
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

func min3(a, b, c float64) float64 {
	return math.Min(math.Min(a, b), c)
}

func hueToRGB(h float64) (float64, float64, float64) {
	kr := math.Mod(5+h*6, 6)
	kg := math.Mod(3+h*6, 6)
	kb := math.Mod(1+h*6, 6)

	r := 1 - math.Max(min3(kr, 4-kr, 1), 0)
	g := 1 - math.Max(min3(kg, 4-kg, 1), 0)
	b := 1 - math.Max(min3(kb, 4-kb, 1), 0)

	return r, g, b
}

type Point struct {
	x, y float64
}
