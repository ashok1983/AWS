package main

import (
    //"github.com/gographics/imagick/imagick"
    "fmt"

    //"gopkg.in/gographics/imagick.v2/imagick"
    "gopkg.in/gographics/imagick.v3/imagick"
)

func main() {
    imagick.Initialize()
    defer imagick.Terminate()

    mw := imagick.NewMagickWand()
    defer mw.Destroy()
    ret, err := imagick.ConvertImageCommand([]string{
      //  "convert", "/tmp/test.jpeg", "-resize", "400x400", "/tmp/test.gif",
       "convert", "-size",  "100x100",  "canvas:khaki",  "canvas_khaki.gif", 
    })
    if err != nil {
        panic(err)
    }
    fmt.Printf("Meta: %+v", ret)
    fmt.Println("hello.... world")
}

// make-tile creates a tileable image from an input image.
// ( +clone -flop ) +append  ( +clone -flip ) -append -resize 50%
// func make_tile(mw *imagick.MagickWand, outfile string) {
//     mwc := mw.Clone()
//     mwc.FlopImage()
//     mw.AddImage(mwc)
//     mwc.Destroy()
//     mwc = mw.AppendImages(false)
//     mwf := mwc.Clone()
//     mwf.FlipImage()
//     mwc.AddImage(mwf)
//     mwf.Destroy()
//     mwf = mwc.AppendImages(true)

//     w := mwf.GetImageWidth()
//     h := mwf.GetImageHeight()
//     // 1 = Don't blur or sharpen image
//     mwf.ResizeImage(w/2, h/2, imagick.FILTER_LANCZOS)
//     mwf.WriteImage(outfile)
//     mwf.Destroy()
//     mwc.Destroy()
// }

// func main() {
//     imagick.Initialize()
//     defer imagick.Terminate()

//     mw := imagick.NewMagickWand()
//     mw.SetSize(200, 200)
//     mw.ReadImage("plasma:red-blue")
//     make_tile(mw, "tile_plasma.png")
//     mw.Destroy()

//     mw = imagick.NewMagickWand()
//     mw.SetSize(100, 100)
//     mw.ReadImage("xc:")
//     mw.AddNoiseImage(imagick.NOISE_RANDOM, 0.0)
//     mw.SetImageVirtualPixelMethod(imagick.VIRTUAL_PIXEL_TILE)
//     mw.BlurImage(0, 10)
//     mw.NormalizeImage()
//     make_tile(mw, "tile_random.png")
//     mw.Destroy()
// }
