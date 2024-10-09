#include "../header_files/path.h"
#include <math.h>
#include <string.h>
#include <SDL2/SDL2_gfxPrimitives.h>

/**
 * Parametric functions for the cartesian coordinates for a circle.
 */
float circle_x(float t) {
    return 3*cosf(t);
}

float circle_y(float t) {
    return 3*sinf(t);
}

const float circle_skip = 0.05f;

/**
 * Translates Path parameter values to a FloatVector position in the plane.
 */
FloatVector *path_get_position(Path *instance, float param) {
    float x = instance->path->x(param);
    float y = instance->path->y(param);

    return fvector_create(x, y);
}

/**
 * Function that, given a Path struct, the static position
 * of a navigator, and the last parameter explored by the
 * navigator, returns the closest parameter to the navigator
 * agent. In case of any error NAN is returned
 */
float path_get_param(Path *instance, FloatVector *navigator_pos, float last_param) {
    //Explore 10 points back and forward of the last explored point
    FloatVector *backward_pos;
    FloatVector *forward_pos;

    float min_distance_param = NAN;
    float min_distance = INFINITY;

    for (int i = 0; i < 10; i++) {
        float backward_param = last_param - ((float)i * instance->path->param_skip);
        float forward_param = last_param + ((float)i * instance->path->param_skip);

        backward_pos = instance->get_position(instance, backward_param);
        forward_pos = instance->get_position(instance, forward_param);

        //Calculate distance of the 2 explored positions to the navigator
        FloatVector *backward_separation = backward_pos->methods->sub_ret(backward_pos, navigator_pos);
        float backward_distance = backward_separation->methods->norm(backward_separation);
        fvector_destroy(backward_separation);

        FloatVector *forward_separation = forward_pos->methods->sub_ret(forward_pos, navigator_pos);
        float forward_distance = forward_separation->methods->norm(forward_separation);
        fvector_destroy(forward_separation);

        if (backward_distance < min_distance) {
            min_distance = backward_distance;
            min_distance_param = backward_param;
        }

        if (forward_distance < min_distance) {
            min_distance = forward_distance;
            min_distance_param = forward_param;
        }

        //Free position vectors for next iteration
        fvector_destroy(backward_pos);
        fvector_destroy(forward_pos);
        backward_pos = NULL;
        forward_pos = NULL;
    }

    return min_distance_param;
}

/**
 * Function that recieves a Path struct and draws it on screen
 * using the provided SDLManager struct. 
 */
void path_draw_circle(Path *instance, SDLManager *sdl) {
    SDL_Renderer *renderer = sdl->renderer;

    circleRGBA(renderer, 320, 240, 192, 255, 255, 255, 255);

    return;
}

/**
 * Function that takes in a curve name and returns a pointer
 * to a Path struct containing the information corresponding
 * to the inputted curve name. Available curves are "circle".
 * If any error is encountered or an inavlid name is inputted
 * a NULL pointer is returned
 */
Path *new_path(char *name) {
    if (!strcmp(name, "circle")) {
        PathCurve *new_instance_curve = (PathCurve *)malloc(sizeof(PathCurve));
        if (new_instance_curve == NULL) {
            return NULL;
        }
        Path *new_instance = (Path *)malloc(sizeof(Path));
        if (new_instance == NULL) {
            free(new_instance_curve);
            return NULL;
        }

        new_instance_curve->param_skip = circle_skip;
        new_instance_curve->x = circle_x;
        new_instance_curve->y = circle_y;

        new_instance->path = new_instance_curve;
        new_instance->last_param = 0.0f;
        new_instance->get_param = path_get_param;
        new_instance->get_position = path_get_position;
        new_instance->draw_path = path_draw_circle;

        return new_instance;
    } else {
        return NULL;
    }
}

/**
 * Function that frees all memory allocated to a heap-allocated
 * Path struct.
 */
void destroy_path(Path *instance) {
    if (instance == NULL) {
        return;
    }

    if (instance->path != NULL) {
        free(instance->path);
        instance->path = NULL;
    }

    free(instance);
    return;
}