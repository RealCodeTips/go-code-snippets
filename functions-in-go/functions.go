// https://www.codetips.co.uk/languages/go/functions-in-go/

package main

import "log"

func printAgeBracket(name string, age int) {
	log.Printf("%s: ", name)

	if age < 13 {
		log.Println("I am considered a child")
	} else if age < 20 {
		log.Println("I am considered a teenager")
	} else if age < 70 {
		log.Println("I am considered an adult")
	} else {
		log.Println("I am considered a pensioner")
	}
}

func main() {

	// Start of our code

	myName := "Simon"
	myAge := 29

	friendOneName := "David"
	friendOneAge := 17

	friendTwoName := "Bill"
	friendTwoAge := 42

	friendThreeName := "Charlie"
	friendThreeAge := 12

	friendFourName := "Abby"
	friendFourAge := 24

	friendFiveName := "Edith"
	friendFiveAge := 74

	log.Printf("Hello World. My name is %s.", myName)

	if myAge < 13 {
		log.Println("I am considered a child")
	} else if myAge < 20 {
		log.Println("I am considered a teenager")
	} else if myAge < 70 {
		log.Println("I am considered an adult")
	} else {
		log.Println("I am considered a pensioner")
	}

	log.Println("I have five friends, and they are as follows:")

	printAgeBracket(friendOneName, friendOneAge)
	printAgeBracket(friendTwoName, friendTwoAge)
	printAgeBracket(friendThreeName, friendThreeAge)
	printAgeBracket(friendFourName, friendFourAge)
	printAgeBracket(friendFiveName, friendFiveAge)

	// End of our code

}
