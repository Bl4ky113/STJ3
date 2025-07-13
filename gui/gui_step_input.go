package gui

import (
	"fmt"
	"math"

	"github.com/pwiecz/go-fltk"
)

const (
	STEP_DIAL_CLICK_LIMIT int = 10
	STEP_DIAL_STEP_VALUE int = 90
	STEP_DIAL_BORDER_SIZE int = 2.0
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
	fmt.Println("redraw")
	fltk.SetDrawColor(0x000000)
	fltk.DrawPie(
		main_wid.X() - STEP_DIAL_BORDER_SIZE,
		main_wid.Y() - STEP_DIAL_BORDER_SIZE,
		main_wid.W() + (2 * STEP_DIAL_BORDER_SIZE),
		main_wid.H() + (2 * STEP_DIAL_BORDER_SIZE),
		0.0,
		360.0,
	)
	fltk.SetDrawColor(fltk.ColorFromRgb(230, 230, 230))
	fltk.DrawPie(
		main_wid.X(),
		main_wid.Y(),
		main_wid.W(),
		main_wid.H(),
		0.0,
		360.0,
	)

	if (*value > 0 || *value == (-1) * STEP_DIAL_CLICK_LIMIT) {
		fltk.SetDrawColor(0xb0bf1a00)
		fltk.DrawPie(
			main_wid.X(),
			main_wid.Y(),
			main_wid.W(),
			main_wid.H(),
			(-1.0 * float64(*step)) + STEP_DIAL_STARTING_DEGREE,
			STEP_DIAL_STARTING_DEGREE,
		)
	}

	radious := float64(size) / 2.0
	offset := float64(size) / (2.0 * math.Sqrt2)

	fltk.SetDrawColor(0x000000)
	fltk.SetLineStyle(0, STEP_DIAL_BORDER_SIZE)
	fltk.DrawLine(
		int(main_wid.X()) + int(radious - offset) - 1,
		int(main_wid.Y()) + int(radious - offset) - 1,
		int(main_wid.X()) + int(radious + offset) + 1,
		int(main_wid.Y()) + int(radious + offset) + 1,
	)
	main_wid.DrawChildren()
	return
}

func step_dial_handle_left_click (main_wid *fltk.Group, step *int, value *int) {
	fmt.Println(*value)
	if (*value == STEP_DIAL_CLICK_LIMIT) {
		return
	} else if (*value < 0) {
		*value = 0
	}

	*value++

	if (*value % 2 != 0 || (*value > 4 && *value < STEP_DIAL_CLICK_LIMIT)) {
		return
	}

	*step = *value * STEP_DIAL_STEP_VALUE
	main_wid.Redraw()
	return
}

func step_dial_handle_right_click (main_wid *fltk.Group, step *int, value *int) {
	fmt.Println(*value)
	if (*value == (-1) * STEP_DIAL_CLICK_LIMIT) {
		return 
	}

	*value--

	if (*value % 2 != 0) {
		return
	} else if (*value < 0 && *value > (-1) * STEP_DIAL_CLICK_LIMIT) {
		return
	} else if (*value > 4 && *value < (STEP_DIAL_CLICK_LIMIT - 1)) {
		return
	}

	*step = *value * STEP_DIAL_STEP_VALUE
	main_wid.Redraw()
	return
}
