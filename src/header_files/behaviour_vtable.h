#ifndef BEHAVIOUR_VTABLE
    #define BEHAVIOUR_VTABLE

    #include "steering_output.h"
    #include "kinematic.h"
    #include "static.h"

    struct GameCharacter_s;
    typedef struct GameCharacter_s GameCharacter;
    union MovementInfo_u;
    typedef union MovementInfo_u MovementInfo;

    typedef struct BehaviourVTable_s {
        SteeringOutput *(*seek)(GameCharacter *self, MovementInfo *blob);
        SteeringOutput *(*flee)(GameCharacter *self, MovementInfo *blob);
        SteeringOutput *(*arrive)(GameCharacter *self, MovementInfo *blob);
        SteeringOutput *(*wander)(GameCharacter *self, MovementInfo *blob);
        SteeringOutput *(*dynamic_seek)(GameCharacter *self, MovementInfo *blob);
        SteeringOutput *(*dynamic_flee)(GameCharacter *self, MovementInfo *blob);
        SteeringOutput *(*dynamic_wander)(GameCharacter *self, MovementInfo *blob);
        SteeringOutput *(*dynamic_arrive)(GameCharacter *self, MovementInfo *blob);
        SteeringOutput *(*dynamic_align)(GameCharacter *self, MovementInfo *blob);
        SteeringOutput *(*pursue)(GameCharacter *self, MovementInfo *blob);
        SteeringOutput *(*evade)(GameCharacter *self, MovementInfo *blob);
        SteeringOutput *(*face)(GameCharacter *self, MovementInfo *blob);
        SteeringOutput *(*path_following)(GameCharacter *self, MovementInfo *blob);
    } BehaviourVTable;

#endif