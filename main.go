package main

import (
	"fmt"

	// Forma importarcion local
	// "./greet"
	// Forma importación desde repositorio
	"github.com/Jonathan-Bello/GO-Paquete-y-Modulos/greet"

	// Importar paquete de terceros
	"rsc.io/quote"

	// Importación de una versión mayor
	quoteV3 "rsc.io/quote/v3" // Se le agrega un alias
)


func main()  {
	// podemos saber que es una funcion importada porque su primer letra es mayuscula
	fmt.Println(greet.English())
	fmt.Println(greet.Spanish())	
	fmt.Println(greet.Italian())
		
	fmt.Println(quote.Hello())
	fmt.Println(quoteV3.Concurrency())
	
	// Para ejecutar el archivo ejecutable generado por build en terminal colocamos
	// ./nombreEjecutable.extension
}