// https://www.codetips.co.uk/languages/go/functions-in-go-challenge-answer/

package main

import "log"

func printAgeBracket(name string, age int, isFriend bool) {
	if isFriend {
		log.Printf("%s: ", name)
	} else {
		log.Printf("Hello World. My name is %s.", name)
	}

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

	printAgeBracket(myName, myAge, false)

	log.Println("I have five friends, and they are as follows:")

	printAgeBracket(friendOneName, friendOneAge, true)
	printAgeBracket(friendTwoName, friendTwoAge, true)
	printAgeBracket(friendThreeName, friendThreeAge, true)
	printAgeBracket(friendFourName, friendFourAge, true)
	printAgeBracket(friendFiveName, friendFiveAge, true)

	// End of our code

}
