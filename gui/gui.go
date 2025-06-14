package gui

import (
	"fmt"

	"github.com/pwiecz/go-fltk"
)

func Run_gui () error {
	err := get_screen_size()
	if err != nil {
		return err
	}

	fmt.Println(screen_width)
	fmt.Println(screen_height)
	create_window(screen_width, screen_height)

	fltk.Run()
	return nil
}

func create_window (screen_width, screen_height int) error {
	window_width = calc_screen_percentage_width(WINDOW_WIDTH_PERCENTAGE)
	window_height = calc_screen_percentage_height(WINDOW_HEIGHT_PERCENTAGE)

	win := fltk.NewWindow(
		window_width,
		window_height,
	)

	fmt.Println((float32) (100 - WINDOW_WIDTH_PERCENTAGE) / 2.0)
	fmt.Println((float32) (100 - WINDOW_HEIGHT_PERCENTAGE) / 2.0)
	win.SetPosition(
		calc_screen_percentage_width((float32) (100 - WINDOW_WIDTH_PERCENTAGE) / 2.0),
		calc_screen_percentage_height((float32) (100 - WINDOW_HEIGHT_PERCENTAGE) / 2.0) + WINDOW_DECORATION_HEIGTH,
	)

	win.End()
	win.Show()

	return nil
}
