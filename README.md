# ci6450-proyecto
Repositorio para proyecto trimestral de Inteligencia Artificial para Juegos (CI6450). Universidad Simón Bolívar - Trimestre Sep-Dic 2024

## Entrega 2: Path Finding y Árboles de Decisión

Para la segunda entrega del proyecto se implementó búsqueda de caminos mediante el algoritmo A* y mecanismos de toma de decisiones para
personajes inteligentes a través de árboles de decisión.

## Contexto y personajes

La simulación está ambientada en una mina donde trabajadores extraen minerales de depósitos alrededor del mapa y los transportan a 
una caja de almacén. Los personajes específicos presentes en el juego y las lógicas que siguen son los siguientes:

* Minero (personaje marrón): Encargado de buscar depósitos de minerales alrededor del mapa y depositarlos en el carrito de mina (caja gris). Tiene un medidor
de stamina que causa que se desmaye y quede inmóvil cuando llega a cero. El stamina del personaje solo disminuye cuando está buscando o depositando minerales.
Su árbol de decisión es el siguiente:
    * Si no tiene estámina. Detenerse
    * En caso contrario
        * Si estoy cargando minerales, llevarlos al carrito
        * En caso contrario
            * Si hay minerales presentes en el mapa, ir a minarlos.
* Colector (personaje naranja): Encargado de buscar los minerales depositados en el carrito por el minero y llevarlos a la caja de almacén. Si no hay minerales
en el carrito empieza a deambular por el mapa (usando Dynamic Wandering). Su árbol de decisión es el siguiente:
    * Si estoy cargando un mineral, llevarlo a la caja de almacén.
    * En caso contrario
        * Si hay minerales en el carrito, ir a buscarlos
        * En caso contrario, deambular
* Médico (personaje cyan): Se encarga de llevar agua a los mineros desmayados y restablecer su stamina. Tiene un inventario interno de botellas de agua de las
que dispone para atender mineros (cada vez que atiende un minero el inventario decrece en 1). Si en algún momento el minero está desmayado y no tiene botellas de
agua para atenderlo, va a buscar más botellas de agua al almacén de agua (caja celeste). Su árbol de decisión es el siguiente:
    * Si el minero no está desmayado, regresar a la enfermería.
    * En caso contrario
        * Si se tiene agua para atenderlo, ir a revivir al minero.
        * En caso contrario, ir a buscar agua al almacén.

## Entorno utilizado y cómo ejecutar

El proyecto está escrito en Golang (>=1.22.7) haciendo uso de los bindings a SDL2 obtenidos del repo `github.com/veandco/go-sdl2`. El proyecto fue desarrollado y
probado en Fedora Linux versión 40/41 haciendo uso del sistema de gráficos Wayland. Para poder compilar los bindings de los que depende el proyecto dicho sistema debe
tener instalados los siguientes paquetes (O sus equivalentes en otras distribuciones):

* SDL2-devel (32 y 64 bits)
* SDL2_image-devel (32 y 64 bits)
* SDL2_gfx-devel (32 y 64 bits)
* Grupo de paquetes "Development Tools" (recomendado)
* Grupo de paquetes "C Development Tools and Libraries" (recomendado)

Para compilar el proyecto a un ejecutable se puede ejecutar el siguiente comando:
```
go build -o <ubicacion_binario> main.go
```

También puede ejecutarse directamente mediante el comando
```
go run main.go
```