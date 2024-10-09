#ifndef STATIC_H
    #define STATIC_H

    #include "float_vector.h"
    typedef struct Static_s {
        FloatVector *position;
        float orientation;
    } Static;

    Static *new_static(float xo, float zo, float orientation);
    void destroy_static(Static *instance);

#endif