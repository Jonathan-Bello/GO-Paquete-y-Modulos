# Paquetes

Son carptas que contien una coleción de archivos y nos provee una funcionalidad

## Creación de nuestro propio paquete greet

Creamos una carpeta con dos archivos, en la primer linea debemos de especificar el apquete al que pertenencen

Cuando definimos un paquete, ese puede compartir variables entre archivos solo
por pertenecer al mismo paquete

Siempre debes poner nombres muy muy especificos en nuestros paquetes, pasa saber que hacen
de manera rapida. Si ponemos nombres genericos como "utils" es una mala practica.

## Private y publico

Para indicar a Go que queremos que se pueda acceder a elementos de nuestro paquete podemos
indicar si queremos que un elemento sea publico o privado.

Esto se realiza simmplemente indicando la primer letra de nuestro elemento.

    var emoji => variable privada
    var Emoji => variable publica

Y esto de igual manera con nuestras funciones o estructuras.
