package main

import "log"

/*
* & es un operador de dirección de memoria, lo que significa que devuelve
* la dirección de memoria de una variable,
* mientras que * es un operador de desreferenciación, lo que significa
* que devuelve el valor almacenado en una
* dirección de memoria.
 */
func main(){
	var myString string
	myString = "Green"
	log.Println("myString is set to", myString)
	changeUsingPointer(&myString)
	log.Println("after func call myString is set to", myString)
}


func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
}