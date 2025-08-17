package gui

// PRIMES USED:
// 	- TAB_ID_PRIME: 2
//  - WIDGET_ID_PRIME: 5

import (
	"github.com/pwiecz/go-fltk"
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

const (
	TAB_ID_PRIME = 2
	WIDGET_ID_PRIME = 7
)

const (
	NUM_TABS = 3
	NUM_WIDGETS = 6
)

const (
	DAILY_TAB_ID int = TAB_ID_PRIME << iota
	THEME_TAB_ID
	THOUGHTS_TAB_ID 
)

const (
	DAILY_SUBMENU_ID int = WIDGET_ID_PRIME << iota
	THEME_SUBMENU_ID
	THOUGHTS_SUBMENU_ID
	DAILY_WIDGET_ID
	THEME_WIDGET_ID
	THOUGHTS_WIDGET_ID
)

const (
	COLOR_INPUT_MAIN fltk.Color = 0x50B7E000
	COLOR_INPUT_DARK fltk.Color = 0x4795B900
	COLOR_INPUT_DARKER fltk.Color = 0x3E729100
	COLOR_INPUT_SPECIAL_POSITIVE fltk.Color = 0x2B2D4200
	COLOR_INPUT_SPECIAL_NEGATIVE fltk.Color = 0xEF233C00
	COLOR_INPUT_BORDER fltk.Color = 0x000000
	COLOR_INPUT_BACKGROUND fltk.Color = 0xefefef00
)

var screen_width, screen_height int
var window_width, window_height int 

var active_tab_id int = DAILY_TAB_ID
var active_tab_ptr *fltk.Group = nil
var tab_ptrs_map map[int]*fltk.Group = nil
var tab_widgets_ids map[int][3]int = nil

var active_widgets_id [3]int = [3]int{DAILY_SUBMENU_ID, THEME_WIDGET_ID, THOUGHTS_WIDGET_ID}
var active_widgets_ptrs [3]*fltk.Group = [3]*fltk.Group{nil, nil, nil}
var widgets_ptr_map map[int]*fltk.Group = nil
var widgets_grid_ptr *fltk.Grid = nil
