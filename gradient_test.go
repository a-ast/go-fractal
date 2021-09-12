package main

import (
	"testing"
)

func TestNewGradientPalette(t *testing.T) {

	start := Colour{0, 0, 0}
	middle := Colour{100, 100, 100}
	end := Colour{200, 200, 200}

	palette := NewGradientPalette(1000, start, end, []GradientPoint{{50, middle}})

	got := palette.GetColor(250)
	want := Colour{50, 50, 50}

	if got != want {
		t.Fatalf(`Got %q, want %q`, got, want)
	}

}
