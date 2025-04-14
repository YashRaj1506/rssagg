package main

import "fmt"

func main () {
	fmt.Println("Hello world")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	fmt.Println("PORT:", portString)
}