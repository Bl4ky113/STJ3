package gui

import (
	"fmt"
	"strconv"

	"github.com/pwiecz/go-fltk"

	g "selfjournal/globals"

	gui_daily "selfjournal/gui/daily"
	
)

func handle_widget_create (widget_id int, parent *fltk.Group) {
	switch widget_id {
	case g.DAILY_SUBMENU_ID:
		parent.Add(gui_daily.Generate_daily_submenu(parent))
	default:
		parent.Add(fltk.NewBox(fltk.BORDER_BOX, parent.X(), parent.Y(), 150, 150, strconv.Itoa(widget_id)))
	}

	return
}

func handle_widgets_show (tab_id int) {
	if (tab_widgets_ids == nil) || (widgets_ptr_map == nil) { return }

	if active_tab_ptr != nil {
		handle_widgets_hide(active_tab_id)
	}

	for i, widget_id := range tab_widgets_ids[tab_id] {
		widget_ptr, widget_exists := widgets_ptr_map[widget_id]

		if !widget_exists {
			// Error, one widget needed to load doesn't exists
			return
		}

		widget_ptr.Show()

		widgets_grid_ptr.SetWidget(widget_ptr, i, 0, fltk.GridFill)

		active_widgets_id[i] = widget_id
		active_widgets_ptrs[i] = widget_ptr
		fmt.Printf("SHOWING WIDGET %d\n", widget_id)
	}

	return
}

func handle_widgets_hide (tab_id int) {
	if (tab_widgets_ids == nil) || (widgets_ptr_map == nil) { return }

	_, tab_exists := tab_ptrs_map[tab_id]

	if !tab_exists {
		return
	}

	widgets_grid_ptr.ClearLayout() // Thanks Me
	widgets_grid_ptr.SetLayout(3, 1, 0, 0)

	for _, widget_id := range tab_widgets_ids[tab_id] {
		widget_ptr, widget_exists := widgets_ptr_map[widget_id]

		if !widget_exists {
			// Error, one widget needed to load doesn't exists
			return
		}

		widget_ptr.Hide()
		fmt.Printf("HIDDING WIDGET %d\n", widget_id)
	}
}
