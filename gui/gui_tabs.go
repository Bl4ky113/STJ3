package gui

import (
	"fmt"

	"github.com/pwiecz/go-fltk"
)

func handle_tab_create (tab_id int, parent *fltk.Group) {
	switch tab_id {
	case DAILY_TAB_ID:
		new_dial := NewStepInputDial(0, 0, 100)
		new_dial.SetValue(180)
		parent.Add(new_dial.main_wid)
		new_dial2 := NewStepInputDial(300, 0, 200)
		parent.Add(new_dial2.main_wid)
		new_dial3 := NewStepInputDial(0, 300, 50)
		new_dial3.SetValue(360)
		parent.Add(new_dial3.main_wid)
		parent.SetColor(fltk.BLUE)
		break
	case THEME_TAB_ID:
		value_box := fltk.NewBox(fltk.NO_BOX, 0, 0, 40, 40, "321")
		value_box.SetLabelSize(26)
		parent.Add(value_box)
		parent.SetColor(fltk.RED)
		break
	case THOUGHTS_TAB_ID:
		parent.SetColor(fltk.GREEN)
		break
	default:
		parent.SetColor(fltk.BLACK)
		break
	}
}

func handle_tab_show (tab_id int) {
	if tab_ptrs_map == nil { return }

	if active_tab_ptr != nil {
		handle_tab_hide(active_tab_id)
	}

	tab_ptr := tab_ptrs_map[tab_id]
	tab_ptr.Show()

	fmt.Printf("SHOWING %d\n", tab_id)
	active_tab_id = tab_id
	active_tab_ptr = tab_ptr
}

func handle_tab_hide (tab_id int) {
	if tab_ptrs_map == nil { return }

	tab_ptr, tab_exists := tab_ptrs_map[tab_id]

	if !tab_exists {
		return
	}
	
	tab_ptr.Hide()
	fmt.Printf("HIDDING %d\n", tab_id)
}
