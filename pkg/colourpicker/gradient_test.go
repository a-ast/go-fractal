package colourpicker

import (
	"testing"
)

func TestNewGradientPalette(t *testing.T) {

	middle := Colour{100, 100, 100}
	end := Colour{200, 200, 200}

	picker := NewGradientPicker(1000, []GradientPoint{
		{50, middle},
		{100, end},
	})

	got := picker.Pick(250)
	want := Colour{50, 50, 50}

	if got != want {
		t.Fatalf(`Got %q, want %q`, got, want)
	}
}
