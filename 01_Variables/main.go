package main

import "fmt"

// also work a public variable when the first letter is capital
const LoginToken string = "ksbkjbjds"

func main() {
	var username string = "shivam"
	fmt.Println(username)
	fmt.Printf("Varible is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Varible is of type: %T \n", isLoggedIn)

	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("Varible is of type: %T \n", smallVal)

	var smallFloat float32 = 255.2272387237872
	fmt.Println(smallFloat)
	fmt.Printf("Varible is of type: %T \n", smallFloat)

	// default values and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var checkstring string
	fmt.Println(checkstring)
	fmt.Printf("Variable is of type: %T \n", checkstring)

	// implicit type

	var website = "techibes.in"
	fmt.Println(website)

	// no var style

	numberOfUser := 30000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Varible is of type: %T \n", LoginToken)
}
