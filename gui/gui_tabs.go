package gui

import (
	"fmt"

	"github.com/pwiecz/go-fltk"

	g "selfjournal/globals"

	gui_daily "selfjournal/gui/daily"
)

func handle_tab_create (tab_id int, parent *fltk.Group) {
	switch tab_id {
	case g.DAILY_TAB_ID:
		tab_widgets_ids[tab_id] = [3]int{g.DAILY_SUBMENU_ID, g.THEME_WIDGET_ID, g.THOUGHTS_WIDGET_ID}

		parent.Add(gui_daily.Generate_daily_tab(parent))
		break
	case g.THEME_TAB_ID:
		tab_widgets_ids[tab_id] = [3]int{g.THEME_SUBMENU_ID, g.DAILY_WIDGET_ID, g.THOUGHTS_WIDGET_ID}

		value_box := fltk.NewBox(fltk.NO_BOX, 0, 0, 40, 40, "321")
		value_box.SetLabelSize(26)
		parent.Add(value_box)
		parent.SetColor(fltk.RED)
		break
	case g.THOUGHTS_TAB_ID:
		tab_widgets_ids[tab_id] = [3]int{g.THOUGHTS_SUBMENU_ID, g.THEME_WIDGET_ID, g.DAILY_WIDGET_ID}

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
	fmt.Printf("HIDDING TAB %d\n", tab_id)
	return
}
