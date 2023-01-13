package main

import "log"

type myStruct struct {
	FirstName string
}

/*
* La función printFirstName se asocia a la estructura myStruct mediante 
* la adición de una etiqueta de receptor de tipo *myStruct a la declaración 
* de la función. Esto le dice al compilador que la función printFirstName 
* se asocia a la estructura myStruct. Esto significa que podemos llamar 
* a la función printFirstName en cualquier instancia de myStruct y que la 
* función tendrá acceso a los campos de la instancia.
*/
func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main(){
	var myVar myStruct
	myVar.FirstName = "Johan"

	myVar2 := myStruct {
		FirstName: "Caro",
	}
	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.printFirstName())
}

/* 
* Este concepto de asociación de funciones a estructuras se puede 
* aplicar en muchos casos prácticos. Por ejemplo, podríamos usarlo 
* para crear una estructura que representa una tarjeta de crédito, 
* con campos como el nombre del titular, el número de tarjeta, etc. Luego, 
* podríamos asociar una función a esta estructura que verifique si el número de 
* tarjeta proporcionado es válido. De esta manera, podemos asegurarnos de que los 
* usuarios estén ingresando información correcta y segura.
*/
