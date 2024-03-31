package main

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
	"math"
	"math/rand"
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
	gc.SetBackgroundColor(color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff})
	gc.SetStrokeColor(color.RGBA{R: 0xff, G: 0, B: 0x44, A: 0xff})
	gc.SetLineWidth(5)
	gc.DrawBorder(300, 300)
	//p0 := Point{0, 100}
	//p1 := Point{300, 100}
	gc.RandomImage(100)
	//var t = 0.0
	//p0 := Point{0, 0}
	//p1 := Point{0, 300}
	//p2 := Point{500, 300}
	//p3 := Point{150, 0}
	//
	//gc.SetPoint(p0, 5)
	//gc.SetPoint(p1, 5)
	//gc.SetPoint(p2, 5)
	//gc.SetPoint(p3, 5)
	//
	//for t <= 1.001 {
	//	//use hue to create rainbow effect
	//	r, g, b := hueToRGB(t)
	//	rgb := color.RGBA{R: uint8(r * 255), G: uint8(g * 255), B: uint8(b * 255), A: 0xff}
	//	//draw the bezier curve
	//	point := gc.BezierCurve4(p0, p1, p2, p3, t)
	//	gc.SetPoint(point, 2)
	//	gc.SetFillColor(rgb)
	//
	//	t += 0.001
	//}
	//gc.Close()
	// Save to file
	draw2dimg.SaveToPngFile("random1.png", dest)
}

// SetPoint draws a point at the given location
func (gc *MyGraphicContext) SetPoint(point Point, width float64) {
	gc.BeginPath()
	gc.ArcTo(point.x, point.y, width, width, 0, 2*math.Pi)
	gc.Fill()
}

// SetLine draws a line part between two points
func (gc *MyGraphicContext) SetLine(p0, p1 Point) {
	//line should be drawn through both points, extending beyond them
	gc.BeginPath()
	gc.MoveTo(p0.x, p0.y)
	gc.LineTo(p1.x, p1.y)
	gc.Stroke()
}

// Lerp returns the linear interpolation between two points
func (gc *MyGraphicContext) Lerp(p0, p1 Point, t float64) Point {
	return Point{
		x: p0.x + (p1.x-p0.x)*t,
		y: p0.y + (p1.y-p0.y)*t,
	}
}

func (gc *MyGraphicContext) BezierCurve3(p0, p1, p2 Point, t float64) Point {
	return gc.Lerp(gc.Lerp(p0, p1, t), gc.Lerp(p1, p2, t), t)
}

func (gc *MyGraphicContext) BezierCurve4(p0, p1, p2, p3 Point, t float64) Point {
	return gc.Lerp(gc.BezierCurve3(p0, p1, p2, t), gc.BezierCurve3(p1, p2, p3, t), t)
}

// min3 returns the minimum of three values
func min3(a, b, c float64) float64 {
	return math.Min(math.Min(a, b), c)
}

// hueToRGB converts a hue value to an RGB color
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

func (gc *MyGraphicContext) SetBackgroundColor(color color.Color) {
	//draw a rectangle with the background color full size
	gc.SetFillColor(color)
	gc.MoveTo(0, 0)
	gc.LineTo(800, 0)
	gc.LineTo(800, 800)
	gc.LineTo(0, 800)
	gc.Close()
	gc.Fill()
}

func (gc *MyGraphicContext) DrawBorder(width, height float64) {
	gc.BeginPath()
	gc.MoveTo(0, 0)
	gc.LineTo(width, 0)
	gc.LineTo(width, height)
	gc.LineTo(0, height)
	gc.Close()
	gc.Stroke()

}

func (gc *MyGraphicContext) RandomImage(amountOfPoints int) {
	//for i := 0; i < amountOfPoints; i++ {
	//	gc.SetPoint(Point{rand.Float64() * 400, rand.Float64() * 400}, 3)
	//}
	t := 0.0
	//each 4 points define a bezier curve
	for i := 0; i < amountOfPoints; i += 4 {
		p0 := Point{rand.Float64() * 400, rand.Float64() * 400}
		p1 := Point{rand.Float64() * 400, rand.Float64() * 400}
		p2 := Point{rand.Float64() * 400, rand.Float64() * 400}
		p3 := Point{rand.Float64() * 400, rand.Float64() * 400}

		gc.SetPoint(p0, 0)
		gc.SetPoint(p1, 0)
		gc.SetPoint(p2, 0)
		gc.SetPoint(p3, 0)

		for t <= 1.001 {
			//use hue to create rainbow effect
			r, g, b := hueToRGB(t)
			rgb := color.RGBA{R: uint8(r * 255), G: uint8(g * 255), B: uint8(b * 255), A: 0xff}
			//draw the bezier curve
			point := gc.BezierCurve4(p0, p1, p2, p3, t)
			gc.SetPoint(point, 2)
			gc.SetFillColor(rgb)

			t += 0.001
		}
		t = 0.0
	}

	gc.Close()

}
