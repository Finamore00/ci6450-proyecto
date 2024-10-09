#ifndef EVASION_H
    #define EVASION_H

    #include "kinematic.h"
    #include "game_character.h"
    typedef struct Evasion_s {
        Kinematic **targets;
        Kinematic *vel_match_target;
        unsigned int target_count;
        float thresshold;
        float decay_coefficient;
    } Evasion;

    typedef struct EvasionCreateBlob_s {
        Kinematic **targets;
        Kinematic *vel_match_target;
        unsigned int target_count;
    } EvasionCreateBlob;

    Evasion *new_evasion(GameCharacter **members, unsigned int member_count);
    void destroy_evasion(Evasion *);

#endif