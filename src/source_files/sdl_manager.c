#include "../header_files/sdl_manager.h"
#include <SDL2/SDL.h>
#include <SDL2/SDL_image.h>
#include <stdio.h>

/**
 * Function that inits the SDL system of an SDLManager instance. It is
 * meant to be called similar to an object method and the SDLManager
 * pointer passed to it is that of the instance itself. The width and
 * height variables specify the dimmensions of the desired window.
 * 
 * Returns -1 on initialization failure and 0 on success.
 */
int sdl_init(SDLManager *instance, int width, int height) {
    if (SDL_Init(SDL_INIT_VIDEO) < 0) { 
        printf("Failed to init SDL system.\n");
        printf("Error string: %s\n", SDL_GetError());
        return -1;
    }

    instance->window = SDL_CreateWindow("Window Title", SDL_WINDOWPOS_CENTERED, SDL_WINDOWPOS_CENTERED, width, height, SDL_WINDOW_SHOWN);
    if (instance->window == NULL) {
        printf("Failed to instantiate window.\n");
        printf("Error string: %s\n", SDL_GetError());
        return -1;
    }

    instance->surface = SDL_GetWindowSurface(instance->window);
    if (instance->surface == NULL) {
        printf("Failed to get window surface instance.\n");
        printf("Error string: %s\n", SDL_GetError());
        return -1;
    }

    instance->renderer = SDL_GetRenderer(instance->window);
    if (instance->renderer == NULL) {
        printf("Failed to get window renderer.\n");
        printf("Error string: %s\n", SDL_GetError());
        return -1;
    }

    return 0;
}

/**
 * Struct method for SDLManager tasked with rendering the window surface.
 * Right now it only displays a white rectangle
 */
void sdl_render(SDLManager *instance) {
    SDL_RenderPresent(instance->renderer);
    return;
}

/**
 * Struct method that lets an SDLManager instance clean after itself. Destroying
 * all relevant objects and quitting all opened SDL subsystems. It is still
 * the caller's responsability to free the SDLManager object. 
 */
void sdl_destroy(SDLManager *instance) {
    SDL_FreeSurface(instance->surface);
    instance->surface = NULL;

    SDL_DestroyWindow(instance->window);
    instance->window = NULL;
        
    SDL_Quit();
    return;
}

/**
 * Function that creates a new instance of an SDLManager
 * struct and returns a pointer to it. It is the caller's
 * responsability to subsequently init the SDL system and 
 * ultimately free the pointer.
 *
 * Definition of SDLManager:
 * typedef struct {
 *     SDL_Window *window;
 *     SDL_Surface *surface;
 *     void (*init)(SDLManager *self, int width, int height);
 * }
 * 
 */
SDLManager *create_sdl(void) {
    SDLManager* instance = (SDLManager*)malloc(sizeof(SDLManager));
    if (instance == NULL) {
        return NULL;
    }
    
    instance->window = NULL;
    instance->surface = NULL;
    instance->init = sdl_init;
    instance->render = sdl_render;
    instance->destroy = sdl_destroy;

    return instance;
}