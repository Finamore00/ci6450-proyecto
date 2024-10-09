#include "../header_files/game_manager.h"
#include "../header_files/evasion.h"
#include <SDL2/SDL.h>
#include <stdlib.h>
#include <stdio.h>

/**
####################################################################
######################### demo stuff ###############################
####################################################################
####################################################################
*/

/**
 * Function that takes in the GameManager instance under creation
 * and adds to it the necessary enemy characters for the kinematic
 * demo. The Kinematic demo contains 
 * 
 *  - 1 character performing Kineamtic Seek
 *  - 1 character performing Kineamtic Flee
 *  - 1 character performing Kinematic Arrive
 *  - 1 character performing Kinematic Wander
 */
void create_enemies_demo_kinematic(GameManager *instance) {
    instance->enemy_count = 4;
    instance->enemies = (GameCharacter **)malloc(sizeof(GameCharacter *) * instance->enemy_count);

    instance->enemies[0] = new_character_enemy(-3.0f, 2.1f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, Seek, (void *)NULL);
    instance->enemies[1] = new_character_enemy(1.0f, 1.3f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, Flee, (void *)NULL);
    instance->enemies[2] = new_character_enemy(0.0f, -4.5f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, Wander, (void *)NULL);
    instance->enemies[3] = new_character_enemy(4.5f, -4.5f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, Arrive, (void *)NULL);

    return;
}

/**
 * Function that takes in a GameManager instance under creation
 * and adds to it the necessary enemy characters for the dynamic
 * demo. The Dynamic demo contains
 * 
 *  - 1 character performing Dynamic Seek
 *  - 1 character performing Dynamic Flee
 *  - 1 character performing Dynamic Arrive
 */
void create_enemies_demo_dynamic(GameManager *instance) {
    instance->enemy_count = 3;
    instance->enemies = (GameCharacter **)malloc(sizeof(GameCharacter *) * instance->enemy_count);

    instance->enemies[0] = new_character_enemy(-2.0f, -2.0f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, DynamicSeek, (void *)NULL);
    instance->enemies[1] = new_character_enemy(2.0f, -2.0f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, DynamicFlee, (void *)NULL);
    instance->enemies[2] = new_character_enemy(0.0f, 2.0f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, DynamicArrive, (void *)NULL);

    return;
}

/**
 * Function that takes in a GameManager instance under creation
 * and adds to it the necessary enemy characters for the face 
 * demo. The Face demo contains
 * 
 *  - 8 characters performing face towards the player character
 */
void create_enemies_demo_face(GameManager *instance) {
    instance->enemy_count = 8;
    instance->enemies = (GameCharacter **)malloc(sizeof(GameCharacter *) * instance->enemy_count);

    instance->enemies[0] = new_character_enemy(0.0f, 3.0f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, Face, (void *)NULL);
    instance->enemies[1] = new_character_enemy(3.0f, 0.0f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, Face, (void *)NULL);
    instance->enemies[2] = new_character_enemy(0.0f, -3.0f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, Face, (void *)NULL);
    instance->enemies[3] = new_character_enemy(-3.0f, 0.0f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, Face, (void *)NULL);
    instance->enemies[4] = new_character_enemy(4.5f, -4.0f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, Face, (void *)NULL);
    instance->enemies[5] = new_character_enemy(-4.5f, -4.0f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, Face, (void *)NULL);
    instance->enemies[6] = new_character_enemy(-4.5f, 4.0f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, Face, (void *)NULL);
    instance->enemies[7] = new_character_enemy(4.5f, 4.0f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, Face, (void *)NULL);

    return;
}

/**
 * Function that takes in a GameManager instance under creation
 * and adds to it the necessary enemy characters for the first
 * Pursue and Evade demo. The first Pursue and Evade demo contains:
 *  - 1 character pursuing the player
 *  - 1 character Evading the player
 */
void create_enemies_pursue_evade_1(GameManager *instance) {
    instance->enemy_count = 2;
    instance->enemies = (GameCharacter **)malloc(sizeof(GameCharacter *) * instance->enemy_count);

    instance->enemies[0] = new_character_enemy(-2.0f, 0.0f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, Pursue, (void *)instance->player->movement);
    instance->enemies[1] = new_character_enemy(2.0f, 0.0f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, Evade, (void *)instance->player->movement);

    return;
}

/**
 * Function that takes in a GameManager instance under creation
 * and adds to it the necessary enemy characters for the second
 * Pursue and Evade demo. The second Pursue and Evade demo contains:
 *  - a first character that pursues the second
 *  - a second character that escapes from the first
 */
void create_enemies_pursue_evade_2(GameManager *instance) {
    instance->enemy_count = 2;
    instance->enemies = (GameCharacter **)malloc(sizeof(GameCharacter *) * instance->enemy_count);

    instance->enemies[0] = new_character_enemy(-4.5f, 1.0f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, Pursue, NULL);
    instance->enemies[1] = new_character_enemy(1.0f, 0.0f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, Evade, NULL);

    instance->enemies[0]->enemy_info->movement_aux = instance->enemies[1]->movement;
    instance->enemies[1]->enemy_info->movement_aux = instance->enemies[0]->movement;

    return;
}

/**
 * Fucntiont that takes in a GameManager instance under creation
 * and adds to it the necessary enemy characters for the Dynamic
 * Wandering demo. The Dynamic Wandering demo contains:
 */
void create_enemies_demo_dynamic_wandering(GameManager *instance) {
    instance->enemy_count = 8;
    instance->enemies = (GameCharacter **)malloc(sizeof(GameCharacter *) * instance->enemy_count);

    instance->enemies[0] = new_character_enemy(0.0f, 1.5f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, DynamicWander, (void *)NULL);
    instance->enemies[1] = new_character_enemy(1.5f, 0.0f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, DynamicWander, (void *)NULL);
    instance->enemies[2] = new_character_enemy(-1.5f, 0.0f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, DynamicWander, (void *)NULL);
    instance->enemies[3] = new_character_enemy(0.0f, -1.5f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, DynamicWander, (void *)NULL);
    instance->enemies[4] = new_character_enemy(3.0f, 3.0f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, DynamicWander, (void *)NULL);
    instance->enemies[5] = new_character_enemy(-3.0f, 3.0f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, DynamicWander, (void *)NULL);
    instance->enemies[6] = new_character_enemy(3.0f, -3.0f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, DynamicWander, (void *)NULL);
    instance->enemies[7] = new_character_enemy(-3.0f, -3.0f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, DynamicWander, (void *)NULL);

    return;
}

/**
 * Function that takes in a GameManager instance under creation
 * and adds to it the necessary enemy characters for the Path Following
 * demo. The Path Following demo contains:
 * 
 *  - 1 character performing path following over a simple circle path
 */
void create_enemies_demo_path_following(GameManager *instance) {
    instance->enemy_count = 1;
    instance->enemies = (GameCharacter **)malloc(sizeof(GameCharacter *) * instance->enemy_count);

    instance->enemies[0] = new_character_enemy(-4.5f, 5.0f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, PathFollowing, (void *)"circle");

    return;
}

/**
 * Function that takes in a GameManager instance under creation and 
 * adds to it the necessary enemy characters for the Separation demo.
 * The Separation demo contains:
 * 
 *  - 10 characters that perform velocity matching against the player
 *    while remaining separated between each other
 */
void create_enemies_demo_separation(GameManager *instance) {
    instance->enemy_count = 10;
    instance->enemies = (GameCharacter **)malloc(sizeof(GameCharacter *) * instance->enemy_count);

    EvasionCreateBlob aux_info = {
        .target_count = instance->enemy_count,
        .targets = instance->enemies,
        .vel_match_target = instance->player->movement
    };

    printf("EnemyCreateBlob: %d, %p, %p", aux_info.target_count, (void *)aux_info.targets, (void *)aux_info.vel_match_target);

    instance->enemies[0] = new_character_enemy(-4.5f, -0.5f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, Separation, (void *)&aux_info);
    instance->enemies[1] = new_character_enemy(-4.1f, -0.5f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, Separation, (void *)&aux_info);
    instance->enemies[2] = new_character_enemy(-3.7f, -0.5f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, Separation, (void *)&aux_info);
    instance->enemies[3] = new_character_enemy(-3.3f, -0.5f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, Separation, (void *)&aux_info);
    instance->enemies[4] = new_character_enemy(-4.3f, 0.0f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, Separation, (void *)&aux_info);
    instance->enemies[5] = new_character_enemy(-3.9f, 0.0f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, Separation, (void *)&aux_info);
    instance->enemies[6] = new_character_enemy(-3.5f, 0.0f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, Separation, (void *)&aux_info);
    instance->enemies[7] = new_character_enemy(-4.1f, 0.5f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, Separation, (void *)&aux_info);
    instance->enemies[8] = new_character_enemy(-3.7f, 0.5f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, Separation, (void *)&aux_info);
    instance->enemies[9] = new_character_enemy(-3.9f, 1.0f, 0.0f, 0.0f, 0.0f, 0.0f, 100, 0.0f, Separation, (void *)&aux_info);

    return;
}

/**
 * ##################################################################
 * ####################### the actual stuff #########################
 * ##################################################################
 */

/**
 * Function that iterates over all the enemies present in the game and
 * updates their movement and poisiton according to their current behaviour.
 * If any errors are encountered an error code of -1 is returned. Otherwise
 * the function returns 0.
 */
int game_manager_update_enemies(GameManager *instance, double time_delta) {
    if (instance == NULL) {
        return -1;
    }

    GameCharacter *player = instance->player;

    for (int i = 0; i < instance->enemy_count; i++) {
        GameCharacter *curr_enemy = instance->enemies[i];
        MovementInfo *target = NULL;
        bool erase_static = false;

        switch (curr_enemy->enemy_info->type) {
            case Seek:
            case Flee:
            case Wander:
            case Arrive:
                target = (MovementInfo *)player->movement->methods->get_static(player->movement);
                erase_static = true;
                break;
            case DynamicSeek:
            case DynamicFlee:
            case DynamicArrive:
            case DynamicAlign:
            case Face:
                target = (MovementInfo *)player->movement;
                break;
            case PathFollowing:
            case DynamicWander:
            case Pursue:
            case Evade:
            case Separation:
                target = (MovementInfo *)curr_enemy->enemy_info->movement_aux;
                break;
        };
        SteeringOutput *movement_correction = curr_enemy->enemy_info->current_behaviour(curr_enemy, target);
        if (erase_static) {
            destroy_static((Static *)target);
        }

        curr_enemy->movement->methods->update(curr_enemy->movement, movement_correction, time_delta);
        destroy_steering_output(movement_correction);
    }

    return 0;
}

/**
 * GameManager struct method that renders all characters on screen.
 * At the beginning of the method the screen is cleared and filled
 * black. So it isn't necessary to clean the display seperately
 * before rendering all game characters.
 * 
 * This will probably be changed when there are more things around.
 */
int game_manager_update_graphics(GameManager *instance) {
    SDL_Renderer *renderer = SDL_GetRenderer(instance->sdl->window);
    if (renderer == NULL) {
        return -1;
    }
    SDL_SetRenderDrawColor(renderer, 0x00, 0x00, 0x00, 0x00);
    SDL_RenderClear(renderer);

    if (instance->player->draw(instance->player, instance->sdl)) {
        printf("Something went wrong drawing player.\n");
        return -1;
    }

    for (int i = 0; i < instance->enemy_count; i++) {
        GameCharacter *curr_enemy = instance->enemies[i];
        if (curr_enemy->draw(curr_enemy, instance->sdl)) {
            printf("Something went wrong drawing enemy.\n");
            return -1;
        }
    }
    return 0;
}

/**
 * Method in charge of rendering the updated display every frame.
 * In case of any failure a status code of -1 is returned. Otherwise
 * the function returns 0.
 */
int game_manager_render(GameManager *instance) {
    if (instance == NULL) {
        return -1;
    }

    instance->sdl->render(instance->sdl);
    return 0;
}

/**
 * Game Manager struct method that processes all SDL input events
 * for the movement keys (WASD) and returns a pointer to a steering
 * output struct containing the necessary movement corrections for
 * the player. If any error is encountered a NULL pointer is returned.
 */
SteeringOutput *game_manager_handle_movement_input(GameManager *instance) {
    SDL_PumpEvents();
    // Get key state
    const unsigned char *keystate = SDL_GetKeyboardState(NULL);
    bool input_detected = false;

    //Build steering output struct for the player
    SteeringOutput *movement = (SteeringOutput *)malloc(sizeof(SteeringOutput));
    if (movement == NULL) {
        return NULL;
    }
    movement->linear = fvector_create(0.0f, 0.0f);
    if (movement->linear == NULL) {
        destroy_steering_output(movement);
        return NULL;
    }
    movement->angular = 0.0f;

    if (keystate[SDL_SCANCODE_W]) {
        input_detected = true;
        movement->linear->z += 5.0f;
    }
    if (keystate[SDL_SCANCODE_S]) {
        input_detected = true;
        movement->linear->z -= 5.0f;
    }
    if (keystate[SDL_SCANCODE_A]) {
        input_detected = true;
        movement->linear->x -= 5.0f;
    }
    if (keystate[SDL_SCANCODE_D]) {
        input_detected = true;
        movement->linear->x += 5.0f;
    }

    //If no keys are pressed down bring character to a halt
    if (!input_detected) {
        fvector_destroy(movement->linear);
        FloatVector *player_velocity = instance->player->movement->velocity;
        movement->linear = player_velocity->methods->multiply_scalar_ret(player_velocity, -5.0f);
    }

    if (input_detected) {
        instance->player->movement->orientation = new_orientation(instance->player->movement->orientation, movement->linear);
    }
    

    return movement;
}

/**
 * Function that takes in the movement corrections generated by input
 * processing and applies those changes to the player character.
 * It is assumed that the passed in SteeringOutput pointer is 
 * not NULL and points to a valid object. This function consumes
 * the inputted SteeringOutput struct and takes care of freeing it, so the
 * user shouldn't try to free it after using this function
 */
void game_manager_update_player(GameManager *instance, SteeringOutput *corrections, double time_delta) {
    instance->player->movement->methods->update(instance->player->movement, corrections, time_delta);
    destroy_steering_output(corrections);
    return;
}

/**
 * Function that creates a new heap-allocated game manager instance and
 * returns a pointer to it. As an argument it takes the name of the demo
 * that is going to be shown. Non of the struct instances owned by the 
 * game manager must be created before hand, the function takes care of
 * creating them. If any errors are encountered a NULL pointer is returned.
 */
GameManager *new_game_manager(char *demo_name) {

    GameManager *new_instance = (GameManager *)malloc(sizeof(GameManager));

    //Create player
    new_instance->player = new_character_player();

    //Create enemies according to chosen demo
    if (!strcmp(demo_name, "kinematic")) {
        create_enemies_demo_kinematic(new_instance);
    } else if (!strcmp(demo_name, "dynamic")) {
        create_enemies_demo_dynamic(new_instance);
    } else if (!strcmp(demo_name, "face")) {
        create_enemies_demo_face(new_instance);
    } else if (!strcmp(demo_name, "pursue-evade-1")) {
        create_enemies_pursue_evade_1(new_instance);
    } else if (!strcmp(demo_name, "pursue-evade-2")) {
        create_enemies_pursue_evade_2(new_instance);
    } else if (!strcmp(demo_name, "dynamic-wandering")) {
        create_enemies_demo_dynamic_wandering(new_instance);
    } else if (!strcmp(demo_name, "path-following")) {
        create_enemies_demo_path_following(new_instance);
    } else if (!strcmp(demo_name, "evasion")) {
        create_enemies_demo_separation(new_instance);
    } else {
        return NULL;
    }

    //Create SDLManager
    new_instance->sdl = create_sdl();

    //Methods
    new_instance->update_enemies = game_manager_update_enemies;
    new_instance->update_graphics = game_manager_update_graphics;
    new_instance->render = game_manager_render;
    new_instance->movement_input = game_manager_handle_movement_input;
    new_instance->update_player = game_manager_update_player;

    printf("Game Manager created.\n");
    return new_instance;
}

/**
 * Method that frees all memory allocated to a GameManager instance.
 * This effectively destroys all game objects.
 */
void destroy_game_manager(GameManager *instance) {
    if (instance == NULL) {
        return;
    }

    if (instance->player != NULL) {
        destroy_character(instance->player);
        instance->player = NULL;
    }

    for (int i = 0; i < instance->enemy_count; i++) {
        if (instance->enemies[i] != NULL) {
            destroy_character(instance->enemies[i]);
            instance->enemies[i] = NULL;
        }
    }
    free(instance->enemies);
    instance->enemies = NULL;

    if (instance->sdl != NULL) {
        instance->sdl->destroy(instance->sdl);
        instance->sdl = NULL;
    }

    free(instance);
    return;
}