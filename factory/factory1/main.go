// lets implement a logger system ,so we have some methods like info() and error() .using factory method we can switch to different implementation like zaplogge or logrus
// without change in existing logger system.

package main

import "fmt"

// Define the 'logrus' struct, representing a logger implementation.
type logrus struct{}

// Define the 'zaplogger' struct, representing another logger implementation.
type zaplogger struct{}

// Define the logger interface with 'info()' and 'error()' methods.
type logger interface {
	info()
	error()
}

// Implement the 'info()' method for the 'logrus' struct.
func (l logrus) info() {
	fmt.Println("logrus info")
}

// Implement the 'info()' method for the 'zaplogger' struct.
func (z zaplogger) info() {
	fmt.Println("zaplogger info")
}

// Implement the 'error()' method for the 'logrus' struct.
func (l logrus) error() {
	fmt.Println("logrus error")
}

// Implement the 'error()' method for the 'zaplogger' struct.
func (z zaplogger) error() {
	fmt.Println("zaplogger error")
}

// Factory utility function to create logger instances based on input string.
// Returns a logger interface.
func factoryUtilityFunc(l string) logger {
	if l == "logrus" {
		return logrus{}
	} else {
		return zaplogger{}
	}
}

func main() {
	// Create an instance of 'logrus' logger using its constructor.
	logrus := logrus{}
	// Call the 'info()' method of the 'logrus' logger instance.
	logrus.info()
	// Call the 'error()' method of the 'logrus' logger instance.
	logrus.error()

	// Create an instance of 'zaplogger' logger using its constructor.
	zaplog := zaplogger{}
	// Call the 'error()' method of the 'zaplogger' logger instance.
	zaplog.error()
	// Call the 'info()' method of the 'zaplogger' logger instance.
	zaplog.info()
}
