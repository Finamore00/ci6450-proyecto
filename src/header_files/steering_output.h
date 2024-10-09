#ifndef STEERING_OUTPUT_H
    #define STEERING_OUTPUT_H

    #include "float_vector.h"
    typedef struct SteeringOutput_s {
        FloatVector *linear;
        float angular;
    } SteeringOutput;

    SteeringOutput *new_steering_output(void);
    void destroy_steering_output(SteeringOutput *instance);

#endif