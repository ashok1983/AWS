package main

import (
	"gopkg.in/gographics/imagick.v3/imagick"
)

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	pw := imagick.NewPixelWand()
	dw := imagick.NewDrawingWand()

	if err := mw.SetSize(1000, 200); err != nil {
		panic(err)
	}
	if err := mw.ReadImage("xc:white"); err != nil {
		panic(err)
	}

	// Now we draw the Drawing wand on to the Magick Wand
	if err := mw.DrawImage(dw); err != nil {
		panic(err)
	}
	// Turn the matte of == +matte
	if err := mw.SetImageMatte(false); err != nil {
		panic(err)
	}
	if err := mw.WriteImage("logo_mask.png"); err != nil {
		panic(err)
	}

	mw.Destroy()
	dw.Destroy()
	pw.Destroy()

	mw = imagick.NewMagickWand()
	pw = imagick.NewPixelWand()
	dw = imagick.NewDrawingWand()

	mwc := imagick.NewMagickWand()

	mw.ReadImage("logo_mask.png")

	dw.Color(0, 0, imagick.PAINT_METHOD_RESET)
	mw.DrawImage(dw)

	mwc.ReadImage("logo_mask.png")
	mwc.SetImageMatte(false)
	mw.CompositeImage(mwc, imagick.COMPOSITE_OP_COPY, true, 0, 0)

	// Annotate gets all the font information from the drawingwand
	// but draws the text on the magickwand
	// Get the first available "*Sans*" font
	//fonts := mw.QueryFonts("*Apple-Color-Emoji*")
	//fmt.Println("Fonts", fonts)

	dw.SetFont("AppleGothic")
	//dw.SetFont(fonts[0])
	dw.SetFontSize(40)
	dw.SetTextAntialias(true)
	pw.SetColor("rgb(125,215,255)")
	pw.SetColor("green")
	dw.SetFillColor(pw)
	dw.SetTextEncoding("UTF-8")
	//dw.SetStrokeColor(pw)
	dw.SetGravity(imagick.GRAVITY_CENTER)
	//pizzaMessage := "'😘'"
	pizzaMessage := "소프트웨어 구매 언제 완료"
	mw.AnnotateImage(dw, 10, 20, 0, pizzaMessage)

	// dw.SetFont("Apple-Color-Emoji")
	// pw.SetColor("blue")
	// dw.SetFillColor(pw)
	// pizzaMessage = "\U0001F920"
	// mw.AnnotateImage(dw, 280, 20, 0, pizzaMessage)

	// pizzaMessage = "\U0001F922"
	// mw.AnnotateImage(dw, 320, 20, 0, pizzaMessage)

	mw.WriteImage("logo_ant.png")
	mwc.Destroy()
}
