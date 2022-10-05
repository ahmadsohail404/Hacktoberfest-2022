package main

import "fmt"
const LoginToken string = "xaydfgenjf" //  starting uppercase letter means it's public

func main() {
	var username string = "Sohail"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n\n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n\n", isLoggedIn)

	var smallVal uint64 = 12345
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n\n", smallVal)

	var smallFloat float64 = 255.12355657567576576766
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n\n", smallFloat)

	// Default values and some aliases.
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n\n", anotherVariable)

	// Implicit type
	var website = "sohailahmad.netlify.app"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n\n", website)

	// no var style
	numberOfUsers := 30000.01 // -> this syntax can only be used inside a method and not global.
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable is of type: %T \n\n", numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n\n", LoginToken)

}
