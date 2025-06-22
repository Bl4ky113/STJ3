package gui

// PRIMES USED:
// 	- TAB_ID_PRIME: 2

import (
	"github.com/pwiecz/go-fltk"
)

type Number interface {
	~int | ~float64 | ~float32
}

const WINDOW_WIDTH_PERCENTAGE int = 75
const WINDOW_HEIGHT_PERCENTAGE int = 75
const WINDOW_DECORATION_HEIGTH int = 30 // 10 padding top and bottom, 10 font size

const MENU_WIDTH_PERCENTAGE int = 5
const MAIN_WIDTH_PERCENTAGE int = 70 // 100 - 5 - 25
const WIDGETS_WIDTH_PERCENTAGE int = 25
const WIDGETS_HEIGHT_PERCENTAGE int = 32

const TITLE_STR string = "Self Thoughts Journal 3"

const (
	NUM_TABS = 3
	TAB_ID_PRIME = 2
)

const (
	DAILY_TAB_ID int = TAB_ID_PRIME << iota
	THEME_TAB_ID
	THOUGHTS_TAB_ID 
)

var screen_width, screen_height int
var window_width, window_height int 

var active_tab_id int = DAILY_TAB_ID
var active_tab_ptr *fltk.Group = nil
var tab_ptrs_map map[int]*fltk.Group = nil

var i int = 0
