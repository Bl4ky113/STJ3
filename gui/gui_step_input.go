package gui

import (
	"math"

	"github.com/pwiecz/go-fltk"
)

const (
	BORDER_SIZE int = 2.0
	STARTING_DEGREE float64 = -45.0
)

type StepInputDial struct {
	main_wid  *fltk.Group
	step *int
}

func NewStepInputDial (x, y, size int) *StepInputDial {
	value := 0
	main_wid := fltk.NewGroup(x, y, size, size)
	main_wid.End()

	main_wid.SetDrawHandler(func(func()) {
		fltk.SetDrawColor(0x000000)
		fltk.DrawPie(
			main_wid.X() - BORDER_SIZE,
			main_wid.Y() - BORDER_SIZE,
			main_wid.W() + (2 * BORDER_SIZE),
			main_wid.H() + (2 * BORDER_SIZE),
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
		fltk.SetDrawColor(0xb0bf1a00)
		fltk.DrawPie(
			main_wid.X(),
			main_wid.Y(),
			main_wid.W(),
			main_wid.H(),
			(-1.0 * float64(value)) + STARTING_DEGREE,
			STARTING_DEGREE,
		)
		fltk.SetDrawColor(0x000000)
		fltk.SetLineStyle(0, BORDER_SIZE)

		radious := float64(size) / 2.0
		offset := float64(size) / (2.0 * math.Sqrt2)
		fltk.DrawLine(
			int(main_wid.X()) + int(radious - offset) - 1,
			int(main_wid.Y()) + int(radious - offset) - 1,
			int(main_wid.X()) + int(radious + offset) + 1,
			int(main_wid.Y()) + int(radious + offset) + 1,
		)
		main_wid.DrawChildren()
	})
	return &StepInputDial{
		main_wid,
		&value,
	}
}

func (d *StepInputDial) SetValue(val int) {
	*d.step = val
	d.main_wid.Redraw()
}

func (d *StepInputDial) Value() int {
	return *d.step
}

