package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
)

var (
	width  = 1920
	height = 1080
)
var img = image.NewRGBA(image.Rect(0, 0, width, height))

var R = 300.0
var r = 55.0
var O = -100.0

var a uint8 = 127
var b uint8 = 128
var c uint8 = 127
var colFG = color.RGBA{a, b, c, 255}

func main() {
	var colBG = color.RGBA{255, 255, 250, 255}

	draw.Draw(img, img.Bounds(), &image.Uniform{colBG}, image.Point{}, draw.Src)

	for t := 0.0; t < 100.0; t += 0.00001 {

		x := (R+r)*math.Cos(t) - O*math.Cos(((R+r)/r)*t)
		y := (R+r)*math.Sin(t) - O*math.Sin(((R+r)/r)*t)

		if (x+float64(width)/2) < float64(width)/2.0 && (y+float64(height)/2) <= float64(height)/2.0 {
			// colFG = color.RGBA{0, 255, 255, 255}
			colFG = color.RGBA{204, 25, 250, 255}
		} else if (x+float64(width)/2) < float64(width)/2.0 && (y+float64(height)/2) > float64(height)/2 {
			// colFG = color.RGBA{255, 0, 255, 255}
			colFG = color.RGBA{127, 22, 222, 255}
		} else if (x+float64(width)/2) >= float64(width)/2.0 && (y+float64(height)/2) > float64(height)/2 {
			// colFG = color.RGBA{255, 255, 0, 255}
			colFG = color.RGBA{89, 36, 245, 255}
		} else if (x+float64(width)/2) >= float64(width)/2.0 && (y+float64(height)/2) <= float64(height)/2 {
			// colFG = color.RGBA{127, 128, 127, 255}
			colFG = color.RGBA{22, 27, 222, 255}
		}
		// log.Println(colFG)
		img.Set(int(x)+width/2, int(y)+height/2, colFG)
	}

	f, err := os.Create("draw.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}
