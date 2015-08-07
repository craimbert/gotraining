// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Create a package named toy with a single exported struct type named Toy. Add
// the exported fields Name and Weight. Then add two unexported fields named
// onHand and sold. Declare a factory function called New to create values of
// type toy and accept parameters for the exported fields. Then declare methods
// that return and update values for the unexported fields.
//
// Create a program that imports the toy package. Use the New function to create a
// value of type toy. Then use the methods to set the counts and display the
// field values of that toy value.
package main

// Add imports.
import (
	"fmt"
	"github.com/craimbert/gotraining/04-packaging_exporting/exercises/charles/toy"
)

// main is the entry point for the application.
func main() {
	// Use the New function from the toy package to create a value of
	// type toy.
	t := toy.New("toy1", 10)
	// Use the methods from the toy value to set some initialize
	// values.
	fmt.Printf("%#v\n", t)
	t.UpdateOnHand(5)
	fmt.Printf("%#v\n", t)
	// Display each field separately from the toy value.

	t.UpdateSold(5)
	fmt.Printf("%#v\n", t)
}
