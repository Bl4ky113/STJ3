package gui_stepdial_input

import (
	"math"

	"github.com/pwiecz/go-fltk"
)

const (
	COLOR_INPUT_MAIN fltk.Color = 0x50B7E000
	COLOR_INPUT_DARK fltk.Color = 0x4795B900
	COLOR_INPUT_DARKER fltk.Color = 0x3E729100
	COLOR_INPUT_SPECIAL_POSITIVE fltk.Color = 0x2B2D4200
	COLOR_INPUT_SPECIAL_NEGATIVE fltk.Color = 0xEF233C00
	COLOR_INPUT_BORDER fltk.Color = 0x000000
	COLOR_INPUT_BACKGROUND fltk.Color = 0xefefef00

	STEP_DIAL_CLICK_LIMIT int = 10
	STEP_DIAL_STEP_VALUE int = 180
	STEP_DIAL_ANIMATION_STEP_VALUE int = 6
	STEP_DIAL_ANIMATION_STEP_TIMEOUT float64 = 0.01
	STEP_DIAL_BORDER_SIZE int = 2
	STEP_DIAL_STARTING_DEGREE float64 = -45.0
)

type StepInputDial struct {
	Widget  *fltk.Group
	degree *int
	value *int
	on_animation *bool
	currently_increasing *int
}

func NewStepInputDial (x, y, size int) *StepInputDial {
	degree := 0
	value := 0
	on_animation := false
	currently_increasing := 0

	main_wid := fltk.NewGroup(x, y, size, size)
	main_wid.End()

	main_wid.SetDrawHandler(func(func()) {
		step_dial_handle_redraw(main_wid, &degree, &value, &on_animation, &currently_increasing, size)
	})

	main_wid.SetEventHandler(func(e fltk.Event) bool {
		if (e != fltk.PUSH) { 
			return false
		}

		switch (fltk.EventButton()) {
		case fltk.LeftMouse:
			step_dial_handle_left_click(main_wid, &degree, &value, &on_animation, &currently_increasing)
			break
		case fltk.RightMouse:
			step_dial_handle_right_click(main_wid, &degree, &value, &on_animation, &currently_increasing)
			break
		}

		return true
	})

	return &StepInputDial{
		main_wid,
		&degree,
		&value,
		&on_animation,
		&currently_increasing,
	}
}

func (d *StepInputDial) Value () int {
	return *d.value / 2
}

func (d *StepInputDial) SetValue(val int) {
	*d.value = val

	if (val > STEP_DIAL_CLICK_LIMIT) {
		*d.value = STEP_DIAL_CLICK_LIMIT
	} else if (val < (-1) * STEP_DIAL_CLICK_LIMIT) {
		*d.value = (-1) * STEP_DIAL_CLICK_LIMIT
	}

	d.Widget.Redraw()
}

func step_dial_handle_redraw (main_wid *fltk.Group, degree *int, value *int, on_animation *bool, currently_increasing *int, size int) {
	fltk.SetDrawColor(COLOR_INPUT_BACKGROUND)
	fltk.DrawPie(
		main_wid.X(),
		main_wid.Y(),
		main_wid.W(),
		main_wid.H(),
		0.0,
		360.0,
	)
	
	step_dial_draw_current_status_pie(main_wid, degree, value, on_animation, currently_increasing)

	radious := float64(size) / 2.0
	offset := float64(size) / (2.0 * math.Sqrt2)

	fltk.SetDrawColor(COLOR_INPUT_BORDER)
	fltk.SetLineStyle(fltk.SOLID, STEP_DIAL_BORDER_SIZE)
	fltk.DrawLine(
		int(main_wid.X()) + int(radious - offset),
		int(main_wid.Y()) + int(radious - offset),
		int(main_wid.X()) + int(radious + offset),
		int(main_wid.Y()) + int(radious + offset),
	)
	fltk.DrawArc(
		main_wid.X(), main_wid.Y(),
		main_wid.W(), main_wid.H(),
		0.0, 360.0,
	)

	main_wid.DrawChildren()
	return
}

func step_dial_draw_current_status_pie (main_wid *fltk.Group, degree *int, value *int, on_animation *bool, currently_increasing *int) {
	var color_input_current_status fltk.Color
	var current_value_max_degree, current_value_min_degree int

	switch (*value) {
	case 2:
		current_value_min_degree = 0
		current_value_max_degree = 180

		if (*currently_increasing == -1) {
			current_value_min_degree = 180
			current_value_max_degree = *degree
		}

		color_input_current_status = COLOR_INPUT_MAIN
		break
	case 4:
		current_value_min_degree = 180
		current_value_max_degree = 360

		color_input_current_status = COLOR_INPUT_MAIN

		if (*currently_increasing == -1) {
			current_value_min_degree = -360
			current_value_max_degree = 0
			if (*degree > 0) {
				color_input_current_status = COLOR_INPUT_SPECIAL_POSITIVE
			} else {
				color_input_current_status = COLOR_INPUT_MAIN
			}
		}
		break
	case 10:
		if (*degree == 0 && !(*on_animation)) {
			*degree = -360
		}

		current_value_min_degree = 0
		current_value_max_degree = 360

		if (*degree > 0) {
			color_input_current_status = COLOR_INPUT_SPECIAL_POSITIVE
		} else {
			color_input_current_status = COLOR_INPUT_MAIN
		}
		break
	case -10:
		current_value_min_degree = -360
		current_value_max_degree = 0
		color_input_current_status = COLOR_INPUT_SPECIAL_NEGATIVE
		break
	default:
		current_value_min_degree = 0
		current_value_max_degree = 0
		color_input_current_status = COLOR_INPUT_BACKGROUND
		
		if (*currently_increasing == -1) { // from 2 -> 0
			current_value_min_degree = 0
			current_value_max_degree = *degree
			color_input_current_status = COLOR_INPUT_MAIN
		} else { // from -10 -> 0
			current_value_min_degree = *degree
			current_value_max_degree = 0
			color_input_current_status = COLOR_INPUT_SPECIAL_NEGATIVE
		}
		break
	}

	fltk.SetDrawColor(color_input_current_status)
	fltk.DrawPie(
		main_wid.X(),
		main_wid.Y(),
		main_wid.W(),
		main_wid.H(),
		(-1.0 * float64(*degree)) + STEP_DIAL_STARTING_DEGREE,
		STEP_DIAL_STARTING_DEGREE,
	)

	if (*currently_increasing == 1) {
		if (*degree < current_value_max_degree) {
			*on_animation = true
			*degree += STEP_DIAL_ANIMATION_STEP_VALUE

			fltk.AddTimeout(
				STEP_DIAL_ANIMATION_STEP_TIMEOUT,
				func () {
					main_wid.Redraw()
				},
			)

		} else {
			*on_animation = false
			*currently_increasing = 0

			if (*value == 10) {
				*degree = 360 // Fix degree from double cicle animation
			}
		}
	} else if (*currently_increasing == -1){
		if (*degree > current_value_min_degree) {
			*on_animation = true
			*degree -= STEP_DIAL_ANIMATION_STEP_VALUE

			fltk.AddTimeout(
				STEP_DIAL_ANIMATION_STEP_TIMEOUT,
				func () {
					main_wid.Redraw()
				},
			)

		} else {
			*on_animation = false
			*currently_increasing = 0

			if (*value == 4) {
				*degree = 360 // Fix degree from double cicle animation
			}
		}
	}


	return
}

func step_dial_handle_left_click (main_wid *fltk.Group, degree *int, value *int, on_animation *bool, currently_increasing *int) {
	if (*value == STEP_DIAL_CLICK_LIMIT || *on_animation == true) {
		return
	} 

	if (*value < 0) {
		*value = -1
	}

	*value++

	if (
		(*value % 2 != 0) || 
		(*value > 4 && *value < STEP_DIAL_CLICK_LIMIT)) {
		return
	}

	if (*value == STEP_DIAL_CLICK_LIMIT) {
		*degree = 0 // Guarantee full animation
	}

	//*step = *value * STEP_DIAL_STEP_VALUE
	//*degree += STEP_DIAL_ANIMATION_STEP_VALUE
	*currently_increasing = 1
	main_wid.Redraw()
	return
}

func step_dial_handle_right_click (main_wid *fltk.Group, degree *int, value *int, on_animation *bool, currently_increasing *int) {
	if (*value == (-1) * STEP_DIAL_CLICK_LIMIT || *on_animation) {
		return
	} 

	*value--

	if (
		(*value % 2 != 0) ||
		(*value < 0 && *value > (-1) * STEP_DIAL_CLICK_LIMIT) ||
		(*value > 4 && *value < STEP_DIAL_CLICK_LIMIT)) {
		return
	}

	if (*value == (-1) * STEP_DIAL_CLICK_LIMIT) {
		*degree = 0 // Guarantee full animation
	}

	//*step = *value * STEP_DIAL_STEP_VALUE
	//*degree -= STEP_DIAL_ANIMATION_STEP_VALUE
	*currently_increasing = -1
	main_wid.Redraw()
	return
}
