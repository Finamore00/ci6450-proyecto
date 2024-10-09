#include "../header_files/game_character.h"
#include "../header_files/path.h"
#include "../header_files/wander_target.h"
#include "../header_files/evasion.h"
#include <stdlib.h>

typedef struct {
    int x;
    int z;
} PixelCoordinate;

extern const int WINDOW_WIDTH;
extern const int WINDOW_HEIGHT;

extern const BehaviourVTable enemy_vtable;

/**
 * ################################################
 * ################ Utilities #####################
 * ################################################
 */

/**
 * Function that translates the float vector position of a game
 * object to concrete pixel coordinates on the screen. The width
 * and height of the screen are hard-coded to 640x480, meaning the
 * functions that perform the horizontal and vertical translateions
 * are hard-coded as well. Used for the Character struct's draw
 * functions.
 */
PixelCoordinate float_to_pixel_pos(FloatVector *position) {
    PixelCoordinate result;

    //Function for horizontal translation is y = 64x + 320
    result.x = (int)(64 * position->x + 320);

    //Function for vertical translation is y = -48x + 240
    result.z = (int)(-48 * position->z + 240);

    return result;
}

/**
 * ################################################
 * ######### Player struct Functions ##############
 * ################################################
 */

/**
 * Function that creates a new heap-allocated Player struct
 * instance and returns a pointer to it.
 */
Player *new_player(unsigned int health) {
    Player *new_instance = (Player *)malloc(sizeof(Player));
    new_instance->health = health;

    return new_instance;
}

/**
 * Procedure that frees all memory allocated to a Player struct instance.
 */
void destroy_player(Player *instance) {
    if (instance == NULL) {
        return;
    }

    free(instance);
    return;
}

/**
 * Drawing method for the Player Character type. All drawing methods
 * return 0 on success and -1 on the occurence of an error.
 */
int player_draw(GameCharacter *instance, SDLManager *sdl) {
    PixelCoordinate ch_pos = float_to_pixel_pos(instance->movement->position);

    SDL_Rect player_sprite = {
        .x = ch_pos.x,
        .y = ch_pos.z,
        .h = 20,
        .w = 20,
    };

    SDL_SetRenderDrawColor(sdl->renderer, 0xFF, 0xA5, 0x00, 0x00); //Player is orange.
    SDL_RenderFillRect(sdl->renderer, &player_sprite);

    FloatVector *orientation_vector = orientation_to_vector(instance->movement->orientation);
    orientation_vector->methods->add(orientation_vector, instance->movement->position);
    PixelCoordinate orientation_vector_px = float_to_pixel_pos(orientation_vector);
    fvector_destroy(orientation_vector);
    SDL_SetRenderDrawColor(sdl->renderer, 0xFF, 0x00, 0x00, 0x00);
    SDL_RenderDrawLine(sdl->renderer, ch_pos.x + 10, ch_pos.z + 10, orientation_vector_px.x + 10, orientation_vector_px.z + 10);

    SDL_SetRenderDrawColor(sdl->renderer, 0xFF, 0xFF, 0xFF, 0xFF); //Return renderer back to white
    return 0;
}

/**
 * ################################################
 * ########## Enemy struct Functions ##############
 * ################################################
 */

/**
 * Function that creates a new heap-allocated Enemy struct
 * instance and returns a pointer to it. A NULL pointer is
 * returned if any error is encountered.
 */
Enemy *new_enemy(unsigned int health, float radius, EnemyBehaviour behaviour, void *misc) {
    Enemy *new_instance = (Enemy *)malloc(sizeof(Enemy));
    if (new_instance == NULL) {
        return NULL;
    }

    new_instance->type = behaviour;
    new_instance->arrive_radius = radius;
    new_instance->health = health;
    new_instance->behaviours = &enemy_vtable;
    new_instance->active = true;

    //Set behaviour
    switch (behaviour) {
        case Seek:
            new_instance->current_behaviour = enemy_vtable.seek;
            new_instance->movement_aux = NULL;
            break;
        case Flee:
            new_instance->current_behaviour = enemy_vtable.flee;
            new_instance->movement_aux = NULL;
            break;
        case Arrive:
            new_instance->current_behaviour = enemy_vtable.arrive;
            new_instance->movement_aux = NULL;
            break;
        case Wander:
            new_instance->current_behaviour = enemy_vtable.wander;
            new_instance->movement_aux = NULL;
            break;
        case DynamicSeek:
            new_instance->current_behaviour = enemy_vtable.dynamic_seek;
            new_instance->movement_aux = NULL;
            break;
        case DynamicFlee:
            new_instance->current_behaviour = enemy_vtable.dynamic_flee;
            new_instance->movement_aux = NULL;
            break;
        case DynamicArrive:
            new_instance->current_behaviour = enemy_vtable.dynamic_arrive;
            new_instance->movement_aux = NULL;
            break;
        case DynamicAlign:
            new_instance->current_behaviour = enemy_vtable.dynamic_align;
            new_instance->movement_aux = NULL;
            break;
        case Pursue:
            new_instance->current_behaviour = enemy_vtable.pursue;
            new_instance->movement_aux = (void *)misc;
            break;
        case Evade:
            new_instance->current_behaviour = enemy_vtable.evade;
            new_instance->movement_aux = (void *)misc;
            break;
        case Face:
            new_instance->current_behaviour = enemy_vtable.face;
            new_instance->movement_aux = NULL;
            break;
        case DynamicWander:
            new_instance->current_behaviour = enemy_vtable.dynamic_wander;
            new_instance->movement_aux = (void *)new_wander_target();
            break;
        case PathFollowing:
            new_instance->current_behaviour = enemy_vtable.path_following;
            char *path_name = (char *)misc;
            new_instance->movement_aux = (void *)new_path(path_name);
            break;
        case Separation:
            new_instance->current_behaviour = enemy_vtable.separation;
            EvasionCreateBlob *blob = (EvasionCreateBlob *)misc;
            new_instance->movement_aux = (void *)new_evasion(blob);
            break;
        default:
            new_instance->current_behaviour = enemy_vtable.wander;
            new_instance->movement_aux = NULL;
            break;
    }

    return new_instance;
}

/**
 * Method that frees all memory allocated to an Enemy struct 
 * instance.
 * 
 * Note: Right now it has some pretty bad memory leaks due to
 * the modifications to enemy structs and auxiliar data things.
 * Will correct it eventually.
 */
void destroy_enemy(Enemy *instance) {
    if (instance == NULL) {
        return;
    }

    instance->current_behaviour = NULL;
    free(instance);
    return;
}

/**
 * Drawing method for the Enemy Character type. All drawing methods
 * return 0 on success and -1 on the occurence of an error.
 */
int enemy_draw(GameCharacter *instance, SDLManager *sdl) {

    if (instance->enemy_info->active == false) {
        return 0; //Don't do anything if the enemy isn't active
    }

    PixelCoordinate ch_pos = float_to_pixel_pos(instance->movement->position);

    SDL_Rect enemy_sprite = {
        .x = ch_pos.x,
        .y = ch_pos.z,
        .h = 20,
        .w = 20,
    };

    SDL_SetRenderDrawColor(sdl->renderer, 0x00, 0x00, 0xFF, 0x00); //Enemy is blue.
    SDL_RenderFillRect(sdl->renderer, &enemy_sprite);

    //Enemy Orientation is denoted by a red line
    FloatVector *orientation_vector = orientation_to_vector(instance->movement->orientation);
    orientation_vector->methods->add(orientation_vector, instance->movement->position);
    PixelCoordinate orientation_vector_px = float_to_pixel_pos(orientation_vector);
    fvector_destroy(orientation_vector);

    SDL_SetRenderDrawColor(sdl->renderer, 0xFF, 0x00, 0x00, 0x00);
    SDL_RenderDrawLine(sdl->renderer, ch_pos.x + 10, ch_pos.z + 10, orientation_vector_px.x + 10, orientation_vector_px.z + 10);

    //If engaging in path following, draw the path
    if (instance->enemy_info->type == PathFollowing) {
        Path *path = (Path *)instance->enemy_info->movement_aux;
        path->draw_path(path, sdl);
    }

    SDL_SetRenderDrawColor(sdl->renderer, 0xFF, 0xFF, 0xFF, 0xFF);  // Return renderer back to white
    return 0;
}


/**
 * Function that creates a new heap-allocated Enemy type character
 * and returns a pointer to it. If any errors are encountered a 
 * NULL pointer is returned instead.
 */
GameCharacter *new_character_enemy(
    float xo,
    float zo,
    float vxo,
    float vzo,
    float orientation,
    float angular,
    unsigned int health,
    float radius,
    EnemyBehaviour behaviour,
    void *misc
) {
    GameCharacter *new_instance = (GameCharacter *)malloc(sizeof(GameCharacter));

    new_instance->type = Enemy_t;
    new_instance->movement = new_kinematic(xo, zo, vxo, vzo, orientation, angular);
    new_instance->enemy_info = new_enemy(health, radius, behaviour, misc);
    new_instance->player_info = NULL;
    new_instance->draw = enemy_draw;

    return new_instance;
}

/**
 * Function that creates a new heap-allocated Player-type character
 * and returns a pointer to it. If any errors are encountered a NULL
 * pointer is returned instead.
 */
GameCharacter *new_character_player(void) {
    GameCharacter *new_instance = (GameCharacter *)malloc(sizeof(GameCharacter));

    new_instance->type = Player_t;
    new_instance->movement = new_kinematic(0.0f, 0.0f, 0.0f, 0.0f, 0.0f, 0.0f);
    new_instance->player_info = new_player(500);
    new_instance->enemy_info = NULL;
    new_instance->draw = player_draw;

    return new_instance;
}

/**
 * Function that frees all memory allocated to an Enemy-type GameCharacter
 * struct. 
 */
void destroy_character(GameCharacter *instance) {
    if (instance == NULL) {
        return;
    }

    if (instance->movement != NULL) {
        destroy_kinematic(instance->movement);
        instance->movement = NULL;
    }

    if (instance->enemy_info != NULL) {
        destroy_enemy(instance->enemy_info);
        instance->enemy_info = NULL;
    }

    if (instance->player_info != NULL) {
        destroy_player(instance->player_info);
        instance->player_info = NULL;
    }

    free(instance);
    return;
}