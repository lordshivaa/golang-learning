package main

import (
	"bufio"
	"fmt"
	"os"
)

// bufio

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")
	// fmt.Printf("Type of reader: %T \n", reader)
	// comma ok || error ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T ", input)
}
