


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


type AnimationType struct {
	sequence *ebiten.Image 
	x, y float64 
	width, height int  // of one frame
	scale float64
	looping bool
	run bool
	numFrames int 
	startAt int  // which frame
	speed int   // 60 is one second 
	numberOfPlays int 
	count int 
	currF int 

}

var (
  mousedownState bool
  //animseq *ebiten.Image
  anim1 AnimationType

)

func (a *AnimationType) init(name string, w int, h int, numF int) {
	a.sequence = readimg(name)
	a.count = 0
	a.run = false
	a.width = w
	a.height = h 
	a.startAt = 0
	a.numFrames = numF

}

func (a *AnimationType) animate(screen *ebiten.Image, x float64, y float64) {
	var x1, y1, x2, y2 int

	a.count++
	if a.count < a.speed {
		return
	}
	a.count = 0
	fmt.Print(a.count," anim \n")
	x1 = a.currF * a.width
	y1 = a.currF * a.height
	x2 = x1 + a.width
	y2 = y1 + a.height
	r := image.Rect(x1,y1,x2,y2)

	draw(screen, a.sequence, x, y, r)
	a.currF += 1
	if a.currF > a.numFrames - 1 {
		a.currF = 0
	}

}

func (a *AnimationType) plate() {

}

//////////////////////////////////////////////////

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

	//draw(screen, animseq, 100, 150)

	anim1.animate(screen, 200, 150)


  	return nil

}

func draw(screen *ebiten.Image, image2draw *ebiten.Image, x float64, y float64, r image.Rectangle) {
	//w, h := image.Size()
	//fmt.Printf("w %d h %d \n",w,h)
	//var r image.Rectangle 
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Reset()
	//opts.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	//opts.GeoM.Rotate(float64(l.pointedDir % 360) * 2 * math.Pi / 360)
	//opts.GeoM.Scale( 1.0, 1.0 )
	opts.GeoM.Scale( 1.0, 1.0 )
	opts.GeoM.Translate(x, y)

	//r := image.Rect(0, 0, 50, 50)
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

	//animseq = readimg("sixframes.png")
	anim1.init("sixframes.png", 50,50, 6)
	anim1.looping = true
	anim1.speed = 60
	anim1.run = true
	anim1.currF = 0


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
