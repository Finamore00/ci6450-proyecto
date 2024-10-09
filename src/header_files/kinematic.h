#ifndef KINEMATIC_H
    #define KINEMATIC_H

    #include "kinematic_vtable.h"
    typedef struct Kinematic_s {
        const KinematicVTable *methods;
        FloatVector *position;
        FloatVector *velocity;
        float orientation;
        float angular_velocity;
    } Kinematic;

    Kinematic *new_kinematic(float xo, float zo, float vxo, float vzo, float orientation, float angular);
    void destroy_kinematic(Kinematic *instance);

#endif