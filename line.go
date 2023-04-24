// Port of http://members.shaw.ca/el.supremo/MagickWand/draw_shapes.htm to Go
package main

import (
	"gopkg.in/gographics/imagick.v3/imagick"
)

func main() {
	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()
	dw := imagick.NewDrawingWand()
	cw := imagick.NewPixelWand()
	// A line from the centre of the circle
	// to the top left edge of the image

	diameter := uint(640)
	cw.SetColor("white")
	mw.NewImage(diameter, diameter, cw)

	cw.SetColor("rgb(0,0,1)")

	dw.SetStrokeColor(cw)
	dw.SetStrokeWidth(2)
	dw.SetStrokeAntialias(true)
	cw.SetColor("green")
	//cw.SetColor("#4000c2")
	dw.SetStrokeColor(cw)

	dw.Line(10, 200, 600, 200)
	mw.DrawImage(dw)

	dw.Line(10, 400, 600, 400)
	mw.DrawImage(dw)

	cw.SetColor("blue")
	dw.SetStrokeColor(cw)
	dw.Line(200, 10, 200, 600)
	mw.DrawImage(dw)

	dw.Line(400, 10, 400, 600)
	mw.DrawImage(dw)
	//cw.SetColor("#4000c2")
	dw.SetStrokeColor(cw)

	mw.WriteImage("line.png")
}
