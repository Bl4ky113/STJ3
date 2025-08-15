#ifndef __c_screen_utils__
#define __c_screen_utils__

#ifdef __linux__

#define _session_type getenv("XDG_SESSION_TYPE")

extern int get_width ();
extern int get_height ();

#endif /* __linux__ */

extern int calc_value_percentage (double percentage, int value);

#endif /* __c_screen_utils__ */
