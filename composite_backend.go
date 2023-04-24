// Port of http://members.shaw.ca/el.supremo/MagickWand/reflect.htm to Go
package main

import "gopkg.in/gographics/imagick.v3/imagick"

func main() {
	imagick.Initialize()
	defer imagick.Terminate()
	dw := imagick.NewDrawingWand()
	defer dw.Destroy()
	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	lw := imagick.NewMagickWand()
	defer lw.Destroy()
	pw := imagick.NewPixelWand()
	defer pw.Destroy()
	rw := imagick.NewMagickWand()
	defer rw.Destroy()
	mw.ReadImage("background.png")
	lw.ReadImage("watermark.png")
	rw.ReadImage("icon.png")

	mw.CompositeImage(lw, imagick.COMPOSITE_OP_PLUS, true, 0, 0)
	mw.CompositeImage(rw, imagick.COMPOSITE_OP_PLUS, true, 0, 0)
	//Text
	dw.SetFont("/Users/ashok/code/image/fonts/NotoSans-SemiBold.ttf")
	pw.SetColor("white")
	dw.SetFontSize(48)
	dw.SetTextAlignment(imagick.ALIGN_CENTER)
	dw.SetFillColor(pw)
	message := "KILLS"
	dw.Annotation(400, 310, message)
	mw.DrawImage(dw)

	message = "12.3 K"
	pw.SetColor("#FFC125")
	//pw.SetColor("rgb(238, 201, 0)")
	dw.SetFontSize(192)
	dw.SetFillColor(pw)
	dw.Annotation(400, 490, message)
	// Draw the image on to the mw
	mw.DrawImage(dw)
	mw.WriteImage("backend-created.png")
}
