package main

import "fmt"

// Define an operator interface with a single method math
type operator interface {
	math(int, int) int
}

// Define addition and multiplication structs to implement the operator interface
type addition struct{}
type multiplication struct{}

// Define a calculator struct that holds an operator
type calculator struct {
	operate operator
}

// Implement the math method for addition, which adds two integers
func (a addition) math(x, y int) int {
	return x + y
}

// Implement the math method for multiplication, which multiplies two integers
func (m multiplication) math(x, y int) int {
	return x * y
}

// NewCalculator is a constructor function that returns a new calculator with the specified operator
// This is where dependency injection happens: the operator dependency is passed into the calculator
func NewCalculator(op operator) *calculator {
	return &calculator{op}
}

// Calculate method uses the operator's math method to perform the calculation
func (c *calculator) Calculate(a, b int) int {
	return c.operate.math(a, b)
}

func main() {
	// Dependency injection in action: passing an addition operator to the calculator
	addcalci := NewCalculator(addition{})
	fmt.Println(addcalci.Calculate(5, 4)) // Output: 9

	// Dependency injection in action: passing a multiplication operator to the calculator
	mulcalci := NewCalculator(multiplication{})
	fmt.Println(mulcalci.Calculate(4, 5)) // Output: 20
}
