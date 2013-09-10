package main

import (
	"fmt"
	"time"
	"strconv"
)

func makeCakeAndSend(cs chan string, jobs int) {
	for i := 1; i<=jobs; i++ {
		cakeName := "Strawberry cake " + strconv.Itoa(i)
		fmt.Println("Making a cake and sending ...", cakeName)
		cs <- cakeName
	}
}

func receiveCakeAndPack(cs chan string, jobs int) {
	for i := 1; i <= jobs; i++ {
		s := <- cs
		fmt.Println("Packing received cake: ", s)
	}
}

func main() {
	cs := make(chan string)
	jobs := 10

	go makeCakeAndSend(cs, jobs)
	go receiveCakeAndPack(cs, jobs)

	//sleep until the code completes - there must be a better way!
	time.Sleep(4 * 1e9)
}
