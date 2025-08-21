package gui_daily

import (
	"fmt"

	"github.com/pwiecz/go-fltk"

	g "selfjournal/globals"
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
		PROGRESS_LABEL_HEIGHT = 16
	)

	var icon_size int
	var icon_id_order [g.NUM_TASKS]int = [g.NUM_TASKS]int{g.HAVE_TASK_ID, g.NICE_TASK_ID, g.SHOULD_TASK_ID, g.LIKE_TASK_ID, g.SPECIAL_GOOD_TASK_ID, g.SPECIAL_BAD_TASK_ID}

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

	for i, icon_id := range icon_id_order {
		var icon_svg_path string

		switch icon_id {
		case g.HAVE_TASK_ID:
			icon_svg_path = g.TASK_HAVE_SVG_PATH
		case g.NICE_TASK_ID:
			icon_svg_path = g.TASK_NICE_SVG_PATH
		case g.SHOULD_TASK_ID:
			icon_svg_path = g.TASK_SHOULD_SVG_PATH
		case g.LIKE_TASK_ID:
			icon_svg_path = g.TASK_LIKE_SVG_PATH
		case g.SPECIAL_GOOD_TASK_ID:
			icon_svg_path = g.TASK_SPECIAL_GOOD_SVG_PATH
		case g.SPECIAL_BAD_TASK_ID:
			icon_svg_path = g.TASK_SPECIAL_BAD_SVG_PATH
		}

		icon_box := fltk.NewBox(
			fltk.FLAT_BOX, progress_grid.X(), progress_grid.Y(),
			icon_size, icon_size, "",
		)

		icon_svg, err := fltk.NewSvgImageLoad(icon_svg_path)
		if err != nil {
			// ERROR LOADING SVG FROM icon_svg_path
			continue
		}

		icon_svg.Scale(icon_size, icon_size, true, true)
		icon_box.SetImage(icon_svg)

		progress_box := fltk.NewBox(
			fltk.FLAT_BOX, progress_grid.X(), progress_grid.Y(),
			icon_size, PROGRESS_LABEL_HEIGHT, 
			"123",
		)

		progress_grid.SetWidget(icon_box, 0, i, fltk.GridBottom)
		progress_grid.SetWidget(progress_box, 1, i, fltk.GridTop)
	}

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
