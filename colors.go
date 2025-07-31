package chart

import "github.com/HootEgor/go-chart/v2/drawing"

var (
	// ColorWhite is white.
	ColorWhite = drawing.Color{R: 255, G: 255, B: 255, A: 255}

	// ColorBlue is a modern theme blue.
	ColorBlue = drawing.Color{R: 52, G: 152, B: 219, A: 255} // Softer, flat blue

	// ColorCyan is a vibrant theme cyan.
	ColorCyan = drawing.Color{R: 26, G: 188, B: 156, A: 255} // Teal-like cyan

	// ColorGreen is a modern theme green.
	ColorGreen = drawing.Color{R: 46, G: 204, B: 113, A: 255} // Fresh green

	// ColorRed is a modern theme red.
	ColorRed = drawing.Color{R: 231, G: 76, B: 60, A: 255} // Softer red

	// ColorOrange is a modern theme orange.
	ColorOrange = drawing.Color{R: 243, G: 156, B: 18, A: 255} // Warm amber

	// ColorYellow is a softer, warm yellow.
	ColorYellow = drawing.Color{R: 241, G: 196, B: 15, A: 255} // Gold-like yellow

	// ColorBlack is a deeper gray-black.
	ColorBlack = drawing.Color{R: 44, G: 62, B: 80, A: 255} // Charcoal black

	// ColorLightGray is a softer light gray.
	ColorLightGray = drawing.Color{R: 236, G: 240, B: 241, A: 255} // Subtle off-white gray

	// Alternate theme colors

	// ColorAlternateBlue is a soft pastel blue.
	ColorAlternateBlue = drawing.Color{R: 116, G: 185, B: 255, A: 255}

	// ColorAlternateGreen is a mint green.
	ColorAlternateGreen = drawing.Color{R: 85, G: 239, B: 196, A: 255}

	// ColorAlternateGray is a medium cool gray.
	ColorAlternateGray = drawing.Color{R: 149, G: 165, B: 166, A: 255}

	// ColorAlternateYellow is a peachy yellow.
	ColorAlternateYellow = drawing.Color{R: 253, G: 203, B: 110, A: 255}

	// ColorAlternateLightGray is a warm light gray.
	ColorAlternateLightGray = drawing.Color{R: 223, G: 228, B: 234, A: 255}

	// ColorTransparent is a transparent (alpha zero) color.
	ColorTransparent = drawing.Color{R: 1, G: 1, B: 1, A: 0}
)

var (
	// DefaultBackgroundColor is the default chart background color.
	// It is equivalent to css color:white.
	DefaultBackgroundColor = ColorWhite
	// DefaultBackgroundStrokeColor is the default chart border color.
	// It is equivalent to color:white.
	DefaultBackgroundStrokeColor = ColorWhite
	// DefaultCanvasColor is the default chart canvas color.
	// It is equivalent to css color:white.
	DefaultCanvasColor = ColorWhite
	// DefaultCanvasStrokeColor is the default chart canvas stroke color.
	// It is equivalent to css color:white.
	DefaultCanvasStrokeColor = ColorWhite
	// DefaultTextColor is the default chart text color.
	// It is equivalent to #333333.
	DefaultTextColor = ColorBlack
	// DefaultAxisColor is the default chart axis line color.
	// It is equivalent to #333333.
	DefaultAxisColor = ColorBlack
	// DefaultStrokeColor is the default chart border color.
	// It is equivalent to #efefef.
	DefaultStrokeColor = ColorLightGray
	// DefaultFillColor is the default fill color.
	// It is equivalent to #0074d9.
	DefaultFillColor = ColorBlue
	// DefaultAnnotationFillColor is the default annotation background color.
	DefaultAnnotationFillColor = ColorWhite
	// DefaultGridLineColor is the default grid line color.
	DefaultGridLineColor = ColorLightGray
)

var (
	// DefaultColors are a couple default series colors.
	DefaultColors = []drawing.Color{
		ColorBlue,
		ColorGreen,
		ColorRed,
		ColorCyan,
		ColorOrange,
	}

	// DefaultAlternateColors are a couple alternate colors.
	DefaultAlternateColors = []drawing.Color{
		ColorAlternateBlue,
		ColorAlternateGreen,
		ColorAlternateGray,
		ColorAlternateYellow,
		ColorBlue,
		ColorGreen,
		ColorRed,
		ColorCyan,
		ColorOrange,
	}
)

// GetDefaultColor returns a color from the default list by index.
// NOTE: the index will wrap around (using a modulo).
func GetDefaultColor(index int) drawing.Color {
	finalIndex := index % len(DefaultColors)
	return DefaultColors[finalIndex]
}

// GetAlternateColor returns a color from the default list by index.
// NOTE: the index will wrap around (using a modulo).
func GetAlternateColor(index int) drawing.Color {
	finalIndex := index % len(DefaultAlternateColors)
	return DefaultAlternateColors[finalIndex]
}

// ColorPalette is a set of colors that.
type ColorPalette interface {
	BackgroundColor() drawing.Color
	BackgroundStrokeColor() drawing.Color
	CanvasColor() drawing.Color
	CanvasStrokeColor() drawing.Color
	AxisStrokeColor() drawing.Color
	TextColor() drawing.Color
	GetSeriesColor(index int) drawing.Color
}

// DefaultColorPalette represents the default palatte.
var DefaultColorPalette defaultColorPalette

type defaultColorPalette struct{}

func (dp defaultColorPalette) BackgroundColor() drawing.Color {
	return DefaultBackgroundColor
}

func (dp defaultColorPalette) BackgroundStrokeColor() drawing.Color {
	return DefaultBackgroundStrokeColor
}

func (dp defaultColorPalette) CanvasColor() drawing.Color {
	return DefaultCanvasColor
}

func (dp defaultColorPalette) CanvasStrokeColor() drawing.Color {
	return DefaultCanvasStrokeColor
}

func (dp defaultColorPalette) AxisStrokeColor() drawing.Color {
	return DefaultAxisColor
}

func (dp defaultColorPalette) TextColor() drawing.Color {
	return DefaultTextColor
}

func (dp defaultColorPalette) GetSeriesColor(index int) drawing.Color {
	return GetDefaultColor(index)
}

// AlternateColorPalette represents the default palatte.
var AlternateColorPalette alternateColorPalette

type alternateColorPalette struct{}

func (ap alternateColorPalette) BackgroundColor() drawing.Color {
	return DefaultBackgroundColor
}

func (ap alternateColorPalette) BackgroundStrokeColor() drawing.Color {
	return DefaultBackgroundStrokeColor
}

func (ap alternateColorPalette) CanvasColor() drawing.Color {
	return DefaultCanvasColor
}

func (ap alternateColorPalette) CanvasStrokeColor() drawing.Color {
	return DefaultCanvasStrokeColor
}

func (ap alternateColorPalette) AxisStrokeColor() drawing.Color {
	return DefaultAxisColor
}

func (ap alternateColorPalette) TextColor() drawing.Color {
	return DefaultTextColor
}

func (ap alternateColorPalette) GetSeriesColor(index int) drawing.Color {
	return GetAlternateColor(index)
}
