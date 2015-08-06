// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/ItPe2EEy9X

// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

// Add imports.
import "fmt"

// Add user type and provide comment.
type user struct {
	name  string
	email string
	age   int
}

// main is the entry point for the application.
func main() {
	// Declare variable of type user and init using a struct literal.
	bill := user{
		name:  "Bill",
		email: "bill@ardanstudios.com",
		age:   25,
	}
	// Display the field values.
	fmt.Printf("%#v\n", bill)
	// Declare a variable using an anonymous struct.

	joe := struct {
		name  string
		email string
		age   int
	}{
		name:  "Joe",
		email: "joe@ardanstudios.com",
		age:   29,
	}
	// Display the field values.
	fmt.Printf("%#v\n", joe)
}
