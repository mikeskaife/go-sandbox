package main

import "fmt"

const userID = 12345

func main() {

	name := "Mike"
	course := "Learn Go"

	fmt.Println("Hi", name, "you're currently watching", course)

	changeCourse(&course)

	fmt.Println("You are now watching", course)
}

func changeCourse(course *string) string {

	*course = "Docker for Numpties"

	fmt.Println("Updating your current course to", *course)

	return *course
}
