# Versionamiento

## go list -m all

Este comando nos permite enlista nuestro modulo.

En la primer linea el nombre de nuiestro modulo

Y en las siguientes nuestras dependencias y su versión

si go va a descargar una dependencia y no encuentra una etiqueta de versión para
descargar la ultima versión, descarga el ultimo codigo fuente y crea una seudoversión+

Seudoversión => v0.0.0-20170915032832-14c0d48ead0c

version seguida de fecha y hora del comit seguido de su hash

## Versionamiento semantico (SemVer)

Se comforma de la letra V seguida de 3 numeros separados por puntos

V 2 . 4 . 10 -alpha

el primer numero(2)  -> versión mayor (major)

el segundo numero(4) -> versión menor (minor)

el tercer numero(10)  -> versión patch

Opcionalmente podemos agregar al final un guion y un identificador para indicar que es una versión de prelanzamiento.

    -alpha

### Para que sirven las versiones

Cuando construimos nuestro paquete o proyecto, cualquier cambio que nosotros hagamos debe de versionarse

v0.0.0 => el paquete no deberia ser cosiderado estable. En desarrollo
v1.0.0 => define la versión del paquete publico. Inicial

Al inicio de un proyecto siempre vamos a usar como primer numero 0.
Y unicamente vamos a actualizar los codigos de la version menor y el patch, pero siempre dejaremos
al inicio el 0, para indicar que aun no es estable.

Cuando ya decimos lanzarlo al publico ya podemos cambiarle al primer valor al numero 1.

### PATCH

Corrige bugs compatibles con la versión anterior

Patch solo se actualiza cuando es cosa de correción de bugs

v1.0.0

v1.0.1  Solucionamos un error, pero sigue siendo compatible con la versión 1.0.0

v1.0.2  Solucionamos un error, pero sigue siendo compatible con la versión 1.0.0 y 1.0.1

### MINOR

Cambios y adiciones que no rompen la compatibilidad de la versiones anteriores del paquete dentro
de la versión major

Agregamos nuevas caracteristicas pero con compatibilidad a la versioness anteriores

v1.0.2

v1.1.0   OJO, cuando cambiamos la versión minor el numero de patch debe reiniciarse a 0

v1.2.0

### MAJOR

Cambios que rompen la compatibilidad con las versiones anteriores del paquete
Cuando aumentamos la version major, debemos de reiniciar el numero de minor y patch

v1.2.0

v2.0.0

v3.0.0
