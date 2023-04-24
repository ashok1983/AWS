// Text effect 1 - shadow effect using MagickShadowImage
// NOTE - if an image has a transparent background, adding a border of any colour other
// than "none" will remove all the transparency and replace it with the border's colour
// Port of http://members.shaw.ca/el.supremo/MagickWand/resize.htm to Go
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
	pw.SetColor("none")
	// Create a new transparent image
	mw.NewImage(350, 100, pw)
	// Set up a 72 point white font
	pw.SetColor("white")
	dw.SetFillColor(pw)
	dw.SetFont("Verdana-Bold-Italic")
	dw.SetFontSize(72)
	// Add a black outline to the text
	pw.SetColor("black")
	dw.SetStrokeColor(pw)
	// Turn antialias on - not sure this makes a difference
	dw.SetTextAntialias(true)
	// Now draw the text
	dw.Annotation(35, 75, "Krafton")
	// Draw the image on to the mw
	mw.DrawImage(dw)
	// Trim the image down to include only the text
	mw.TrimImage(0)
	// equivalent to the command line +repage
	mw.ResetImagePage("")
	// Make a copy of the text image
	cw := mw.Clone()
	// Set the background colour to blue for the shadow
	pw.SetColor("blue")
	mw.SetImageBackgroundColor(pw)
	// Opacity is a real number indicating (apparently) percentage
	mw.ShadowImage(70, 4, 5, 5)
	// Composite the text on top of the shadow
	mw.CompositeImage(cw, imagick.COMPOSITE_OP_OVER, true, 5, 5)
	cw.Destroy()
	cw = imagick.NewMagickWand()
	defer cw.Destroy()
	// Create a new image the same size as the text image and put a solid colour
	// as its background
	pw.SetColor("rgb(125,215,255)")
	cw.NewImage(mw.GetImageWidth(), mw.GetImageHeight(), pw)
	// Now composite the shadowed text over the plain background
	cw.CompositeImage(mw, imagick.COMPOSITE_OP_OVER, true, 0, 0)
	// and write the result
	cw.WriteImage("text_shadow.png")
}
