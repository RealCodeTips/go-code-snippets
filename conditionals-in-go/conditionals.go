// https://www.codetips.co.uk/languages/go/conditionals-in-go/

package main

import "log"

func main() {

	// Start of our code

	name := "Simon"
	age := 29
	hobbies := [4]string{"example1", "mountain biking", "example2", "example3"}

	log.Printf("Hello World. My name is %s.", name)

	if age < 13 {
		log.Println("I am considered a child")
	} else if age >= 13 && age < 20 {
		log.Println("I am considered a teenager")
	} else if age >= 20 && age < 70 {
		log.Println("I am considered an adult")
	} else {
		log.Println("I am considered a pensioner")
	}

	log.Printf("I have %d hobbies, and they are: ", len(hobbies))

	for i := 0; i < len(hobbies); i++ {
		log.Println(hobbies[i])
	}

	// End of our code

}
