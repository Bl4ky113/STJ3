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
	window_width = calc_screen_percentage(75, false)
	window_height = calc_screen_percentage(70, true)

	win := fltk.NewWindow(
		window_width,
		window_height,
	)

	win.End()
	win.Show()

	return nil
}
