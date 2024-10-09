#include "../header_files/evasion.h"
#include <stdlib.h>

/**
 * Function that creates a new heap-allocated Evasion struct
 * instance and returns a pointer to it. The members array
 * passed as an argument is an array of all the game
 * characters from the world who will be members of the
 * separation group. Returns NULL on error.
 */
Evasion *new_evasion(GameCharacter **members, unsigned int member_count) {
    Evasion *new_instance = (Evasion *)malloc(sizeof(Evasion));
    if (new_instance == NULL) {
        return NULL;
    }

    Kinematic **new_instance_members = (Kinematic **)malloc(sizeof(Kinematic *) * member_count);
    if (new_instance_members == NULL) {
        free(new_instance);
        return NULL;
    }

    for (unsigned int i = 0; i < member_count; i++) {
        new_instance_members[i] = members[i]->movement;
    }

    new_instance->decay_coefficient = 0.5f;
    new_instance->thresshold = 1.0f;

    return new_instance;
}

/**
 * Function that recieves a pointer to a heap-allocated Evasion
 * struct instance and frees all memory allocated to it. It's
 * worth noting that the pointers in the targets array are shared
 * with the GameCharacter objects the struct was constructed with, 
 * so we don't free them as they will be freed when the GameCharacter
 * objects are destroyed.
 */
void destroy_evasion(Evasion *instance) {
    if (instance == NULL) {
        return;
    }

    if (instance->targets != NULL) {
        free(instance->targets);
    }

    free(instance);
    return;
}