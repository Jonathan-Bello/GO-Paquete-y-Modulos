# Importar desde repositorio de terceros

Para realizarlo ponemos la ruta de importación del paquete
en import

Ahora para que go obtenga la dependencia de terceros usamos el comando

    go build

## go build

Este comando nos permite compilar nuestro proyecto y crear ya sea un archivo binario o ejecutable.

Pero adiconalmente verifica que dependencias hacen falta y las descarga.
Y leera el archvo go.mod y descargara las despendencias

## go env

Este comando nos muestra la lista de toda las dependencias que usa go para trabajar

Aqui podemos encontrar el GOPATH = C:\Users\fenix\go

Ahi es donde se clonan nuestras dependencias

## GOPATH

Si entramos al directorio de la carpeta de nuestro GOPATH
Podemos ir a la carpeta pkg (package), y dentro de ella se encuentra la carpeta mod
y aqui podemos ver los respositorios de dependencias que clonamos.

## Aciones del comando build

Adicionalemnete al ejecutar go build
Vemos como en  go.mod nos muestra nuestros paquetes requeridos

Y crea un archivo llamadao go.sum que lleva el control de las dependencias directas
e indirectas con las que estamos trabajando

Las dependencias indirectas son las que ocupan las dependencias que usamos de manera directa.

## Ejecutar build

Para ejecutar el archivo ejecutable generado por build en terminal colocamos

    ./nombreEjecutable.extension
