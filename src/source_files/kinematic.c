#include "../header_files/kinematic.h"
#include <stdlib.h>
#include <stdio.h>
#define PIF 3.14

void kinematic_update(Kinematic *instance, SteeringOutput *steering, double time);
Static *kinematic_get_static(Kinematic *instance);

static const KinematicVTable vtable = {
    .update = kinematic_update,
    .get_static = kinematic_get_static
};              

/**
 * Function that updates the linear and angular velocities
 * stored in the Kinematic instance given a SteeringOutput
 * and a time delta. All inputted pointers are assumed to be
 * non-null.
 */
void kinematic_update(Kinematic *instance, SteeringOutput *steering, double time) {

    FloatVector *velocity_time = instance->velocity->methods->multiply_scalar_ret(instance->velocity, time);
    instance->position->methods->add(instance->position, velocity_time);
    fvector_destroy(velocity_time);

    //Bound position to map dimmensions
    instance->position->x = instance->position->x < -5.0 ? -5.0 : instance->position->x;
    instance->position->x = instance->position->x > 5.0 ? 5.0 : instance->position->x;
    instance->position->z = instance->position->z < -5.0 ? -5.0 : instance->position->z;
    instance->position->z = instance->position->z > 5.0 ? 5.0 : instance->position->z;

    instance->orientation += instance->angular_velocity * time;

    steering->linear->methods->multiply_scalar(steering->linear, time);
    instance->velocity->methods->add(instance->velocity, steering->linear);  // I don't know what the fuck is going on here

    if (instance->velocity->methods->norm(instance->velocity) > 1.0f) {
        instance->velocity->methods->normalize(instance->velocity);
    }

    instance->angular_velocity += steering->angular * time;

    return;
}

Static *kinematic_get_static(Kinematic *instance) {
    Static *new_instance = new_static(instance->position->x, instance->position->z, instance->orientation);
    
    return new_instance;
}

/**
 * Function that creates a new heap-allocated instance of the Kinematic struct
 * and returns a pointer to it. This function is rarely if ever called by itself
 * and is usually used by other structs that contain a Kinematic instance within
 * them. So yeah, I wouldn't call it.
 */
Kinematic *new_kinematic(float xo, float zo, float vxo, float vzo, float orientation, float angular) {
    Kinematic *new_instance = (Kinematic *)malloc(sizeof(Kinematic));

    FloatVector *new_instance_vel = fvector_create(vxo, vzo);

    FloatVector *new_instance_pos = fvector_create(xo, zo);

    new_instance->position = new_instance_pos;
    new_instance->velocity = new_instance_vel;
    new_instance->angular_velocity = angular;
    new_instance->orientation = orientation;
    new_instance->methods = &vtable;

    return new_instance;
}

/**
 * Method that deallocates all memory allocated to a Kinematic struct
 * instance. This method is usually called by the destructor of 
 * another struct that owns a Kinematic struct instance, so it typically
 * isn't called by itself.
 */
void destroy_kinematic(Kinematic *instance) {
    if (instance == NULL) {
        return;
    }

    if (instance->position != NULL) {
        fvector_destroy(instance->position);
        instance->position = NULL;
    }

    if (instance->velocity != NULL) {
        fvector_destroy(instance->velocity);
        instance->velocity = NULL;
    }

    free(instance);
    return;
}