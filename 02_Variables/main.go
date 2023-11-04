package main

import "fmt"

//Starts with a capital letter means that it's public
const LoginToken string ="TRHBJ788BHJ"


func main() {
	var username string = "Peluca"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)

	fmt.Println()

	var isLogged bool = true
	fmt.Println(isLogged)
	fmt.Printf("Variable is of type: %T\n", isLogged)

	fmt.Println()

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T\n", smallVal)

	fmt.Println()

	var smallFloat float64 = 255.4555135765135746543
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T\n", smallFloat)

	fmt.Println()

	// default values and some alisaes

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T\n", anotherVariable)

	fmt.Println()

	// implicit type

	var website = "google.es"
	//website = 3 ## ERROR 
	fmt.Println(website)

	fmt.Println()

	// no var style

	numberOfUser := 1500 // This only works inside a method
	fmt.Println(numberOfUser)

	fmt.Println()

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T\n", LoginToken)
}
