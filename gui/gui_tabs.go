package gui

import (
	"fmt"

	"github.com/pwiecz/go-fltk"
)

func handle_tab_create (tab_id int, parent *fltk.Group) {
	switch tab_id {
	case DAILY_TAB_ID:
		tab_widgets_ids[tab_id] = [3]int{DAILY_SUBMENU_ID, THEME_WIDGET_ID, THOUGHTS_WIDGET_ID}

		daily_tab_group := generate_daily_tab(parent)
		parent.Add(daily_tab_group)
		break
	case THEME_TAB_ID:
		tab_widgets_ids[tab_id] = [3]int{THEME_SUBMENU_ID, DAILY_WIDGET_ID, THOUGHTS_WIDGET_ID}

		value_box := fltk.NewBox(fltk.NO_BOX, 0, 0, 40, 40, "321")
		value_box.SetLabelSize(26)
		parent.Add(value_box)
		parent.SetColor(fltk.RED)
		break
	case THOUGHTS_TAB_ID:
		tab_widgets_ids[tab_id] = [3]int{THOUGHTS_SUBMENU_ID, THEME_WIDGET_ID, DAILY_WIDGET_ID}

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

	fmt.Printf("SHOWING TAB %d\n", tab_id)
	active_tab_id = tab_id
	active_tab_ptr = tab_ptr

	return
}

func handle_tab_hide (tab_id int) {
	if tab_ptrs_map == nil { return }

	tab_ptr, tab_exists := tab_ptrs_map[tab_id]

	if !tab_exists {
		return
	}
	
	tab_ptr.Hide()
	fmt.Printf("HIDDING %d\n", tab_id)
	return
}
