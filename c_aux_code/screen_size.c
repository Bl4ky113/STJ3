
#ifndef __c_screen_utils__
#define __c_screen_utils__

#include <stdio.h>

#ifdef __linux__

#include <stdlib.h>
#include <X11/Xlib.h>
#define _session_type getenv("XDG_SESSION_TYPE")

int get_width () {
    Display *display;
    Window win;
    int screen_num, width, height;

    display = XOpenDisplay(0);
    screen_num = DefaultScreen(display);
    width = DisplayWidth(display, screen_num);
    height = DisplayHeight(display, screen_num);
        
    return width;
}

int get_height () {
    char *win_id;
    Display *display;
    Window win;
    int screen_num, width, height;

    display = XOpenDisplay(0);
    screen_num = DefaultScreen(display);
    height = DisplayHeight(display, screen_num);
    
    return height;
}

#endif /* __linux__ */

int calc_value_percentage (int percentage, int value) {
    return (int) (((double) value) * ((double) percentage / 100.0 ));
}

#endif /* __c_screen_utils__ */
