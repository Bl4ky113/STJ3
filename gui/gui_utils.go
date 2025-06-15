package gui

import (
	"errors"

	// #cgo LDFLAGS: -lX11
	// int get_width ();
	// int get_height ();
	// int calc_screen_percentage (double percentage, int screen_measure);
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

func calc_screen_percentage_width [N Number](percentage N) int {
	return _call_c_calc_value_percentage(percentage, screen_width)
}

func calc_screen_percentage_height [N Number](percentage N) int {
	return _call_c_calc_value_percentage(percentage, screen_height)
}

func calc_window_percentage_width [N Number](percentage N) int {
	return _call_c_calc_value_percentage(percentage, window_width)
}

func calc_window_percentage_height [N Number](percentage N) int {
	return _call_c_calc_value_percentage(percentage, window_height)
}

func _call_c_calc_value_percentage [N Number](percentage N, value int) int {
	return int(C.calc_value_percentage(C.double(percentage), C.int(value)))
}
