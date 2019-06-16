// https://www.codetips.co.uk/languages/go/conditionals-in-go/

package main

import "log"

func ageToAgeBracket(age int) string {
	return ""
}

func main() {

	// Start of our code

	myName := "Simon"
	myAge := 29

	friendOneName := "David"
	friendOneAge := 17

	friendTwoName := "Bill"
	friendTwoAge := 42

	// friendThreeName := "Charlie"
	// friendThreeAge := 12

	// friendFourName := "Abby"
	// friendFourAge := 24

	// friendFiveName := "Edith"
	// friendFiveAge := 74
	
	log.Printf("Hello World. My name is %s.", myName)

	if myAge < 13 {
		log.Println("I am considered a child")
	} else if myAge >= 13 && myAge < 20 {
		log.Println("I am considered a teenager")
	} else if myAge >= 20 && myAge < 70 {
		log.Println("I am considered an adult")
	} else {
		log.Println("I am considered a pensioner")
	}

	log.Println("I have two friends, and they are as follows:")

	log.Printf("%s: ", friendOneName)

	if friendOneAge < 13 {
		log.Println("They are considered a child")
	} else if friendOneAge >= 13 && friendOneAge < 20 {
		log.Println("They are considered a teenager")
	} else if friendOneAge >= 20 && friendOneAge < 70 {
		log.Println("They are considered an adult")
	} else {
		log.Println("They are considered a pensioner")
	}

	log.Printf("%s: ", friendTwoName)

	if friendTwoAge < 13 {
		log.Println("They are considered a child")
	} else if friendTwoAge >= 13 && friendTwoAge < 20 {
		log.Println("They are considered a teenager")
	} else if friendTwoAge >= 20 && friendTwoAge < 70 {
		log.Println("They are considered an adult")
	} else {
		log.Println("They are considered a pensioner")
	}



	// End of our code

}
