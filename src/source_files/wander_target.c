#include "../header_files/wander_target.h"
#include <stdlib.h>

/**
 * Function that creates a new heap-allocated WanderTarget struct
 * instance and returns a pointer to it. For now the attributes
 * of the WanderTarget struct are hard-wired for all characters
 * that use it. This will change only if it's needed.
 */
WanderTarget *new_wander_target(void) {
    WanderTarget *new_instance = (WanderTarget *)malloc(sizeof(WanderTarget));
    if (new_instance == NULL) {
        return NULL;
    }

    new_instance->wander_offset = 1.0f;
    new_instance->wander_orientation = 0.0f;
    new_instance->wander_radius = 0.5f;

    return new_instance;
}

/**
 * Function that recieves a pointer to a heap-allocated instance
 * of a WanderTarget struct a frees all memory allocated to it.
 */
void destroy_wander_target(WanderTarget *instance) {
    if (instance == NULL) {
        return;
    }

    free(instance);
    return;
}