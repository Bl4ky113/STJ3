package gui_daily

import (
	"fmt"

	"github.com/pwiecz/go-fltk"

	stepdial "selfjournal/gui/stepdial_input"
)

const (
	DAILY_TAB_ROW_HEIGHT = 64
	DAILY_TAB_GRID_COLUMNS = 8
	DAILY_TAB_GRID_MARGIN = 0
	DAILY_TAB_GRID_GAP = 12
	DAILY_TAB_LABELS_WIDTH = 312
	DAILY_TAB_DIAL_INPUT_SIZE = 48
)

func Generate_daily_tab (parent *fltk.Group) *fltk.Scroll {
	fmt.Println("GENERATING DAILY TAB")
	wrapper := fltk.NewScroll(
		parent.X(),
		parent.Y(),
		parent.W(),
		parent.H(),
	)
	wrapper.SetType(fltk.SCROLL_VERTICAL)

	grid := fltk.NewGrid(
		wrapper.X(),
		wrapper.Y(),
		wrapper.W(),
		(DAILY_TAB_ROW_HEIGHT + DAILY_TAB_GRID_MARGIN + DAILY_TAB_GRID_GAP) * (32 + 1),
	)
	grid.SetLayout(
		32 + 1, // Extra row for the Days Labels
		DAILY_TAB_GRID_COLUMNS + 1, // Extra column for right spacing
		DAILY_TAB_GRID_MARGIN,
		DAILY_TAB_GRID_GAP,
	)

	add_days_labels(grid)

	for i := 1; i < 32 + 1; i++ {
		grid_label := fltk.NewBox(
			fltk.BORDER_BOX, 
			0, 0,
			DAILY_TAB_LABELS_WIDTH, DAILY_TAB_ROW_HEIGHT,
			"TEST Consectetur dolores fuga obcaecati quibusdam ducimus debitis id Cupiditate veritatis officiis nostrum veniam reprehenderit distinctio! Commodi perspiciatis saepe illo assumenda saepe ut. Aliquid id aliquam laboriosam corporis consectetur. Recusandae perspiciatis.",
		)

		grid_dial_input_1 := stepdial.NewStepInputDial(0,0,DAILY_TAB_DIAL_INPUT_SIZE)
		grid_dial_input_2 := stepdial.NewStepInputDial(0,0,DAILY_TAB_DIAL_INPUT_SIZE)
		grid_dial_input_3 := stepdial.NewStepInputDial(0,0,DAILY_TAB_DIAL_INPUT_SIZE)
		grid_dial_input_4 := stepdial.NewStepInputDial(0,0,DAILY_TAB_DIAL_INPUT_SIZE)
		grid_dial_input_5 := stepdial.NewStepInputDial(0,0,DAILY_TAB_DIAL_INPUT_SIZE)
		grid_dial_input_6 := stepdial.NewStepInputDial(0,0,DAILY_TAB_DIAL_INPUT_SIZE)
		grid_dial_input_7 := stepdial.NewStepInputDial(0,0,DAILY_TAB_DIAL_INPUT_SIZE)
		
		grid.SetWidget(grid_label, i, 0, fltk.GridHorizontal)
		grid.SetWidget(grid_dial_input_1.Widget, i, 1, fltk.GridCenter)
		grid.SetWidget(grid_dial_input_2.Widget, i, 2, fltk.GridCenter)
		grid.SetWidget(grid_dial_input_3.Widget, i, 3, fltk.GridCenter)
		grid.SetWidget(grid_dial_input_4.Widget, i, 4, fltk.GridCenter)
		grid.SetWidget(grid_dial_input_5.Widget, i, 5, fltk.GridCenter)
		grid.SetWidget(grid_dial_input_6.Widget, i, 6, fltk.GridCenter)
		grid.SetWidget(grid_dial_input_7.Widget, i, 7, fltk.GridCenter)
	}

	wrapper.Add(grid)
	grid.End()
	wrapper.End()

	return wrapper
}

func add_days_labels (grid *fltk.Grid) {
	const (
		DAILY_TAB_DAY_WIDTH = DAILY_TAB_DIAL_INPUT_SIZE
		DAILY_TAB_DAY_HEIGHT = 32
	)

	var days_labels [7]string = [7]string{"MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN"}

	for i := 1; i <= len(days_labels); i++ {
		day_label := fltk.NewBox(
			fltk.BORDER_BOX, 0, 0, 
			DAILY_TAB_DAY_WIDTH, DAILY_TAB_DAY_HEIGHT,
			days_labels[i - 1],
		)
		grid.SetWidget(day_label, 0, i, fltk.GridFill)
	}

	grid.SetWidget(fltk.NewBox(fltk.FLAT_BOX, 0, 0, 0, 0, ""), 0, 8, fltk.GridFill)
	return
}
