package greet

// emoji variable a nivel de paquete, no puede usarse la sintaxis corta
var emoji = "🖐️"

// English retorna saludo en ingles
func English() string{
	return "Hi "+ emoji
}

// Italian retorna un saludo en italiano
func Italian() string {
	return "Ciao " + emoji
}