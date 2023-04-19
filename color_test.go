package color

import "testing"

func TestRGBColor(t *testing.T) {
	c := RGBColor{0, 0, 0, "Black"}
	if c.R != 0 || c.G != 0 || c.B != 0 {
		t.Errorf("Triplet color is %v, want %v", []uint8{c.R, c.G, c.B}, []uint8{0, 0, 0})
	}
	if c.Name != "Black" {
		t.Errorf("Color name is %v, want %v", c.Name, "Black")
	}
}

func TestAppleColor(t *testing.T) {
	c := AppleColor{0.0, 0.0, 0.0, "Black"}
	if c.R != 0.0 || c.G != 0.0 || c.B != 0.0 {
		t.Errorf("Triplet color is %v, want %v", []float32{c.R, c.G, c.B}, []float32{0.0, 0.0, 0.0})
	}
	if c.Name != "Black" {
		t.Errorf("Color name is %v, want %v", c.Name, "Black")
	}
}
