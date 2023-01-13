package main

import "log"



func main(){
	var isTrue bool
	isTrue = false
	if isTrue {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}
	//---------------------------------------
	cat := "cat"
	if(cat == "cat") {
		log.Println("Cat is cat")
	} else {
		log.Println("Cat is not cat")
	}
	//---------------------------------------

	myNum := 100 
	if myNum > 99 && !isTrue {
		log.Println("myNum is generate than 99 and is True is set to true")
	} else if myNum < 100 && isTrue {
			
			log.Println("1")
	} else if myNum == 101 || isTrue {
		log.Println("2")

	} else if myNum > 1000 && isTrue == false {
		log.Println("3")
	}
	//---------------------------------------

	myVar := "cat"

	switch myVar {
		case "cat":
			log.Println("cat is set to Cat")
		case "dog":
			log.Println("cat is set to Dog")
		case "fish":
			log.Println("cat is set to Fish")
		default: 
			log.Println("cat is somthing else")
	}
}
