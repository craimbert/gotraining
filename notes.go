/*
MULTI LINES COMMENT
MULTI LINES COMMENT
*/

////////////////////////////////////////////////////////////////////////////////
// Imports
////////////////////////////////////////////////////////////////////////////////
import (
	_ "temporary unused package" // use that to declare a package you want to use later
	"fmt"
)

// PRINT stuff => 'println(stuff)'  WITHOUT any import

////////////////////////////////////////////////////////////////////////////////
// Variables
////////////////////////////////////////////////////////////////////////////////

// Declare variables that are set to their zero value.
var age int

// Display the value of those variables.
fmt.Println(age)

// Declare variables and initialize.
// Using the short variable declaration operator.
month := 10

var bool1, bool2 bool // declare N var of same type with just ","
// -> default bool value = false


////////////////////////////////////////////////////////////////////////////////
// Structs
////////////////////////////////////////////////////////////////////////////////

// example represents a type with different fields.   // NO COMMA !!!
type example struct {
   flag bool // NO COMMA !!!
   counter int16 // NO COMMA !!!
   pi float32 // NO COMMA !!!
}

// Declare a variable of type example set to its
// zero value.
var e1 example

// init using a struct literal
e2 := example{
		flag:  true, // COMMA !!!
		counter:   25, // COMMA !!!
		pi: 3.14, // COMMA !!!
	}

// Display the value.
fmt.Printf("%v\n", e1) //=> {false 0 0}
fmt.Printf("%+v\n", e1) //=> {flag:false counter:0 pi:0}
fmt.Printf("%#v\n", e1) //=> main.example{flag:false, counter:0, pi:0} ==> BETTER !


// Declare a variable using an anonymous struct.
ed := struct {
	name  string
	email string
	age   int
}{
	name:  "Ed",
	email: "ed@ardanstudios.com",
	age:   46,
}


////////////////////////////////////////////////////////////////////////////////
// Pointers
////////////////////////////////////////////////////////////////////////////////

//In golang: 
// everything gets passed by value !! https://play.golang.org/p/nNnsK6hWdP
// => Pass var by reference: https://play.golang.org/p/FWmGnVUDoA

// NO pointer arythmetic

// *inc = memory pointed to value inc

// look at /gotraining/01-language_syntax/02-struct_types/my-example-pointers-struct.go


////////////////////////////////////////////////////////////////////////////////
// Constants: numeric, string, bool => NO built-in type
////////////////////////////////////////////////////////////////////////////////
const ui = 12345    // kind: integer
const ti int = 12345        // type: int64
const myUint8 uint8 = 1000 // => NO ! no const with built-in type

// This is an example of constant arithmetic between typed and
// untyped constants. Must have like types to perform math.
const one int8 = 1
const two = 2 * one











