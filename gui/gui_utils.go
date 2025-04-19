package gui

import (
	"errors"

    // #cgo LDFLAGS: -lX11
	// int get_width ();
	// int get_height ();
	// int calc_screen_percentage (int percentage, int screen_measure);
    // #include "../c_aux_code/screen_size.c"
    "C"
)

func get_screen_size () error {
	width := int(C.get_width())
	if width < 0 {
		return errors.New("Error on getting Screen Width") 
	}

	height := int(C.get_height())
	if height < 0 {
		return errors.New("Error on getting Screen Height")
	}

	screen_width = width
	screen_height = height
	return nil
}

func calc_screen_percentage (percentage int, use_height bool) int {
	if use_height {
		return _call_c_calc_value_percentage(percentage, screen_height)
	}

	return _call_c_calc_value_percentage(percentage, screen_width)
}

func calc_window_percentage(percentage int, use_height bool) int {
	if use_height {
		return _call_c_calc_value_percentage(percentage, window_height)
	}

	return _call_c_calc_value_percentage(percentage, window_width)
}

func _call_c_calc_value_percentage (percentage int, value int) int {
	return int(C.calc_value_percentage(C.int(percentage), C.int(value)))
}
