package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input program" // walrus syntax
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")

	// comma ok || comma err
	input, _ := reader.ReadString('\n') // here we dont care about errors.
	//  _ , err := reader.ReadString('\n)	// here we dont care about input.

	fmt.Println("Thanks for the rating:", input)
	fmt.Printf("Type of this rating is %T \n", input)
}
