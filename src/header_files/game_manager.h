#ifndef GAME_MANAGER_H
    #define GAME_MANAGER_H

    #include "game_character.h"
    #include "sdl_manager.h"
    //#include "player.h" TO-DO
    typedef struct GameManager_s {
        SDLManager *sdl;
        GameCharacter *player;
        GameCharacter **enemies;
        void (*process_input)(struct GameManager_s *self);
        SteeringOutput *(*movement_input)(struct GameManager_s *self);
        void (*update_player)(struct GameManager_s *self, SteeringOutput *corrections, double time_delta);
        int (*update_enemies)(struct GameManager_s *self, double time_delta);
        int (*update_graphics)(struct GameManager_s *self);
        int (*render)(struct GameManager_s *self);
        unsigned int enemy_count;
    } GameManager;

    GameManager *new_game_manager(char *demo_name);
    void destroy_game_manager(GameManager *instance);

#endif
