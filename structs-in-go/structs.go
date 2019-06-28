// https://www.codetips.co.uk/languages/go/structs-in-go/

package main

import "log"

type person struct {
	name string
	age  int
}

func printAgeBracket(p person, isFriend bool) {
	if isFriend {
		log.Printf("%s: ", p.name)
	} else {
		log.Printf("Hello World. My name is %s.", p.name)
	}

	if p.age < 13 {
		log.Println("I am considered a child")
	} else if p.age < 20 {
		log.Println("I am considered a teenager")
	} else if p.age < 70 {
		log.Println("I am considered an adult")
	} else {
		log.Println("I am considered a pensioner")
	}
}

func main() {

	// Start of our code

	me := person{"Simon", 29}

	friends := [5]person{
		{"David", 17},
		{"Bill", 42},
		{"Charlie", 12},
		{"Abby", 24},
		{"Edit", 74},
	}

	printAgeBracket(me, false)

	log.Printf("I have %d friends, and they are as follows:", len(friends))

	for i := 0; i < len(friends); i++ {
		printAgeBracket(friends[i], true)
	}

	// End of our code

}
