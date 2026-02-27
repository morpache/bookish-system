package main

import (
	"fmt"
	"log"
)

func main() {
	var username string
	var age int
	println("Hey! Print ur name pls :3")
	_, err := fmt.Scan(&username)
	if err != nil {
		log.Fatal(err)
	}
	println("Oh dear " + username + ", how old are you?")
	_, err = fmt.Scan(&age)
	if err != nil {
		log.Fatal(err)
	}
	if age < 30 {
		println("OMG who's this twinkie-tinkie")
	} else {
		println("Omg men, u r dilf :-0")
	}
	println("is vibecoding a good thing?")
}
