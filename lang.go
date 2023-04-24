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
	pw.SetColor("white")
	mw.NewImage(800, 200, pw)
	// Set up a 72 point font
	dw.SetFont("/Users/ashok/code/image/fonts/AppleGothic.ttf")
	dw.SetFontSize(40)
	// Set up a 72 point white font
	dw.SetGravity(imagick.GRAVITY_WEST)
	pw.SetColor("rgb(125,215,255)")
	pw.SetColor("green")
	dw.SetFillColor(pw)
	message := "소프트웨어 구매 언제 소프트"
	// Now draw the text
	dw.Annotation(10, 10, message)
	// Draw the image on to the mw
	mw.DrawImage(dw)
	dw.SetFont("/Users/ashok/code/image/fonts/AppleColorEmoji.ttc")
	dw.SetGravity(imagick.GRAVITY_WEST)
	//dw.SetFillColor(pw)
	emoji := "\U0001F920"
	mw.AnnotateImage(dw, 550, 10, 0, emoji)
	emoji = "\U0001F922"
	mw.AnnotateImage(dw, 600, 10, 0, emoji)
	emoji = "\U0001F926"
	mw.AnnotateImage(dw, 650, 10, 0, emoji)
	mw.WriteImage("text_bevel.png")
	return
}
