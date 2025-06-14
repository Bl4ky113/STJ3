package gui

type Number interface {
	~int | ~float64 | ~float32
}

const WINDOW_WIDTH_PERCENTAGE int = 75
const WINDOW_HEIGHT_PERCENTAGE int = 75
const WINDOW_DECORATION_HEIGTH int = 30

var screen_width, screen_height int
var window_width, window_height int

var i int = 0
