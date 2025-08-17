package gui

import (
	"fmt"
	"strconv"

	"github.com/pwiecz/go-fltk"
)

func handle_widget_create (widget_id int, parent *fltk.Group) {
	foo := fltk.NewBox(fltk.BORDER_BOX, parent.X(), parent.Y(), 100, 100, strconv.Itoa(widget_id))

	parent.Add(foo)
	return
}

func handle_widgets_show (tab_id int) {
	if (tab_widgets_ids == nil) || (widgets_ptr_map == nil) { return }

	if active_tab_ptr != nil {
		handle_widgets_hide(active_tab_id)
	}

	for i, widget_id := range tab_widgets_ids[tab_id] {
		widget_ptr := widgets_ptr_map[widget_id]
		widget_ptr.Show()

		widgets_grid_ptr.SetWidget(widget_ptr, i, 0, fltk.GridFill)

		active_widgets_id[i] = widget_id
		active_widgets_ptrs[i] = widget_ptr
		fmt.Printf("SHOWING WIDGET %d\n", widget_id)
	}

	return
}

func handle_widgets_hide (tab_id int) {

}
