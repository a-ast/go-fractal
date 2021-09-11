package main

import (
	"image/color"
	"testing"
)

func TestNewGradientPalette(t *testing.T) {

	color1 := color.RGBA{0, 0, 0, 0}
	color2 := color.RGBA{100, 100, 100, 100}

	palette := NewGradientPalette(1000, color1, color2)

	got := palette.GetColor(500)
	want := color.RGBA{50, 50, 50, 50}

	if got != want {
		t.Fatalf(`Got %q, want %q`, got, want)
	}

}
