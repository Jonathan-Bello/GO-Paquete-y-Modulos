package main

import (
	"fmt"
	"./greet"
)


func main()  {
	// podemos saber que es una funcion importada porque su primer letra es mayuscula
	fmt.Println("Hello")

	greet.English()
}