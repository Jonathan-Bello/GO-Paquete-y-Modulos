# Modulos

Los modulos de go nos permiten administrar las dependencias de nuestros paquetes, y controlar
las versiones de los mismos.

Debemos tener un solo modulo en la raiz de nuestro proyecto.

## Crear modulo

Nos vamos a la raiz de nuestro proyecto y usamos el siuiente comando

    go mod init nombreModulo

El nombre del modulo: Este debe ser la misma ruta en donde esta alojado nuestro cosdigo fuente, es decir
la ruta a nuestra repositorio remoto.

Esto es porque nuestro repositorio es publico, asi que si otro desarrollador quisiera usar nuestro paquete, solo
tendria quue importarlo atravez de esa ruta.

    go mod init github.com/Jonathan-Bello/GO-Paquete-y-Modulos

## Uso del modulo

Se nos crea un archivo llamado go.mod
Aqui es donde se llevan acabo el control de dependencias.

Y para importar nuestro modulo, en nuestro archivo main colocamos, la ruta a githib/nombreModulo

    import (
        "github.com/Jonathan-Bello/GO-Paquete-y-Modulos/greet"
    )
