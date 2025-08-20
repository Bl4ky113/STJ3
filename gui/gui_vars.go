package gui

import (
	"github.com/pwiecz/go-fltk"

	g "selfjournal/globals"
)

type Number interface {
	~int | ~float64 | ~float32
}

const (
	WINDOW_WIDTH_PERCENTAGE int = 75
	WINDOW_HEIGHT_PERCENTAGE int = 75
	WINDOW_DECORATION_HEIGTH int = 30 // 10 padding top and bottom, 10 font size
)

const (
	MENU_WIDTH_PERCENTAGE int = 5
	MAIN_WIDTH_PERCENTAGE int = 70 // 100 - 5 - 25
	WIDGETS_WIDTH_PERCENTAGE int = 25
	WIDGETS_HEIGHT_PERCENTAGE int = 30
)

const TITLE_STR string = "Self Thoughts Journal 3"

var screen_width, screen_height int
var window_width, window_height int 

var active_tab_id int = g.DAILY_TAB_ID
var active_tab_ptr *fltk.Group = nil
var tab_ptrs_map map[int]*fltk.Group = nil
var tab_widgets_ids map[int][3]int = nil

var active_widgets_id [3]int = [3]int{g.DAILY_SUBMENU_ID, g.THEME_WIDGET_ID, g.THOUGHTS_WIDGET_ID}
var active_widgets_ptrs [3]*fltk.Group = [3]*fltk.Group{nil, nil, nil}
var widgets_ptr_map map[int]*fltk.Group = nil
var widgets_grid_ptr *fltk.Grid = nil
