package main

import (
	"fmt"
	"image/color"

	"github.com/llgcode/draw2d/draw2dpdf"
)

func main() {
	fmt.Println("Ulam Spiral in Go v0.0.1")
	printPrimes(2045)
	dest := draw2dpdf.NewPdf("L", "mm", "A4")
	gc := draw2dpdf.NewGraphicContext(dest)
	// Set some properties
	gc.SetFillColor(color.RGBA{0x99, 0x99, 0x99, 0xff})
	gc.SetStrokeColor(color.RGBA{0x00, 0x00, 0x00, 0xff})
	gc.SetLineWidth(1)

	// Draw a closed shape
	gc.MoveTo(10, 10) // should always be called first for a new path
	gc.LineTo(100, 50)
	gc.QuadCurveTo(100, 10, 10, 10)
	gc.Close()
	gc.FillStroke()

	// Save to file
	draw2dpdf.SaveToPdfFile("hello.pdf", dest)
}

func checkPrime(num int) bool {
	var mod int = 0
	var isPrime bool = false

	for i := 1; i <= num; i++ {
		if num%i == 0 {
			mod += 1
		}
	}
	if mod == 2 {
		isPrime = true
	}
	return isPrime

}

func printPrimes(max int) {
	for n := 1; n <= max; n++ {
		if checkPrime(n) == false {
			fmt.Printf("◉ ")
		} else if checkPrime(n) == true {
			fmt.Printf("◌ ")
		}
	}
}
