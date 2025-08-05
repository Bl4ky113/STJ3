package gui

import (
	"fmt"
	"math"

	"github.com/pwiecz/go-fltk"
)

const (
	STEP_DIAL_CLICK_LIMIT int = 10
	STEP_DIAL_STEP_VALUE int = 90
	STEP_DIAL_BORDER_SIZE int = 2
	STEP_DIAL_STARTING_DEGREE float64 = -45.0
)

type StepInputDial struct {
	main_wid  *fltk.Group
	step *int
	value *int
}

func NewStepInputDial (x, y, size int) *StepInputDial {
	step := 0
	value := 0

	main_wid := fltk.NewGroup(x, y, size, size)
	main_wid.End()

	main_wid.SetDrawHandler(func(func()) {
		step_dial_handle_redraw(main_wid, &step, &value, size)
	})

	main_wid.SetEventHandler(func(e fltk.Event) bool {
		if (e != fltk.PUSH) { 
			return false
		}

		switch (fltk.EventButton()) {
		case fltk.LeftMouse:
			step_dial_handle_left_click(main_wid, &step, &value)
			break
		case fltk.RightMouse:
			step_dial_handle_right_click(main_wid, &step, &value)
			break
		}

		return true
	})

	return &StepInputDial{
		main_wid,
		&step,
		&value,
	}
}

func (d *StepInputDial) SetValue(val int) {
	*d.value = val

	if (val > STEP_DIAL_CLICK_LIMIT) {
		*d.value = STEP_DIAL_CLICK_LIMIT
	} else if (val < (-1) * STEP_DIAL_CLICK_LIMIT) {
		*d.value = (-1) * STEP_DIAL_CLICK_LIMIT
	}

	d.main_wid.Redraw()
}

func (d *StepInputDial) Value () int {
	return *d.value / 2
}

func (d *StepInputDial) CurrentStep() int {
	return *d.step
}

func step_dial_handle_redraw (main_wid *fltk.Group, step *int, value *int, size int) {
	var color_input_current_status fltk.Color

	fmt.Println("redraw")
	fltk.SetDrawColor(COLOR_INPUT_BACKGROUND)
	fltk.DrawPie(
		main_wid.X(),
		main_wid.Y(),
		main_wid.W(),
		main_wid.H(),
		0.0,
		360.0,
	)
	
	switch (*value) {
		case 10:
			color_input_current_status = COLOR_INPUT_SPECIAL_POSITIVE
			break
		case -10:
			color_input_current_status = COLOR_INPUT_SPECIAL_NEGATIVE
			break
		default:
			color_input_current_status = COLOR_INPUT_MAIN
			break
	}

	if (*value > 0 || *value == (-1) * STEP_DIAL_CLICK_LIMIT) {
		fltk.SetDrawColor(color_input_current_status)
		fltk.DrawPie(
			main_wid.X() - 1,
			main_wid.Y() - 1,
			main_wid.W() + 1,
			main_wid.H() + 1,
			(-1.0 * float64(*step)) + STEP_DIAL_STARTING_DEGREE,
			STEP_DIAL_STARTING_DEGREE,
		)
	}

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
	fltk.SetLineStyle(fltk.SOLID, STEP_DIAL_BORDER_SIZE)
	fltk.DrawArc(
		main_wid.X(), main_wid.Y(),
		main_wid.W(), main_wid.H(),
		0.0, 360.0,
	)

	main_wid.DrawChildren()
	return
}

func step_dial_handle_left_click (main_wid *fltk.Group, step *int, value *int) {
	if (*value == STEP_DIAL_CLICK_LIMIT) {
		return
	} else if (*value < 0) {
		*value = 0
	}

	*value++

	if (
		*value % 2 != 0 || 
		(*value > 4 && *value < STEP_DIAL_CLICK_LIMIT)) {
		return
	}

	*step = *value * STEP_DIAL_STEP_VALUE
	main_wid.Redraw()
	return
}

func step_dial_handle_right_click (main_wid *fltk.Group, step *int, value *int) {
	if (*value == (-1) * STEP_DIAL_CLICK_LIMIT) {
		return 
	}

	*value--

	if (
		*value % 2 != 0 ||
		*value < 0 && *value > (-1) * STEP_DIAL_CLICK_LIMIT ||
		*value > 4 && *value < (STEP_DIAL_CLICK_LIMIT - 2)) {
		return
	} 

	*step = *value * STEP_DIAL_STEP_VALUE
	main_wid.Redraw()
	return
}
