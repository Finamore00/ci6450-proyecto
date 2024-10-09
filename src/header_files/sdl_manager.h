#ifndef SDL_MANAGER_H
#define SDL_MANAGER_H
    #include <SDL2/SDL.h>

    typedef struct SDLManager {
        SDL_Window *window;
        SDL_Surface *surface;
        SDL_Renderer *renderer;
        int (*init)(struct SDLManager *self, int width, int height);
        void (*render)(struct SDLManager *self);
        void (*destroy)(struct SDLManager *self);
    } SDLManager;

    SDLManager *create_sdl(void);
#endif