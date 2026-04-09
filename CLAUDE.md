# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Build and run the HTTP server
cd cmd/server && go build && ./server
# Server listens on :10000

# Build and run the file generator (outputs img/fractal.png)
cd cmd/file-generator && go build && ./file-generator

# Run all tests
go test ./...

# Run a single test file
go test ./cmd/server/
go test ./pkg/colourpicker/
```

## Architecture

Two entry points, one shared library:

- **`cmd/server/`** - HTTP server that renders fractals as PNG on each request. Query params control fractal type, dimensions, zoom, center, color palette, and fractal-specific params. `query.go` parses URL params with defaults; `createfractal.go` maps params to a `CreateFractal` struct; `fractalfactory.go` instantiates the correct `fractals.Fractal`; `main.go` renders and writes the PNG response.

- **`cmd/file-generator/`** - CLI tool that renders a fractal to a local PNG file (`img/fractal.png`). Hardcoded parameters.

- **`pkg/fractals/`** - Core fractal math. `base.go` defines the `Fractal` interface (`Render() *FractalElements`), `Canvas`, `Element`, and coordinate mapping logic. Each fractal type (Julia, Mandelbrot, BurningShip) is its own file implementing `Render()`. `ImageFactory.FromItems()` converts `FractalElements` to an `*image.RGBA`.

- **`pkg/colourpicker/`** - Color palette system. `GradientPicker` linearly interpolates between `GradientPoint`s across a fixed-size color array. `predefined_pickers.go` defines named palettes (`arcticsun`, `electro`). `PickerByName()` selects a palette by string.

## Key data flow

```
URL query params
  → Query (parsed, typed, with defaults)
  → CreateFractal struct
  → fractals.Fractal (Julia/Mandelbrot/BurningShip)
  → Render() → *FractalElements ([]Element{X, Y, Value})
  → ImageFactory.FromItems() → *image.RGBA
  → PNG response
```

`Element.Value` is a float32 in [0, 1] representing iteration ratio (`k / maxIterations`). The color picker maps this value to an RGB color.

## Adding a new fractal type

1. Add a new file in `pkg/fractals/` implementing `Render() *FractalElements`.
2. Register it in `pkg/fractals/factory.go` (for the file-generator CLI).
3. Register it in `cmd/server/fractalfactory.go` (for the HTTP server).
4. Add any new query params to `cmd/server/createfractal.go` and `cmd/server/query.go`.
