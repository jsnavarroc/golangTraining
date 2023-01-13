package main

import (
	"log"

	"github.com/jsnavarroc/golangTraining/helpers"
)

const numPool = 1000 

func CalculateValue(intChan chan<- int) {	
	randomNumer := helpers.RandomNumer(numPool)
	intChan <- randomNumer	
}

func logRadomValurChan(intChan <-chan int) {
	num := <-intChan
	log.Println(num)
}

func main(){
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)
	logRadomValurChan(intChan)
	
}

