package gui

import (
	"strconv"

	"github.com/pwiecz/go-fltk"
)

func Run_gui () error {
	err := get_screen_size()
	if err != nil {
		return err
	}

	create_window(screen_width, screen_height)

	fltk.Run()
	return nil
}

var testptr *fltk.Box = nil
var grpptr *fltk.Group = nil

func create_window (screen_width, screen_height int) error {
	window_width = calc_screen_percentage_width(WINDOW_WIDTH_PERCENTAGE)
	window_height = calc_screen_percentage_height(WINDOW_HEIGHT_PERCENTAGE)

	win := fltk.NewWindowWithPosition(
		calc_screen_percentage_width((float32) (100 - WINDOW_WIDTH_PERCENTAGE) / 2.0),
		calc_screen_percentage_height((float32) (100 - WINDOW_HEIGHT_PERCENTAGE) / 2.0) + WINDOW_DECORATION_HEIGTH,
		window_width,
		window_height,
		TITLE_STR,
	)

	window_flex_wrapper := fltk.NewFlex(0, 0, window_width, window_height)
	window_flex_wrapper.SetType(fltk.ROW)
	window_flex_wrapper.SetGap(10)

	menu_wrapper := create_menu_section()
	main_wrapper := create_main_section()
	widgets_wrapper := create_widget_section()

	window_flex_wrapper.End()
	window_flex_wrapper.Fixed(menu_wrapper, calc_window_percentage_width(MENU_WIDTH_PERCENTAGE))
	window_flex_wrapper.Fixed(main_wrapper, calc_window_percentage_width(MAIN_WIDTH_PERCENTAGE))
	window_flex_wrapper.Fixed(widgets_wrapper, calc_window_percentage_width(WIDGETS_WIDTH_PERCENTAGE))

	win.End()
	win.Add(window_flex_wrapper)
	win.Show()

	return nil
}

func create_menu_section () fltk.Widget {
	menu_wrapper := fltk.NewFlex(
		0, 0, 
		calc_window_percentage_width(MENU_WIDTH_PERCENTAGE), window_height,
	)
	menu_wrapper.SetGap(5)

	const btn_size int = 32

	dec := fltk.NewButton(0, 0, btn_size, btn_size, "Show")
	menu_wrapper.Fixed(dec, btn_size)
	inc := fltk.NewButton(0, 0, btn_size, btn_size, "Hide")
	menu_wrapper.Fixed(inc, btn_size)

	inc.SetCallback(func() {
		(*testptr).Hide()
	})
	dec.SetCallback(func() {
		(*testptr).Show()
	})

	menu_wrapper.End()
	return menu_wrapper
}

func create_main_section () fltk.Widget {
	main_width := calc_window_percentage_width(MAIN_WIDTH_PERCENTAGE)
	curr_main_height := window_height

	main_wrapper := fltk.NewFlex(
		calc_window_percentage_width(MENU_WIDTH_PERCENTAGE), 0, 
		main_width, window_height,
	)

	testptr = fltk.NewBox(fltk.DIAMOND_UP_BOX, 0, 0, main_width, curr_main_height, "TEST TESTS")
	testptr.Hide()

	grpptr = fltk.NewGroup(0, 0, main_width, window_height)
	grpptr.Add(testptr)

	main_wrapper.Fixed(grpptr, window_height)

	main_wrapper.End()
	return main_wrapper
}

func create_widget_section () fltk.Widget {
	widgets_wrapper := fltk.NewFlex(
		calc_window_percentage_width(MENU_WIDTH_PERCENTAGE + MAIN_WIDTH_PERCENTAGE), 0, 
		calc_window_percentage_width(WIDGETS_WIDTH_PERCENTAGE), window_height,
	)
	inc := fltk.NewButton(0, 0, 0, 0, "Increment")
	widgets_wrapper.Fixed(inc, 40)
	box := fltk.NewBox(fltk.FLAT_BOX, 0, 0, 0, 0, "0")
	dec := fltk.NewButton(0, 0, 0, 0, "Decrement")
	inc.SetCallback(func() {
		i++
		box.SetLabel(strconv.Itoa(i))
	})
	dec.SetCallback(func() {
		i--
		box.SetLabel(strconv.Itoa(i))
	})

	widgets_wrapper.Fixed(dec, 40)

	widgets_wrapper.End()
	return widgets_wrapper
}
