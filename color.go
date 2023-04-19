package color

import "fmt"

type RGBColor struct {
	R, G, B uint8
	Name    string
}

func (c RGBColor) String() string {
	return fmt.Sprintf("%v %v %v %v", c.R, c.G, c.B, c.Name)
}

type AppleColor struct {
	R, G, B float32
	Name    string
}

func (c AppleColor) String() string {
	return fmt.Sprintf("%f %f %f %v", c.R, c.G, c.B, c.Name)
}

func RGBColorToAppleColor(c RGBColor) AppleColor {
	r1 := float32(c.R) / 255.0
	g1 := float32(c.G) / 255.0
	b1 := float32(c.B) / 255.0
	return AppleColor{r1, g1, b1, c.Name}
}
