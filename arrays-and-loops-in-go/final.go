package main

import "log"

func main() {

	// Start of our code

	name := "Simon"
	age := 29
	profession := "Developer"
	favouriteAnimal := "Dog"

	hobbies := [4]string{"example1", "mountain biking", "example2", "example3"}

	log.Printf(
		"Welcome to my first Go programme. My name is %s. I am %d years old. I work as a %s. My favourite animal is a %s",
		name, age, profession, favouriteAnimal,
	)

	log.Printf("I have %d hobbies, and they are: ", len(hobbies))

	for i := 0; i < len(hobbies); i++ {
		log.Println(hobbies[i])
	}

	// End of our code

}
