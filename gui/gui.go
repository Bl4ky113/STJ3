package gui

import (
	"github.com/pwiecz/go-fltk"

	g "selfjournal/globals"
)

func Run_gui () error {
	err := get_screen_size()
	if err != nil {
		return err
	}

	create_window(screen_width, screen_height)

	fltk.Run()
	return nil
}

func create_window (screen_width, screen_height int) error {
	window_width = calc_screen_percentage_width(WINDOW_WIDTH_PERCENTAGE)
	window_height = calc_screen_percentage_height(WINDOW_HEIGHT_PERCENTAGE)

	win := fltk.NewWindowWithPosition(
		calc_screen_percentage_width((float32) (100 - WINDOW_WIDTH_PERCENTAGE) / 2.0),
		calc_screen_percentage_height((float32) (100 - WINDOW_HEIGHT_PERCENTAGE) / 2.0) + WINDOW_DECORATION_HEIGTH,
		window_width,
		window_height,
		TITLE_STR,
	)

	window_flex_wrapper := fltk.NewFlex(0, 0, window_width, window_height)
	window_flex_wrapper.SetType(fltk.ROW)
	//window_flex_wrapper.SetGap(10) // Save it for later, for the makeup of the app

	menu_wrapper := create_menu_section()
	main_wrapper := create_main_section()
	widgets_wrapper := create_widget_section()

	window_flex_wrapper.End()
	window_flex_wrapper.Fixed(menu_wrapper, calc_window_percentage_width(MENU_WIDTH_PERCENTAGE))
	window_flex_wrapper.Fixed(main_wrapper, calc_window_percentage_width(MAIN_WIDTH_PERCENTAGE))
	window_flex_wrapper.Fixed(widgets_wrapper, calc_window_percentage_width(WIDGETS_WIDTH_PERCENTAGE))

	handle_tab_show(g.DAILY_TAB_ID)
	handle_widgets_show(g.DAILY_TAB_ID)

	win.End()
	win.Add(window_flex_wrapper)
	win.Show()
	return nil
}

func create_menu_section () fltk.Widget {
	const ( 
		MENU_BTN_SIZE = 32
		MENU_BTN_GAP = 5
	)

	menu_wrapper := fltk.NewFlex(
		0, 0, 
		calc_window_percentage_width(MENU_WIDTH_PERCENTAGE), window_height,
	)
	menu_wrapper.SetGap(MENU_BTN_GAP)

	daily_btn := fltk.NewButton(0, 0, MENU_BTN_SIZE, MENU_BTN_SIZE, "Daily")
	menu_wrapper.Fixed(daily_btn, MENU_BTN_SIZE)

	theme_btn := fltk.NewButton(0, 0, MENU_BTN_SIZE, MENU_BTN_SIZE, "Theme")
	menu_wrapper.Fixed(theme_btn, MENU_BTN_SIZE)

	thoughts_btn := fltk.NewButton(0, 0, MENU_BTN_SIZE, MENU_BTN_SIZE, "Thoughts")
	menu_wrapper.Fixed(thoughts_btn, MENU_BTN_SIZE)

	daily_btn.SetCallback(func () {
		handle_tab_show(g.DAILY_TAB_ID)
		handle_widgets_show(g.DAILY_TAB_ID)
	})
	theme_btn.SetCallback(func () {
		handle_tab_show(g.THEME_TAB_ID)
		handle_widgets_show(g.THEME_TAB_ID)
	})
	thoughts_btn.SetCallback(func () {
		handle_tab_show(g.THOUGHTS_TAB_ID)
		handle_widgets_show(g.THOUGHTS_TAB_ID)
	})

	menu_wrapper.End()
	return menu_wrapper
}

func create_main_section () fltk.Widget {
	main_width := calc_window_percentage_width(MAIN_WIDTH_PERCENTAGE)

	main_wrapper := fltk.NewFlex(
		calc_window_percentage_width(MENU_WIDTH_PERCENTAGE), 0, 
		main_width, window_height,
	)

	grp_wrapper := fltk.NewGroup(0, 0, main_width, window_height)

	tab_ptrs_map = make(map[int]*fltk.Group, g.NUM_TABS)
	tab_widgets_ids = make(map[int][3]int, g.NUM_TABS)
	for i := 0; i < g.NUM_TABS; i++ {
		curr_tab_id := (g.TAB_ID_PRIME << i)
		curr_tab := fltk.NewGroup(0, 0, main_width, window_height)
		
		handle_tab_create(curr_tab_id, curr_tab)
		curr_tab.Hide()
		curr_tab.End()

		grp_wrapper.Add(curr_tab)
		tab_ptrs_map[curr_tab_id] = curr_tab
	}

	grp_wrapper.End()
	main_wrapper.Add(grp_wrapper)

	main_wrapper.End()
	return main_wrapper
}

func create_widget_section () fltk.Widget {
	widgets_width := calc_window_percentage_width(WIDGETS_WIDTH_PERCENTAGE)
	widgets_height := calc_window_percentage_height(WIDGETS_HEIGHT_PERCENTAGE)

	widgets_wrapper := fltk.NewGrid(
		calc_window_percentage_width(MENU_WIDTH_PERCENTAGE) + calc_window_percentage_width(MAIN_WIDTH_PERCENTAGE), 
		0, 
		widgets_width, window_height,
	)
	widgets_wrapper.SetLayout(3, 1, 0, 0)

	widgets_ptr_map = make(map[int]*fltk.Group, g.NUM_WIDGETS)
	for i := 0; i < g.NUM_WIDGETS; i++ {
		curr_widget_id := (g.WIDGET_ID_PRIME << i)
		curr_widget := fltk.NewGroup(
			widgets_wrapper.X(), widgets_wrapper.Y(), 
			widgets_width, widgets_height,
		)

		handle_widget_create(curr_widget_id, curr_widget)
		curr_widget.Hide()
		curr_widget.End()

		widgets_wrapper.Add(curr_widget)
		widgets_ptr_map[curr_widget_id] = curr_widget
	}

	widgets_wrapper.End()
	widgets_grid_ptr = widgets_wrapper
	return widgets_wrapper
}
