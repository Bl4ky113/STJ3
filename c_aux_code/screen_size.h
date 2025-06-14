#ifndef __c_screen_utils__
#define __c_screen_utils__

#include <stdio.h>

#ifdef __linux__

#include <stdlib.h>
#include <X11/Xlib.h>
#define _session_type getenv("XDG_SESSION_TYPE")

#include "./screen_size.c"

int get_width ();
int get_height ();

#endif /* __linux__ */

int calc_value_percentage (double percentage, int value);

#endif /* __c_screen_utils__ */
