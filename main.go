package main

import (
	"log"
	"mygo/helpers")


const numPool =1000
func calculateValue(myChan chan int) {

	rendomNumber := helpers.RendomNumber(numPool)
	myChan <- rendomNumber

}

func main() {
	myChan := make(chan int)

	defer close(myChan)

	go calculateValue(myChan)

	num := <-myChan
	log.Println(num)


}
