package examples

import (
	"fmt"

	color "github.com/nilmonster/go-color"
)

func ExampleRGBColor() {
	black := color.RGBColor{R: 0, G: 0, B: 0, Name: "Black"}
	white := color.RGBColor{R: 255, G: 255, B: 255, Name: "White"}
	fmt.Println(black)
	fmt.Println(white)
}

func ExampleAppleColor() {
	black := color.AppleColor{R: 0.0, G: 0.0, B: 0.0, Name: "Black"}
	white := color.AppleColor{R: 1.0, G: 1.0, B: 1.0, Name: "White"}
	fmt.Println(black)
	fmt.Println(white)
}
