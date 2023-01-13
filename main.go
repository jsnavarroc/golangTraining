package main

import (
	"log"
	"time"
)



type User struct {
	 firstName string
 	lastName string
 	phoneNumber string
 	age int
 	birthDate time.Time
}


func main(){
	user := User {
		firstName: "Johan",
		lastName: "Navarro",
		age: 31,
		phoneNumber: "+573166996469",
	}

	log.Println(user.firstName, user.lastName, "Birthdate:", user.birthDate)
}
