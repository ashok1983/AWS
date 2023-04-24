// Port of http://members.shaw.ca/el.supremo/MagickWand/draw_shapes.htm to Go
package main

import (
	"gopkg.in/gographics/imagick.v3/imagick"
)

func main() {
	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	dw := imagick.NewDrawingWand()
	defer dw.Destroy()
	pw := imagick.NewPixelWand()
	defer pw.Destroy()
	// Create a 320x100 canvas
	pw.SetColor("black")
	mw.NewImage(800, 800, pw)
	// Set up a 72 point font
	dw.SetFont("/Users/ashok/code/image/fonts/NotoSans-SemiBold.ttf")
	dw.SetFontSize(48)
	// Set up a 72 point white font
	pw.SetColor("white")
	mw.SetGravity(imagick.GRAVITY_CENTER)
	dw.SetGravity(imagick.GRAVITY_CENTER)

	dw.SetFillColor(pw)

	message := "KILLS"
	dw.Annotation(293, 400, message)
	mw.DrawImage(dw)
	//mw.WriteImage("text_bevel.png")
	//Now draw the text
	message = "12.3 K"
	pw.SetColor("#FFD700")
	dw.SetFontSize(192)
	dw.SetFillColor(pw)
	dw.Annotation(418, 400, message)
	// Draw the image on to the mw
	mw.DrawImage(dw)
	mw.WriteImage("text_bevel.png")
	return
}
