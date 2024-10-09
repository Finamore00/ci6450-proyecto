#ifndef GAME_CHARACTER_H
    #define GAME_CHARACTER_H

    #include "static.h"
    #include "kinematic.h"
    #include "behaviour_vtable.h"
    #include "sdl_manager.h"
    #include <stdbool.h>

    typedef enum CharacterType_e {
        Player_t,
        Enemy_t,
        OtherCh_t
    } CharacterType;

    typedef enum EnemyBehaviour_e {
        Seek,
        Flee,
        Arrive,
        Wander,
        DynamicSeek,
        DynamicFlee,
        DynamicArrive,
        DynamicWander,
        DynamicAlign,
        Pursue,
        Evade,
        Face,
        PathFollowing,
        Separation
    } EnemyBehaviour;

    typedef union MovementInfo_u {
        Kinematic k;
        Static s;
    } MovementInfo;

    // Enemy character type
    typedef struct Enemy_s {
        EnemyBehaviour type;
        SteeringOutput *(*current_behaviour)(GameCharacter *self, MovementInfo *target);
        void *movement_aux;
        const BehaviourVTable *behaviours;
        float arrive_radius;
        unsigned short int health;
        bool active;
    } Enemy;
    Enemy *new_enemy(unsigned int health, float radius, EnemyBehaviour behaviour, void *misc);
    void destroy_enemy(Enemy *instance);
    int enemy_draw(GameCharacter *instance, SDLManager *sdl);

    //Player character type
    typedef struct Player_s {
        unsigned int health;
    } Player;
    Player *new_player(unsigned int health);
    void destroy_player(Player *player);
    int player_draw(GameCharacter *instance, SDLManager *sdl);

    typedef struct GameCharacter_s {
        CharacterType type;
        Kinematic *movement;
        Player *player_info;
        Enemy *enemy_info;
        int (*draw)(GameCharacter *self, SDLManager *sdl);
    } GameCharacter;

    GameCharacter *new_character_enemy(float xo, float zo, float vxo, float vzo, float orientation, float angular, unsigned int health, float radius, EnemyBehaviour behaviour, void *misc);
    GameCharacter *new_character_player(void); //TO-DO
    void destroy_character(GameCharacter *instance);

#endif