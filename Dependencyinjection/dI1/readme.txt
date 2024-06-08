Task:
You are required to create a simple calculator program in Go that uses dependency injection to support multiple operations. Specifically, the calculator should be able to perform addition and multiplication using different implementations of an operator interface.

Requirements:
Define an operator interface with a single method math that takes two integers and returns an integer.
Create two structs, addition and multiplication, that implement the operator interface.
Create a calculator struct that holds an operator.
Implement a method for the calculator struct named Calculate that uses the operator's math method to perform the calculation.
Write a constructor function NewCalculator that takes an operator and returns a new calculator with this operator.
In the main function, demonstrate the use of your calculator by performing an addition and a multiplication operation and printing the results.