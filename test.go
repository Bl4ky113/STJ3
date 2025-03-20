package main

import (
    //"fmt"
	"strconv"

    "github.com/pwiecz/go-fltk"
    // #cgo LDFLAGS: -lX11
    // #include "./c_aux_code/screen_size.c"
    "C"
)

var i = 0

func main() {
    screen_width := int(C.get_width())
    screen_height := int(C.get_height())

	win := fltk.NewWindow(screen_width, screen_height)
	column := fltk.NewFlex(0, 0, screen_width, screen_height)
	column.SetType(fltk.COLUMN)
	column.SetGap(5)
	inc := fltk.NewButton(0, 0, 0, 0, "Increment")
	column.Fixed(inc, 40)
	box := fltk.NewBox(fltk.FLAT_BOX, 0, 0, 0, 0, "0")
	dec := fltk.NewButton(0, 0, 0, 0, "Decrement")
	inc.SetCallback(func() {
		i++
		box.SetLabel(strconv.Itoa(i))
	})
	dec.SetCallback(func() {
		i--
		box.SetLabel(strconv.Itoa(i))
	})

	column.Fixed(dec, 40)
	column.End()
	win.End()
	win.Resizable(column)
	win.Show()
	fltk.Run()
}
