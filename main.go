// Created by: Infinity de Guzman
// Created on: May 2021
//
// This program registers your address

package main

import "fmt"

func main() {
	// This function prints the address
	var length int
	var width int

	// input
	fmt.Println("This program gets a user's address.")
	fmt.Println()
	fmt.Print("Enter the length: ")
	fmt.Scanln(&length)
	fmt.Print("Enter the width: ")
	fmt.Scanln(&width)

	// output
	fmt.Println("The area is: ", (length * width), "cm²")
	fmt.Println("The perimeter is: ", (2 * (length + width)), "cm")
	fmt.Println("\nDone.")
}
