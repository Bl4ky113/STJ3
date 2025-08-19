package gui_daily

import (
	"fmt"

	"github.com/pwiecz/go-fltk"
)

const (
	DAILY_SUBMENU_GRID_MARGIN = 12
	DAILY_SUBMENU_GRID_GAP = 8
)	

func Generate_daily_submenu (parent *fltk.Group) *fltk.Group {
	fmt.Println("GENERATING DAILY SUBMENU")
	wrapper := fltk.NewGroup(
		parent.X(), parent.Y(),
		parent.W(), parent.H(),
	)

	grid := fltk.NewGrid(
		wrapper.X(), wrapper.Y(),
		wrapper.W(), wrapper.H(),
	)
	grid.SetLayout(3, 1, DAILY_SUBMENU_GRID_MARGIN, DAILY_SUBMENU_GRID_GAP)

	daily_submenu_add_week_btns(grid)
	daily_submenu_add_progress(grid)
	daily_submenu_add_daily_btns(grid)
	grid.End()

	wrapper.Add(grid)
	wrapper.End()
	return wrapper
}

func daily_submenu_add_week_btns (parent *fltk.Grid) {
	const (
		DAILY_SUBMENU_WEEK_BTN_WIDTH = 64
		DAILY_SUBMENU_WEEK_BTN_HEIGHT = 32
	)

	btns_flex := fltk.NewFlex(
		parent.X(), parent.Y(),
		parent.W() - (DAILY_SUBMENU_GRID_MARGIN * 2), DAILY_SUBMENU_WEEK_BTN_HEIGHT,
	)
	btns_flex.SetType(fltk.ROW)

	week_prev_btn := fltk.NewButton(
		btns_flex.X(), btns_flex.Y(),
		DAILY_SUBMENU_WEEK_BTN_WIDTH, DAILY_SUBMENU_WEEK_BTN_HEIGHT, 
		"PREV",
	)
	btns_flex.Fixed(week_prev_btn, DAILY_SUBMENU_WEEK_BTN_WIDTH)

	week_number_label := fltk.NewBox(
		fltk.BORDER_BOX, 0, 0, 
		DAILY_SUBMENU_WEEK_BTN_WIDTH, DAILY_SUBMENU_WEEK_BTN_HEIGHT,
		"WEEK #",
	)
	btns_flex.Add(week_number_label)

	week_next_btn := fltk.NewButton(
		(btns_flex.X() + btns_flex.W()) - DAILY_SUBMENU_WEEK_BTN_WIDTH, btns_flex.Y(),
		DAILY_SUBMENU_WEEK_BTN_WIDTH, DAILY_SUBMENU_WEEK_BTN_HEIGHT, 
		"NEXT",
	)
	btns_flex.Fixed(week_next_btn, DAILY_SUBMENU_WEEK_BTN_WIDTH)

	btns_flex.End()
	parent.SetWidget(btns_flex, 0, 0, fltk.GridTop)
	return
}

func daily_submenu_add_progress (parent *fltk.Grid) {
	return
}

func daily_submenu_add_daily_btns (parent *fltk.Grid) {
	const (
		DAILY_SUBMENU_END_BTN_WIDTH = 128
		DAILY_SUBMENU_END_BTN_HEIGHT = 32
	)

	end_day_btn := fltk.NewButton(
		parent.X(), parent.Y(),
		DAILY_SUBMENU_END_BTN_WIDTH, DAILY_SUBMENU_END_BTN_HEIGHT,
		"Finish Day",
	)

	parent.SetWidget(end_day_btn, 2, 0, fltk.GridCenter)
	return
}
