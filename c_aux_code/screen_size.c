#include <stdio.h>

#ifdef __linux__

#include <stdlib.h>
#include <string.h>
#include <X11/Xlib.h>
#define _session_type getenv("XDG_SESSION_TYPE")

int hello() {
    puts(_session_type);
    return 0;
}

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

