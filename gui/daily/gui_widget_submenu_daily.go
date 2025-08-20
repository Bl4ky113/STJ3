package gui_daily

import (
	"fmt"

	"github.com/pwiecz/go-fltk"
)

const (
	DAILY_SUBMENU_GRID_MARGIN = 12
	DAILY_SUBMENU_GRID_GAP = 0
	DAILY_SUBMENU_BTN_HEIGHT = 32
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

	submenu_add_week_btns(grid)
	submenu_add_progress(grid)
	submenu_add_daily_btns(grid)
	grid.End()

	wrapper.Add(grid)
	wrapper.End()
	return wrapper
}

func submenu_add_week_btns (parent *fltk.Grid) {
	const (
		DAILY_SUBMENU_WEEK_BTN_WIDTH = 64
		DAILY_SUBMENU_WEEK_BTN_HEIGHT = DAILY_SUBMENU_BTN_HEIGHT
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

func submenu_add_progress (parent *fltk.Grid) {
	const (
		PROGRESS_GRID_ROWS = 2
		PROGRESS_GRID_COLUMNS = 6
		PROGRESS_GRID_MARGIN = 0
		PROGRESS_GRID_GAP = 0
		
		DEFAULT_SVG_PATH = "gui/src/"
		TASK_NICE_SVG_PATH = DEFAULT_SVG_PATH + "mug.svg"
		TASK_SHOULD_SVG_PATH = DEFAULT_SVG_PATH + "gears.svg"
		TASK_LIKE_SVG_PATH = DEFAULT_SVG_PATH + "seedling.svg"
		TASK_HAVE_SVG_PATH = DEFAULT_SVG_PATH + "briefcase.svg"
		TASK_SPECIAL_GOOD_SVG_PATH = DEFAULT_SVG_PATH + "medal.svg"
		TASK_SPECIAL_BAD_SVG_PATH = DEFAULT_SVG_PATH + "bus.svg"
	)

	var icon_size int

	progress_grid := fltk.NewGrid(
		parent.X(), parent.Y(),
		parent.W() - (DAILY_SUBMENU_GRID_MARGIN * 2), parent.H() - ((DAILY_SUBMENU_GRID_GAP + DAILY_SUBMENU_GRID_MARGIN + DAILY_SUBMENU_BTN_HEIGHT) * 2),
	)
	progress_grid.SetShowGrid(true)
	progress_grid.SetLayout(
		PROGRESS_GRID_ROWS, 
		PROGRESS_GRID_COLUMNS,
		PROGRESS_GRID_MARGIN, PROGRESS_GRID_GAP,
	)

	icon_size = progress_grid.W() / PROGRESS_GRID_COLUMNS
	fmt.Printf("icon size %d, %d = %d\n", progress_grid.W(), PROGRESS_GRID_COLUMNS, icon_size)

	bar := fltk.NewBox(fltk.BORDER_FRAME, 0, 0, icon_size, icon_size, "")
	foo, _ := fltk.NewSvgImageLoad(TASK_LIKE_SVG_PATH)
	foo.Scale(icon_size, icon_size, true, false)
	bar.SetImage(foo)

	progress_grid.SetWidget(bar, 0, 0, fltk.GridCenter)

	progress_grid.End()
	parent.SetWidget(progress_grid, 1, 0, fltk.GridFill)
	return
}

func submenu_add_daily_btns (parent *fltk.Grid) {
	const (
		DAILY_SUBMENU_END_BTN_WIDTH = 128
		DAILY_SUBMENU_END_BTN_HEIGHT = DAILY_SUBMENU_BTN_HEIGHT
	)

	end_day_btn := fltk.NewButton(
		parent.X(), parent.Y(),
		DAILY_SUBMENU_END_BTN_WIDTH, DAILY_SUBMENU_END_BTN_HEIGHT,
		"Finish Day",
	)

	parent.SetWidget(end_day_btn, 2, 0, fltk.GridBottom)
	return
}
