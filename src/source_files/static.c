#include "../header_files/static.h"
#include <stdlib.h>

/**
 * Function that allocates a new instance of the Static struct
 * and returns a pointer to it. This struct should not be instantiated
 * on its own and is usually utilized within other structs, which have
 * the responsability of freeing the returned memory when cleaning
 * after themselves. The initial values of the position vector and
 * rotation field are passed in as arguments
 */
Static *new_static(float xo, float zo, float orientation) {
    Static *new_instance = (Static *)malloc(sizeof(Static));
    new_instance->orientation = 0.0f;
    new_instance->position = NULL;

    FloatVector *new_instance_pos = fvector_create(xo, zo);

    new_instance->position = new_instance_pos;
    new_instance->orientation = orientation;

    return new_instance;
}

/**
 * Method that deallocates all memory allocated to a Static struct
 * instance. This method is usually called by the destructor of 
 * another struct that owns a Static struct instance, so it typically
 * isn't called by itself.
 */
void destroy_static(Static *instance) {
    if (instance == NULL) {
        return;
    }

    if (instance->position != NULL) {
        fvector_destroy(instance->position);
    }

    free(instance);
    return;
}