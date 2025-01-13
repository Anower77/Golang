package main

import "fmt"

// Print welcome message
func welcomeMassage() {
	fmt.Println("welcome to function!!")
}

// Get user name as input
func getUserName() string {
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)
	return name
}

// calculateBetweenTwoNumber
func calculateBetweenTwoNumber() (int, int) {
	var num1, num2 int
	fmt.Println("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter second number: ")
	fmt.Scanln(&num2)
	return num1, num2
}

// sum
func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

// Display results
func display(name string, sum int) {
	fmt.Println("Hello - ", name)
	fmt.Println("Sum - ", sum)
}

// Goobye massage
func goodbye() {
	fmt.Println("Thank you for using this application!!")
}

// Main function
func main() {
	// Call Function
	welcomeMassage()
	name := getUserName()
	num1, num2 := calculateBetweenTwoNumber()
	sum := add(num1, num2)
	display(name, sum)
}
