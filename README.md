# ci6450-proyecto
Repositorio para proyecto trimestral de Inteligencia Artificial para Juegos (CI6450). Universidad Simón Bolívar - Trimestre Sep-Dic 2024

## Entrega 1: Steering Behaviours

Para esta primera entrega se realizó la implementación de varios algoritmos de Steering Behaviours y de un motor de juego básico que permite hacer uso de los mismos dentro un *game loop* funcional. Más específicamente, se implementaron los siguientes SteeringBehaviours siguiendo la bibliografía:

* Kinematic Seek
* Kinematic Flee
* Kinematic Arrive
* Kinematic Wandering
* Dynamic Seek
* Dynamic Flee
* Dynamic Arrive
* Velocity Matching
* Dynamic Wandering
* Align
* Face (haciendo uso de Align)
* Pursue + Look Where You're Going
* Evade + Look Where You're Going
* Path Following 
* Separation (Usando Velocity Matching)

## Entorno utilizado y cómo ejecutar

El proyecto se desarrolló y probó en Fedora Linux versión 40 bajo el entorno de escritorio KDE con el subsistema de gráficos Wayland. Para poder compilar exitosamente el proyecto y ejecutarlo se recomienda tener instalados los siguientes paquetes y librerías:

* Grupos de paquetes de DNF "Development Tools" y "Development Libraries"
* Librería SDL2 Y sus headers de desarrollo para 32 y 64 bits (SLD2-devel)
* Librería SDL2_gfx y sus headers de desarrollos para 32 y 64 bits (SDL2_gfx-devel)

## ¿Por qué hay tan pocos commits en el repo?

Al iniciar con la decisión de implementar mi propio motor no estaba del todo comprometido con la idea. Por este motivo no quería comprometerme del todo subiendo de una vez al repositorio de GitHub. Para el momento en que ya era demasiado tarde para cambiar de opinión, estaba ocupado implementando esta monstruosidad, así que subirlo al repo se volvió un *afterthought* hasta último minuto. De ahora en adelante si estaré haciendo adiciones al repositorio de manera regular.