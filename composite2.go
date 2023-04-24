// Port of http://members.shaw.ca/el.supremo/MagickWand/reflect.htm to Go
package main

import "gopkg.in/gographics/imagick.v3/imagick"

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	lw := imagick.NewMagickWand()
	rw := imagick.NewMagickWand()

	mw.ReadImage("composite-layer-3.jpg")
	lw.ReadImage("composite-layer-1.png")
	rw.ReadImage("composite-layer-2.png")

	// mw.ReadImage("composite-layer-1.png")
	// lw.ReadImage("composite-layer-3.jpg")

	// Note that MagickSetImageCompose is usually only used for the MagickMontageImage
	// function and isn't used or needed by MagickCompositeImage
	mw.CompositeImage(lw, imagick.COMPOSITE_OP_PLUS, true, 0, 0)
	/* Write the new image */
	//mw.WriteImage("composite-layer-result.png")
	//pw.ReadImage("composite-layer-result.png")
	mw.CompositeImage(rw, imagick.COMPOSITE_OP_PLUS, true, 0, 0)
	mw.WriteImage("composite-layer-result-final.png")
}
