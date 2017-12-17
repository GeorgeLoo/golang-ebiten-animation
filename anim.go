


/*

anim.go by George Loo 15.12.2017



jj
*/


package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
	"path/filepath"
	"image/color"
	"image"

)

const (
	version = "0.1"
	datafolder = "data"
	screenwidth = 800
	screenheight = 400
)

var (
  mousedownState bool
  animseq *ebiten.Image
)


func mouseLeftdown() {
	fmt.Print("mousedown \n")
}

func mouseLeftup() {
	fmt.Print("mouseup \n")
}


func update(screen *ebiten.Image) error {

	if ebiten.IsRunningSlowly() {
		return nil
		//fmt.Print("running slowly! \n")
	}

	screen.Fill(color.NRGBA{255, 255, 0, 0xff})  // yellow

  	if mousedownState {
		if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			mousedownState = false
			mouseLeftup()
		}
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !mousedownState {
			mousedownState = true
			mouseLeftdown()
		}
	}

	draw(screen, animseq, 100, 150)

  	return nil

}

func draw(screen *ebiten.Image, image2draw *ebiten.Image, x float64, y float64) {
	//w, h := image.Size()
	//fmt.Printf("w %d h %d \n",w,h)
	var r image.Rectangle 
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Reset()
	//opts.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	//opts.GeoM.Rotate(float64(l.pointedDir % 360) * 2 * math.Pi / 360)
	//opts.GeoM.Scale( 1.0, 1.0 )
	opts.GeoM.Scale( 1.0, 1.0 )
	opts.GeoM.Translate(x, y)

	r = image.Rect(0, 0, 50, 50)
	opts.SourceRect = &r

	screen.DrawImage(image2draw, opts)

}


func readimg(fn string) *ebiten.Image {
	var err error
	var fname string
	fname = filepath.Join(datafolder, fn)
	img, _, err := ebitenutil.NewImageFromFile(
		fname,
		ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}
	return img

}

func initprog() {

	animseq = readimg("sixframes.png")
}


func main() {

	initprog()

    scale := 1.0
    // Initialize Ebiten, and loop the update() function
    if err := ebiten.Run(update, screenwidth, screenheight, scale, "Animation test 0.0 by George Loo"); err != nil {
      panic(err)
    }
    fmt.Printf("Program ended -----------------\n")

}
