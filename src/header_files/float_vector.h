#ifndef FLOAT_VECTOR_H
    #define FLOAT_VECTOR_H

    #include "float_vector_vtable.h"
    typedef struct FloatVector_s {
        float x;
        float z;
        const FloatVectorVTable *methods;
    } FloatVector;

    float new_orientation(float current, FloatVector *velocity);
    FloatVector *orientation_to_vector(float orientation);
    FloatVector *fvector_create(float x, float z);
    void fvector_destroy(FloatVector *instance);

#endif