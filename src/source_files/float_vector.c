#include "../header_files/float_vector.h"
#include <stdlib.h>
#include <math.h>
#include <string.h>
#include <stdio.h>


//Preemptive declaration of functions
int fvector_add(FloatVector *instance, FloatVector *other);
FloatVector *fvector_add_ret(FloatVector *instance, FloatVector *other);
int fvector_sub(FloatVector *instance, FloatVector *other);
FloatVector *fvector_sub_ret(FloatVector *instance, FloatVector *other);
int fvector_multiply_scalar(FloatVector *instance, float n);
FloatVector *fvector_multiply_scalar_ret(FloatVector *instance, float n);
float fvector_dot_product(FloatVector *instance, FloatVector *other);
float fvector_norm(FloatVector *instance);
void fvector_normalize(FloatVector *instance);
FloatVector *fvector_make_copy(FloatVector *instance);

//VTable declaration
static const FloatVectorVTable vtable = {
    .add = fvector_add,
    .add_ret = fvector_add_ret,
    .sub = fvector_sub,
    .sub_ret = fvector_sub_ret,
    .multiply_scalar = fvector_multiply_scalar,
    .multiply_scalar_ret = fvector_multiply_scalar_ret,
    .dot_product = fvector_dot_product,
    .norm = fvector_norm,
    .normalize = fvector_normalize,
    .make_copy = fvector_make_copy
};

/**
 * Function that creates a new FloatVector instance of capacity 
 * 'capacity' and length 0 and returns a pointer to it. If any
 * error is encountered a NULL pointer is returned instead.
 */
FloatVector *fvector_create(float x, float z) {
    FloatVector *new_instance = (FloatVector *)malloc(sizeof(FloatVector));
    if (new_instance == NULL) {
        return NULL;
    }

    new_instance->x = x;
    new_instance->z = z;
    new_instance->methods = &vtable;

    return new_instance;
}

/**
 * Procedures that frees all memory allocated to a FloatVector instance.
 */
void fvector_destroy(FloatVector *instance) {
    if (instance == NULL) {
        return; //Nothing to do here
    }

    free(instance);
    return;
}

/**
 * Function that returns a pointer to a deep copy of the calling struct instance.
 * In case any error is found a NULL pointer is returned instead.
 */
FloatVector *fvector_make_copy(FloatVector *instance) {
    if (instance == NULL) {
        return NULL;
    }

    FloatVector *new_instance = fvector_create(instance->x, instance->z);
    if (new_instance == NULL) {
        return NULL;
    }

    return new_instance;
}

/**
 * Function that recieves another FloatVector instance pointer and adds
 * it to the calling FloatVector instance. This operation is destructive to
 * the calling instance, meaning the calling struct is overwritten with the
 * result instead of giving a return value. Returns 0 on success and -1 on
 * the occurence of any error.
 */
int fvector_add(FloatVector *instance, FloatVector *other) {
    if (instance == NULL || other == NULL) {
        return -1;
    }

    instance->x += other->x;
    instance->z += other->z;

    return 0;
}

/**
 * Function that returns a pointer to a new FloatVector instance equal to
 * the sum of the calling instance and the inputted 'other' vector.
 * In case of any error a NULL pointer is returned instead.
 */
FloatVector *fvector_add_ret(FloatVector *instance, FloatVector *other) {
    if (instance == NULL || other == NULL) {
        return NULL;
    }

    FloatVector *result = fvector_create(instance->x + other->x, instance->z + other->z);

    return result;
}

/**
 * Function that recieves another FloatVector instance pointer and substracts
 * it to the calling FloatVector instance. This operation is destructive to
 * the calling instance, meaning the calling struct is overwritten with the
 * result instead of giving a return value. Returns 0 on success and -1 on
 * the occurence of any error.
 */
int fvector_sub(FloatVector *instance, FloatVector *other) {
    if (instance == NULL || other == NULL) {
        return -1;
    }

    instance->x -= other->x;
    instance->z -= other->z;

    return 0;
}

/**
 * Function that returns a pointer to a new FloatVector instance equal to
 * the substraction of the calling instance and the inputted 'other' vector.
 * In case of any error a NULL pointer is returned instead.
 */
FloatVector *fvector_sub_ret(FloatVector *instance, FloatVector *other) {
    if (instance == NULL || other == NULL) {
        return NULL;
    }

    FloatVector *result = fvector_create(instance->x - other->x, instance->z - other->z);

    return result;
}

/**
 * Function that multiplies the calling instance by a scalar value passed in
 * as an argument. This operation is destructive to the calling struct instance,
 * meaning it modifies it's values directly instead of returning a result value.
 * On success the function returns 0, if any error is encountered, -1 is returned.
 */
int fvector_multiply_scalar(FloatVector *instance, float n) {
    if (instance == NULL) {
        return -1;
    }

    instance->x *= n;
    instance->z *= n;

    return 0;
}

/**
 * Function that returns a pointer to a new FloatVector instance equal
 * to the multiplication of the calling instance by the scalar value n.
 * If any errors are encountered a NULL pointer is returned instead.
 */
FloatVector *fvector_multiply_scalar_ret(FloatVector *instance, float n) {
    
    if (instance == NULL) {
        return NULL;
    }

    FloatVector *result = fvector_create(instance->x * n, instance->z * n);

    return result;
}

/**
 * Function that returns a float equal to the dot product of the calling 
 * instance and the inputted 'other' FloatVector. If any errors are encountered
 * NAN is returned instead.
 */
float fvector_dot_product(FloatVector *instance, FloatVector *other) {
    if (instance == NULL || other == NULL) {
        return NAN;
    }

    return instance->x * other->x + instance->z * other->z;
}

/**
 * Function that returns a float value equal to the norm of the calling instance.
 * If any errors are enocountered then NAN is returned instead.
 */
float fvector_norm(FloatVector *instance) {
    if (instance == NULL) {
        return NAN;
    }

    float norm_sq = instance->x * instance->x + instance->z * instance->z;

    return sqrt(norm_sq);
}

/**
 * Function that normlizes the calling FloatVector instance. This operation is
 * destructive and directly modifies the calling struct. The passwd in FloatVector
 * pointer is assumed to be valid
 */
void fvector_normalize(FloatVector *instance) {
    float norm = instance->methods->norm(instance);

    if (norm != 1.0f) {
        instance->methods->multiply_scalar(instance, 1.0f / norm);
    }
    return;
}

/**
 * Function that creates a new orientation value for kinematic
 * algorithms where angular velocity isn't taken into account.
 * The new orientation value will simply follow the orientation
 * of the Static argument velocity vector.
 */
float new_orientation(float current, FloatVector *velocity) {
    if (velocity->methods->norm(velocity) > 0) {
        return atan2f(velocity->x, velocity->z);
    }

    return current;
}

/**
 * Function that takes an orientation expressed as an angle against the z axis
 * and returns a vector pointing in the orientation of the given angle. The returned
 * vector will have a length of 1.
 */
FloatVector *orientation_to_vector(float orientation) {
    return fvector_create(sinf(orientation), cosf(orientation));
}