// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/kzZZ24O23g

// Declare an untyped and typed constant and display their values.
//
// Multiply two literal constants into a typed variable and display the value.
package main

// Add imports.

// Declare a constant of kind string and assign a value.
const bill = "bill"

// Declare a constant of type integer and assign a value.
const age int16 = 25

// main is the entry point for the application.
func main() {
	// Display the value of both constants.
	println("val: ", bill)
	println(age)
	// Divide a constant of kind integer and kind floating point and
	// assign the result to a variable.
	test := age / 1.0
	// Display the value of the variable.
	println(test)
}
