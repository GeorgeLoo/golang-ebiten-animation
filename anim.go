


/*

anim.go by George Loo 15.12.2017




*/


package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
)

const (
	screenwidth = 800
	screenheight = 400
)

var (
  mousedownState bool
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

  return nil

}

func main() {

    scale := 1.0
    // Initialize Ebiten, and loop the update() function
    if err := ebiten.Run(update, screenwidth, screenheight, scale, "Animation test 0.0 by George Loo"); err != nil {
      panic(err)
    }
    fmt.Printf("Program ended -----------------\n")

}
