#ifndef PATH_H
    #define PATH_H

    #include "static.h"
    #include "sdl_manager.h"
    typedef struct {
        float (*x)(float t);
        float (*y)(float t);
        float param_skip;
    } PathCurve;

    typedef struct Path_s {
        PathCurve *path;
        float (*get_param)(struct Path_s *self, FloatVector *position, float last_param);
        FloatVector *(*get_position)(struct Path_s *self, float param);
        void (*draw_path)(struct Path_s *self, SDLManager *sdl);
        float last_param;
    } Path;

    Path *new_path(char *name);
    void destroy_path(Path *instance);

#endif