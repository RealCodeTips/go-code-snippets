//

package main

import "log"

type person struct {
	name     string
	age      int
	isFriend bool
}

func printAgeBracket(p person) {
	if p.isFriend {
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

	me := person{"Simon", 29, false}

	friends := [5]person{
		{"David", 17, true},
		{"Bill", 42, true},
		{"Charlie", 12, true},
		{"Abby", 24, true},
		{"Edit", 74, true},
	}

	printAgeBracket(me)

	log.Printf("I have %d friends, and they are as follows:", len(friends))

	for i := 0; i < len(friends); i++ {
		printAgeBracket(friends[i])
	}

	// End of our code

}
