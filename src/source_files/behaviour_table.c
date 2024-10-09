#include "../header_files/game_character.h"
#include "../header_files/path.h"
#include "../header_files/wander_target.h"
#include "../header_files/evasion.h"
#include <stdio.h>
#include <stdlib.h>
#include <math.h>
#define PIF 3.14f

/**
 * Util function that returns a random floating point value
 * between 0 and 1. Not much to explain here
 */
float random_binomial(void) {
    return ((float)rand() / (float)RAND_MAX) - ((float)rand() / (float)RAND_MAX);
}

/**
 * Util function that maps a radian angle to the range [-pi, pi].
 */
float map_angle_to_range(float angle) {
    return fmodf((angle + PIF), (2.0f * PIF)) - PIF;
}


/**
 * Function that returns a pointer to a SteeringOutput struct instance
 * containing the necessary movement corrections for a wandering behaviour.
 * It is the caller's responsability to free the memory allocated to the
 * result struct after using it. The function returns NULL in case of any
 * failure.
 */
SteeringOutput *enemy_kinematic_wander(GameCharacter *instance, MovementInfo *target_ptr) {
    SteeringOutput *result = (SteeringOutput *)malloc(sizeof(SteeringOutput));
    if (result == NULL) {
        return NULL;
    }

    float static const max_speed = 1.0f;
    float static const max_rotation = 4.71f; //3pi/2

    result->linear = orientation_to_vector(instance->movement->orientation);
    result->linear->methods->multiply_scalar(result->linear, max_speed);

    result->angular = random_binomial() * max_rotation;

    return result;
}

/**
 * Function that returns a pointer to a SteeringOutput struct
 * containing the necessary movement corrections for a seek
 * behaviour. It is the caller's responsability to free the
 * SteeringOutput pointer after calling.
 */
SteeringOutput *enemy_kinematic_seek(GameCharacter *instance, MovementInfo *target_ptr) {
    Static *target = (Static *)target_ptr;
    SteeringOutput *result = (SteeringOutput *)malloc(sizeof(SteeringOutput));
    result->angular = 0.0f;
    result->linear = NULL;

    static const float max_speed = 1.0f;

    FloatVector *target_pos = target->position;
    FloatVector *ch_pos = instance->movement->position;

    result->linear = target_pos->methods->sub_ret(target_pos, ch_pos);
    result->linear->methods->normalize(result->linear);
    result->linear->methods->multiply_scalar(result->linear, max_speed);

    instance->movement->orientation = new_orientation(instance->movement->orientation, result->linear);

    result->angular = 0;
    return result;
}

/**
 * Function that returns a pointer to a SteeringOutput struct
 * containing the necessary movement corrections for a flee
 * behaviour. It is the caller's responsability to free the
 * SteeringOutput pointer after calling.
 */
SteeringOutput *enemy_kinematic_flee(GameCharacter *instance, MovementInfo *target_ptr) {
    Static *target = (Static *)target_ptr;
    SteeringOutput *result = (SteeringOutput *)malloc(sizeof(SteeringOutput));
    result->angular = 0.0f;
    result->linear = NULL;

    static const float max_speed = 1.0f;
    static const float max_distance = 3.0f;

    FloatVector *target_pos = target->position;
    FloatVector *ch_pos = instance->movement->position;

    result->linear = ch_pos->methods->sub_ret(ch_pos, target_pos);
    if (result->linear->methods->norm(result->linear) > max_distance) {
        //If character is too far away from target come to a halt
        fvector_destroy(result->linear);
        result->linear = instance->movement->velocity->methods->multiply_scalar_ret(instance->movement->velocity, -1.0f);
        instance->movement->orientation = new_orientation(instance->movement->orientation, result->linear);
        return result;
    }
    result->linear->methods->normalize(result->linear);
    result->linear->methods->multiply_scalar(result->linear, max_speed);

    instance->movement->orientation = new_orientation(instance->movement->orientation, result->linear);

    result->angular = 0;
    return result;
}

/**
 * Function that returns a pointer to a SteeringOutput struct instance
 * containing the necessary movement corrections for Kinematic Arrive 
 * behaviour. It is the caller's responsability to destroy the SteeringOutput
 * instance after ussage.
 */
SteeringOutput *enemy_kinematic_arrive(GameCharacter *instance, MovementInfo *target_ptr) {
    Static *target = (Static *)target_ptr;
    SteeringOutput *result = (SteeringOutput *)malloc(sizeof(SteeringOutput));
    if (result == NULL) {
        return NULL; //I should probably return some kind of error but I'm not making a special struct for that
    }
    result->angular = 0.0f;
    result->linear = NULL;

    static const float max_speed = 1.0f;
    static const float arrive_radius = 2.0f;
    static const float time_to_target = 0.25f;

    result->linear = target->position->methods->sub_ret(target->position, instance->movement->position);
    if (result->linear->methods->norm(result->linear) < arrive_radius) {
        result->linear = fvector_create(0.0f, 0.0f);
        return result;
    }

    result->linear->methods->multiply_scalar(result->linear, 1.0f / time_to_target);

    if (result->linear->methods->norm(result->linear) > max_speed) {
        result->linear->methods->normalize(result->linear);
        result->linear->methods->multiply_scalar(result->linear, max_speed);
    }

    instance->movement->orientation = new_orientation(instance->movement->orientation, result->linear);
    result->angular = 0.0f;

    return result;
}

/**
 * Function that returns a pointer to a SteeringOutput struct instance
 * containing the necessary movement corrections for a Dynamic Seek behaviour.
 * It is the caller's responsability to free the memory allocated to the result
 * struct after using it. The function returns NULL in case of any failure.
 */
SteeringOutput *enemy_dynamic_seek(GameCharacter *instance, MovementInfo *target_ptr) {
    Kinematic *target = (Kinematic *)target_ptr;

    const static float max_acceleration = 4.0f;
    SteeringOutput *result = (SteeringOutput *)malloc(sizeof(SteeringOutput));
    if (result == NULL) {
        return NULL;
    }

    result->linear = target->position->methods->sub_ret(target->position, instance->movement->position);
    if (result == NULL) {
        free(result);
        return NULL;
    }

    result->linear->methods->normalize(result->linear);
    result->linear->methods->multiply_scalar(result->linear, max_acceleration);
    result->angular = 0.0f;

    return result;
}

/**
 * Function that returns a pointer to a SteeringOutput struct instance
 * containing the necessary movement corrections for a Dynamic Flee behaviour.
 * It is the caller's responsability to free the memory allocated to the result
 * struct after using it. The function returns NULL in case of any failure.
 */
SteeringOutput *enemy_dynamic_flee(GameCharacter *instance, MovementInfo *target_ptr) {
    Kinematic *target = (Kinematic *)target_ptr;

    const static float max_acceleration = 4.0f;
    SteeringOutput *result = (SteeringOutput *)malloc(sizeof(SteeringOutput));
    if (result == NULL) {
        return NULL;
    }

    result->linear = instance->movement->position->methods->sub_ret(instance->movement->position, target->position);
    if (result == NULL) {
        free(result);
        return NULL;
    }

    result->linear->methods->normalize(result->linear);
    result->linear->methods->multiply_scalar(result->linear, max_acceleration);
    result->angular = 0.0f;

    return result;
}

/**
 * Function that returns a pointer to a SteeringOutput struct instance
 * containing the necessary movement corrections for a Dynamic Arrive behaviour.
 * It is the caller's responsability to free the allocated memory after
 * it is utilized. In the occurrence of an error a value of NULL is 
 * returned.
 */
SteeringOutput *enemy_dynamic_arrive(GameCharacter *instance, MovementInfo *target_ptr) {
    Kinematic *target = (Kinematic *)target_ptr;

    const float max_acceleration = 1.0f;
    const float max_speed = 1.0f;
    const float target_radius = 0.1f;
    const float slow_radius = 4.0f;
    const float time_to_target = 0.1f;

    SteeringOutput *result = (SteeringOutput *)malloc(sizeof(SteeringOutput));
    if (result == NULL) {
        return NULL;
    }

    FloatVector *direction = target->position->methods->sub_ret(target->position, instance->movement->position);
    float distance = direction->methods->norm(direction);

    if (distance < target_radius) {
        result->linear = fvector_create(0.0f, 0.0f);
        result->angular = 0;
        fvector_destroy(direction);
        return result;
    }

    float target_speed;
    if (distance > slow_radius) {
        target_speed = max_speed;
    } else {
        target_speed = max_speed * (distance - target_radius) / slow_radius;
    }

    FloatVector *target_velocity = direction;
    target_velocity->methods->normalize(target_velocity);
    target_velocity->methods->multiply_scalar(target_velocity, target_speed);

    result->linear = target_velocity->methods->sub_ret(target_velocity, instance->movement->velocity);
    result->linear->methods->multiply_scalar(result->linear, 1.0f/time_to_target);

    if (result->linear->methods->norm(result->linear) > max_acceleration) {
        result->linear->methods->normalize(result->linear);
        result->linear->methods->multiply_scalar(result->linear, max_acceleration);
    }

    result->angular = 0.0f;
    fvector_destroy(target_velocity);
    return result;
}

/**
 * Function that returns a pointer to a SteeringOutput struct containing
 * the necessary movement corrections for an Align behaviour. It is the
 * caller's responsability to free the memory allocated to the result
 * struct after the function is called. In case of any error a NULL pointer
 * is returned.
 */
SteeringOutput *enemy_align(GameCharacter *instance, MovementInfo *info) {
    Kinematic *target = (Kinematic *)info;

    const float max_angular_acceleration = 2.0f * PIF;
    const float max_rotation = PIF;

    const float slow_radius = PIF / 4.0f;
    const float target_radius = PIF / 64.0f;
    const float time_to_target = 0.1f;

    SteeringOutput *result = (SteeringOutput *)malloc(sizeof(SteeringOutput));
    float rotation = target->orientation - instance->movement->orientation;
    rotation = map_angle_to_range(rotation);
    float rotation_size = fabsf(rotation);

    if (rotation_size < target_radius) {
        result->linear = fvector_create(0.0f, 0.0f);
        result->angular = 0.0f;
        return result;
    }

    float target_rotation;
    if (rotation_size > slow_radius) {
        target_rotation = max_rotation;
    } else {
        target_rotation = max_rotation * (rotation_size - target_radius)/ slow_radius;
    }
    target_rotation *= rotation / rotation_size;

    result->angular = target_rotation - instance->movement->angular_velocity;
    result->angular /= time_to_target;

    float angular_acceleration = fabsf(result->angular);
    if (angular_acceleration > max_angular_acceleration) {
        result->angular /= angular_acceleration;
        result->angular *= max_angular_acceleration;
    }

    result->linear = fvector_create(0.0f, 0.0f);
    return result;
}

/**
 * Function that returns a pointer to a SteeringOutput struct containing
 * the necessary movement corrections for a Face behaviour. It is
 * the caller's responsability to free the allocated memory for the 
 * result struct after calling. In case of any errors NULL is returned.
 */
SteeringOutput *enemy_face(GameCharacter *instance, MovementInfo *blob) {
    Kinematic *target = (Kinematic *)blob;

    FloatVector *direction = target->position->methods->sub_ret(target->position, instance->movement->position);
    if (direction->methods->norm(direction) == 0) {
        SteeringOutput *result = (SteeringOutput *)malloc(sizeof(SteeringOutput));
        result->linear = fvector_create(0.0f, 0.0f);
        result->angular = 0.0f;
    }

    //Create new target for align
    FloatVector *target_pos = target->position;
    FloatVector *target_vel = target->velocity;
    Kinematic *new_target = new_kinematic(target_pos->x, target_pos->z, target_vel->x, target_vel->z, target->orientation, target->angular_velocity);
    new_target->orientation = atan2f(direction->x, direction->z);

    SteeringOutput *result = enemy_align(instance, (MovementInfo *)new_target);
    destroy_kinematic(new_target);

    return result;
}

/**
 * Function that returns a pointer to a SteeringOutput struct
 * containing the necessary movement corrections for a Velocity
 * Matching behaviour. It is the caller's responsability to free
 * the memory allocated to the result struct after calling. NULL
 * is returned in case of any errors.
 */
SteeringOutput *enemy_velocity_matching(GameCharacter *instance, MovementInfo *data) {
    Kinematic *target = (Kinematic *)data;

    const float max_acceleration = 3.0f;
    const float time_to_target = 0.1f;

    SteeringOutput *result = (SteeringOutput *)malloc(sizeof(SteeringOutput));
    if (result == NULL) {
        return NULL;
    }

    result->linear = target->velocity->methods->sub_ret(target->velocity, instance->movement->velocity);
    result->linear->methods->multiply_scalar(result->linear, 1.0f / time_to_target);

    if (result->linear->methods->norm(result->linear) > max_acceleration) {
        result->linear->methods->normalize(result->linear);
        result->linear->methods->multiply_scalar(result->linear, max_acceleration);
    }

    result->angular = 0.0f;
    return result;
}

/**
 * Function that returns a pointer to a SteeringBehaviour struct instance
 * containing the necessary movement corrections for a Look Where You're
 * Going behaviour. It is the caller's responsability to free the memory
 * allocated to the result struct after calling. In case of any errors
 * NULL is returned.
 */
SteeringOutput *enemy_look_where_youre_going(GameCharacter *instance, MovementInfo *blob) {
    Kinematic *target = (Kinematic *)blob;

    FloatVector *velocity = instance->movement->velocity;
    if (velocity->methods->norm(velocity) == 0) {
        SteeringOutput *result = (SteeringOutput *)malloc(sizeof(SteeringOutput));
        result->linear = fvector_create(0.0f, 0.0f);
        result->angular = 0.0f;
        return result;
    }

    // Put new target together for Align
    FloatVector *target_pos = target->position;
    FloatVector *target_vel = target->velocity;

    Kinematic *new_target = new_kinematic(target_pos->x, target_pos->z, target_vel->x, target_vel->z, target->orientation, target->angular_velocity);
    new_target->orientation = atan2f(velocity->x, velocity->z);

    SteeringOutput *result = enemy_align(instance, (MovementInfo *)new_target);
    destroy_kinematic(new_target);
    return result;
}

/**
 * Function that returns a pointer to a SteeringOutput struct
 * containing the necessary movement corrections for a Pursue
 * behaviour. It is the caller's responsability to free the memory
 * allocated to the result struct after calling the function.
 * In case of any error NULL is returned.
 */
SteeringOutput *enemy_delegated_behaviour_pursue(GameCharacter *instance, MovementInfo *blob) {
    Kinematic *target = (Kinematic *)blob;

    const float max_prediction = 2.0f;

    FloatVector *direction = target->position->methods->sub_ret(target->position, instance->movement->position);
    const float distance = direction->methods->norm(direction);

    const float speed = instance->movement->velocity->methods->norm(instance->movement->velocity);

    float prediction;
    if (speed <= distance / max_prediction) {
        prediction = max_prediction;
    } else {
        prediction = distance / speed;
    }

    //Put new target for seek together
    FloatVector *target_pos = target->position;
    FloatVector *target_vel = target->velocity;

    Kinematic *new_target = new_kinematic(target_pos->x, target_pos->z, target_vel->x, target_vel->z, target->orientation, target->angular_velocity);
    new_target->position->methods->multiply_scalar(new_target->position, prediction);

    SteeringOutput *result = enemy_dynamic_seek(instance, (MovementInfo *)new_target);

    //Get orientation corrections from Look Where You're Going
    SteeringOutput *orientation_result = enemy_look_where_youre_going(instance,  (MovementInfo *)new_target);
    result->angular = orientation_result->angular;
    destroy_steering_output(orientation_result);
    destroy_kinematic(new_target);

    return result;
}

/**
 * Function that returns a pointer to a SteeringOutput struct
 * containing the necessary movement corrections for a Flee
 * behaviour. It is the caller's responsability to free the memory
 * allocated to the result struct after calling the function.
 * In case of any error NULL is returned.
 */
SteeringOutput *enemy_delegated_behaviour_evade(GameCharacter *instance, MovementInfo *blob) {
    Kinematic *target = (Kinematic *)blob;

    const float max_prediction = 2.0f;

    FloatVector *direction = target->position->methods->sub_ret(target->position, instance->movement->position);
    const float distance = direction->methods->norm(direction);

    const float speed = instance->movement->velocity->methods->norm(instance->movement->velocity);

    float prediction;
    if (speed <= distance / max_prediction) {
        prediction = max_prediction;
    } else {
        prediction = distance / speed;
    }

    //Put new target for seek together
    FloatVector *target_pos = target->position;
    FloatVector *target_vel = target->velocity;

    Kinematic *new_target = new_kinematic(target_pos->x, target_pos->z, target_vel->x, target_vel->z, target->orientation, target->angular_velocity);
    new_target->position->methods->multiply_scalar(new_target->position, prediction);

    SteeringOutput *result = enemy_dynamic_flee(instance, (MovementInfo *)new_target);

    //Get orientation corrections from Look Where You're Going
    SteeringOutput *orientation_result = enemy_look_where_youre_going(instance,  (MovementInfo *)new_target);
    result->angular = orientation_result->angular;
    destroy_steering_output(orientation_result);
    destroy_kinematic(new_target);

    return result;
}

/**
 * Function that returns a pointer to a SteeringOutput struct
 * containing the necessary movement corrections for a Dynamic
 * Wandering behaviour. It is the caller's responsability to 
 * later free the memory associated to the returned result struct.
 * In the case of any errors NULL is returned.
 */
SteeringOutput *enemy_dynamic_wandering(GameCharacter *instance, MovementInfo *blob) {
    WanderTarget *wander_target = (WanderTarget *)blob;
    const float wander_rate = 0.2f;
    const float max_acceleration = 2.0f;

    wander_target->wander_orientation += random_binomial() * wander_rate;
    float target_orientation = wander_target->wander_orientation + instance->movement->orientation;

    //target = character.position + wanderOffSet * character.position.asVector()
    FloatVector *character_position = instance->movement->position;
    FloatVector *character_orientation_vec = orientation_to_vector(instance->movement->orientation); //heap allocated
    character_orientation_vec->methods->multiply_scalar(character_orientation_vec, wander_target->wander_offset);
    Kinematic *face_target = new_kinematic(character_position->x, character_position->z, 0.0f, 0.0f, 0.0f, 0.0f); //heap allocated
    face_target->position->methods->add(face_target->position, character_orientation_vec);
    fvector_destroy(character_orientation_vec); //heap free

    //target += wanderRadius * targetOrientation.asVector()
    FloatVector *target_orientation_vec = orientation_to_vector(target_orientation); //heap allocated
    target_orientation_vec->methods->multiply_scalar(target_orientation_vec, wander_target->wander_radius);
    face_target->position->methods->add(face_target->position, target_orientation_vec);
    fvector_destroy(target_orientation_vec); //heap free

    //result.linear = maxAccelertion * character.orientation.asVector()
    SteeringOutput *result = enemy_face(instance, (MovementInfo *)face_target);
    result->linear->methods->multiply_scalar(result->linear, max_acceleration);
    FloatVector *character_orientation_vec2 = orientation_to_vector(instance->movement->orientation);
    character_orientation_vec2->methods->multiply_scalar(character_orientation_vec2, max_acceleration);
    result->linear->methods->add(result->linear, character_orientation_vec2);
    fvector_destroy(character_orientation_vec2);

    destroy_kinematic(face_target);  // heap free

    return result;
}

/**
 * Function that returns a pointer to a SteeringOutput struct containing
 * the necessary movement correctinos for a Path Following behaviour.
 * It is the caller's responsability to free the memory after calling
 * the function. In case of any error a NULL pointer is returned.
 */
SteeringOutput *enemy_path_following(GameCharacter *instance, MovementInfo *blob) {
    Path *path = (Path *)blob;
    const float path_offset = 0.05f;

    path->last_param = path->get_param(path, instance->movement->position, path->last_param);
    const float target_param = path->last_param + path_offset;

    //Build new target to delegate to Seek
    FloatVector *target_position = path->get_position(path, target_param);
    Kinematic *target = new_kinematic(target_position->x, target_position->z, 0.0f, 0.0f, 0.0f, 0.0f);
    fvector_destroy(target_position);

    SteeringOutput *result = enemy_dynamic_seek(instance, (MovementInfo *)target);
    destroy_kinematic(target);
    return result;
}

/**
 * Function that returns a pointer to a SteeringOutput struct containing
 * the necessary movement corrections for a single individual of an Evasion
 * behaviour. The Evasion behaviour is combined with Velocity Matching in order
 * for the effect to be shown in demo. It is the caller's responsability to
 * free the memory allocated to the result struct after use. Returns NULL on any
 * error.
 */
SteeringOutput *enemy_evasion(GameCharacter *instance, MovementInfo *blob) {
    Evasion *evasion = (Evasion *)blob;

    const float max_acceleration = 2.0f;
    // First get velocity matching corrections
    SteeringOutput *result = enemy_velocity_matching(instance, (MovementInfo *)evasion->vel_match_target);

    //Correct it using evasion algorithm
    for (int i = 0; i < evasion->target_count; i++) {
        FloatVector *target_pos = evasion->targets[i]->position;
        FloatVector *direction = target_pos->methods->sub_ret(target_pos, instance->movement->position);
        float distance = direction->methods->norm(direction);

        if (distance < evasion->thresshold) {
            float strength = fminf(evasion->decay_coefficient / (distance * distance), max_acceleration);
            direction->methods->normalize(direction);
            direction->methods->multiply_scalar(direction, strength);
            result->linear->methods->add(result->linear, direction);
        }
    }

    return result;
}

//VTable declaration 
const BehaviourVTable enemy_vtable = {
    .seek = enemy_kinematic_seek,
    .flee = enemy_kinematic_flee,
    .arrive = enemy_kinematic_arrive,
    .wander = enemy_kinematic_wander,
    .dynamic_seek = enemy_dynamic_seek,
    .dynamic_flee = enemy_dynamic_flee,
    .dynamic_arrive = enemy_dynamic_arrive,
    .dynamic_align = enemy_align,
    .pursue = enemy_delegated_behaviour_pursue,
    .evade = enemy_delegated_behaviour_evade,
    .face = enemy_face,
    .path_following = enemy_path_following,
    .dynamic_wander = enemy_dynamic_wandering
};