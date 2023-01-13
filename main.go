package main

import "log"


func main(){
	for i := 0; i <=3; i++ {
		log.Println(i)
	}
	animals := []string {"dog", "fish", "horse", "cow", "cat"}
	for _,animal := range animals {
		log.Println(animal)
	}

	animals2 := make(map[string]string)
	animals2["dog"] = "Fido"
	animals2["cat"] = "Fluffy"
	for animalType, animal:=range animals2 {
		log.Println(animalType, animal)
	}

	/*
	* Cuando se recorre un string o se imprime un string en go
	* se imprime con el unicode del caracter al que corresponde
	* por ejeplo para O sería 79, si lo que queremos es que muestre 
	* la letra se tiene que utilizar la funcion rune que recupera
	* el caracter al que pertence el unicode  string(rune(l))
	*/
	var firstLine = "Once upon a midnight dreary"
	
	for i, l:=range firstLine {
		log.Println(i,":",  string(rune(l)))
	}


	type User struct {
		FirstName string 
		LastName string
		Email string
		Age int
	}

	var users []User 
	users = append(users, User{"John", "Smith", "john@smith.com", 30})
	users = append(users, User{"Johan", "Navarro", "johan@navarro.com", 31})
	users = append(users, User{"Caro", "Anduquia", "caro@anduquia.com", 28})
	for _, l:= range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}

	
}
