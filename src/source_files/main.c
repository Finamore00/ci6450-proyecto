#include "../header_files/game_manager.h"
#include <SDL2/SDL.h>
#include <stdio.h>
#include <stdbool.h>
#include <time.h>
#include <unistd.h>

//Global variables
const int WINDOW_WIDTH = 640;
const int WINDOW_HEIGHT = 480;

void print_wrong_arugment_count(void) {
    printf("Número errado de argumentos. El uso del programa viene dado por:\n");
    printf("\ttarget/Game <NOMBRE_DEMO>\n");
    printf("Nombres de demos disponibles:\n");
    printf("\t* kinematic\n");
    printf("\t* dynamic\n");
    printf("\t* face\n");
    printf("\t* pursue-evade-1\n");
    printf("\t *pursue-evade-2\n");
    printf("\t* dynamic-wandering\n");
    printf("\t* path-following\n");
    printf("\t* evasion\n");
    return;
}

void print_wrong_demo_name(void) {
    printf("Nombre de demo inválido. El uso del programa viene dado por:\n");
    printf("\ttarget/Game <NOMBRE_DEMO>\n");
    printf("Nombres de demos disponibles:\n");
    printf("\t* kinematic\n");
    printf("\t* dynamic\n");
    printf("\t* face\n");
    printf("\t* pursue-evade-1\n");
    printf("\t* pursue-evade-2\n");
    printf("\t* dynamic-wandering\n");
    printf("\t* path-following\n");
    printf("\t* evasion\n");
    return;
}

char *verify_demo_name(char *demo_name) {
    char *demo_names[8] = {
        "kinematic", "dynamic", "face",
        "pursue-evade-1", "pursue-evade-2",
        "dynamic-wandering", "path-following",
        "evasion"
    };

    for (int i = 0; i < 8; i++) {
        if (!strcmp(demo_name, demo_names[i])) {
            return demo_name;
        }
    }

    return NULL;
}

int main(int argC, char *argV[]) {

    //Validate input
    if (argC != 2) {
        print_wrong_arugment_count();
        return -1;
    }

    char *demo_name = verify_demo_name(argV[1]);
    if (demo_name == NULL) {
        print_wrong_demo_name();
        return -1;
    }

    GameManager *game = new_game_manager(demo_name);

    game->sdl->init(game->sdl, WINDOW_WIDTH, WINDOW_HEIGHT);
    SDL_Event e;
    bool quit = false;

    double time_delta = 1.0l / 60.0l;

    while (!quit) {

        while (SDL_PollEvent(&e) != 0) {
            if (e.type == SDL_QUIT) {
                quit = true;
                continue;
            }
        }

        time_t frame_start_t = time(NULL);

        game->update_player(game, game->movement_input(game), time_delta);
        game->update_enemies(game, time_delta);
        game->update_graphics(game);
        game->render(game);

        time_t frame_end_t = time(NULL);

        double time_delta = difftime(frame_start_t, frame_end_t);
        if (time_delta < 16.66l) {
            double sleep_time = 16.66l - time_delta;
            usleep(sleep_time*1000);
        }
    }

    destroy_game_manager(game);
    return 0;
}