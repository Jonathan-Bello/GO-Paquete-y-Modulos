# Control de dependencias

Podemos actualizarnos y movernos de distnitas versiones de dependencias

## go mod why paqueteName

Nos permite conocer el porque se importa un paquete

## go list -m -version paqueteName

Nos enlistan todas las versiones que tiene disponible ese paquete
Las dependencias indirectas se importaran dependiendo de la versi+on usada
por la dependencia directa, asi que no sera una dependencia totalmente actulizada en la ultima versión por lo general

## Actualizar dependencia en versiones menores

Recuerda comprobar con el primer numero para berificar que sean compatible

### Actualizar a la ultima versión

    go get nombrePaquete

### Actualizar a una versión especifica

    go get nombrePaquete@v1.3.1

## Actualizar dependencia versión mayor

Importación de versionamiento semantico

En go para subir de versión lo que se hace es agregar a la ruta /v#

    por ejemplo: /v2

por lo mismo cuando se crean nuevas versiones que pasan de la versión 1
En los paquetes de go crean de manera especifica una subcarpeta con la nomenclatura de /v#
con el fin de que asi se pueda acceder de manera especifica a la versión ejemplo:

        import "rsc.io/quote"   => trae la versión v1
        import "rsc.io/quote/v3"=> trea la versión v3

ahora para llamar al paquete es igual

        quote.Hello()

Nota: no podemos tener dos paquetes con el mismo nombre por lo tanto si quisieramos tener dos versiones
de un mismo paquete para utlizar debemos usar los alias

    import (
        "rsc.io/quote"
        quoteV3 "rsc.io/quote/v3" // Se le agrega un alias
    )

    quote.Hello()
    quoteV3.Hello() // se usa el alias

## Limpiar dependencias no usadas

Cuando ya no usamos una dependencia en nuestro proyecto

el comando go build, nos incluye las dependencias que nos hacen falta, pero no nos quita las que no usamos.

Para ello debemos usar:

    go mod tidy

Y nos limpia las dependencias no usadas.
Es muy importante hacer la limpieza.
