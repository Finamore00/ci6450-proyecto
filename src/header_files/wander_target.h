#ifndef WANDER_TARGET_H
    #define WANDER_TARGET_H

    typedef struct WanderTarget_s {
        float wander_offset;
        float wander_radius;
        float wander_orientation;
    } WanderTarget;

    WanderTarget *new_wander_target(void); //For now we can make every Wander target the same
    void destroy_wander_target(WanderTarget *instance);

#endif