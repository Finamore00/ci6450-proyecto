#ifndef KINEMATIC_VTABLE_H
    #define KINEMATIC_VTABLE_H

    struct Kinematic_s;
    typedef struct Kinematic_s Kinematic;
    
    #include "steering_output.h"
    #include "float_vector.h"
    #include "static.h"
    typedef struct KinematicVTable_s {
        void (*update)(Kinematic *instance, SteeringOutput *steering, double time);
        Static *(*get_static)(Kinematic *instance);
    } KinematicVTable;

#endif