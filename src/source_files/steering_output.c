#include "../header_files/steering_output.h"
#include <stdlib.h>

/**
 * Function that creates a new SteeringOutput instance and
 * returns a pointer to it. Responsability to free the 
 * newly created instance is left to the caller.
 */
SteeringOutput *steering_output_create(void) {
    return NULL;
}

void destroy_steering_output(SteeringOutput *instance) {
    if (instance == NULL) {
        return;
    }

    if (instance->linear != NULL) {
        fvector_destroy(instance->linear);
        instance->linear = NULL;
    }

    free(instance);
    return;
}