// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/b6-FNFOToO

// Declare a struct type and create a value of this type. Declare a function
// that can change the value of some field in this struct type. Display the
// value before and after the call to your function.
package main

// Add imports.
import "fmt"

// Declare a type named user.
type user struct {
	name  string
	email string
}

// Create a function that changes the value of one of the user fields.
func updateEmail(oldValue *string, newValue string) {
	// Use the pointer to change the value that the
	// pointer points to.
	*oldValue = newValue
}

// main is the entry point for the application.
func main() {
	// Create a variable of type user and initialize each field.
	bill := user{
		name:  "Bill",
		email: "bill@ardanstudios.com",
	}
	fmt.Printf("%#v\n", bill)
	// Display the value of the variable.
	println(bill.email)
	// Share the variable with the function you declared above.
	updateEmail(&bill.email, "charles@gm.com")
	// Display the value of the variable.
	fmt.Printf("%#v\n", bill)
	println(bill.email)
}
