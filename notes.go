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

// %v	the value in a default format
// %#v	a Go-syntax representation of the value
// %T	a Go-syntax representation of the type of the value
// %t	Boolean: the word true or false
// %d	Integer: base 10
// %s	the uninterpreted bytes of the string or slice
// %p	pointer: base 16 notation, with leading 0x


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

// Time
now := time.Now()


////////////////////////////////////////////////////////////////////////////////
// Functions and errors
////////////////////////////////////////////////////////////////////////////////
// in main, handle error:
u, err := doSomething("input")
if err != nil {
	fmt.Println(err)
	return
}
// TRY TO NOT PUT else block, try to DE-TENT (as much left as possible)

func retrieveUser(name string) (*user, error) { // in a func: error is always the LAST returned value
	// Make a call to get the user in a json response.
	r, err := getUser(name)
	if err != nil {
		return nil, err
	}
}

// => CAUTION: list of return values HAS TO BE inside parenthesis (return1, return2)

////////////////////////////////////////////////////////////////////////////////
// Arrays & Slices
////////////////////////////////////////////////////////////////////////////////

// NO dynamic size array => an array is ALWAYS fixed size
// NO uninitialized var => by default ANY var gets initialized with the 0 of the type

var fruits [5]string // declare unintiliazed array

// Declare an array of 4 integers that is initalized with some values.
numbers := [4]int{10, 20, 30, 40}

// Iterate over the array of strings.
for i, fruit := range strings { // => for index, value := array
	fmt.Println(i, fruit)
}

len(array) // get Array size

// Slices
slice := make([]string, 5, 8) // length of 5 elements and a capacity of 8.
fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))  // Length[5] Capacity[8]

// Dynamic array -> using slices
var data []string
for record := 1; record <= 10240; record++ {
	data = append(data, fmt.Sprintf("Rec: %d", record))
	// append a string to the slice of strings, and return the newly generated slice
}

...user // dynamic nbr of objects of type user
func display(users ...user) {
	for i := 0; i < len(users); i++ {
		fmt.Printf("%+v\n", users[i])
	}
}

// Declare a slice of strings.
names := []string{"Bill", "Lisa", "Jim", "Cathy"}
// Take a slice of index 1 and 2.
slice := names[1:3] // => element of index 3 is NOT included !
//=> slice = ["Lisa", "Jim"]


////////////////////////////////////////////////////////////////////////////////
// Maps
////////////////////////////////////////////////////////////////////////////////
type user struct {
	name  string
	email string
}
users := make(map[string]user)

// Iterate over the map.
for key, value := range users {
	// iteration -> pseudo random order !!! CAUTION !!!
	// no concept of keys order. if you want: use slice from map
	fmt.Println(key, value)
}

// map litteral declaration
users := map[string]user{
		"Rob":     user{"Roy", "Rob"},
		"Ford":    user{"Henry", "Ford"},
		"Mouse":   user{"Mickey", "Mouse"},
		"Jackson": user{"Michael", "Jackson"},
	}

// Map lookup
user_joe, ok := users["Joe"]  // val, ok := map[key] => ok gets true/false based on the presence of key


////////////////////////////////////////////////////////////////////////////////
// Methods
////////////////////////////////////////////////////////////////////////////////

//Methods are functions that are declared with a receiver which binds the method to a type.
type user struct {
	name  string
	email string
}
// notify implements a method with a value receiver.
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}
// affects only a copy of the object => wont affect the original object

// changeEmail implements a method with a pointer receiver.
func (u *user) changeEmail(email string) {
	u.email = email
}
// => affects the pointer to the value => affect the original object itself

////////////////////////////////////////////////////////////////////////////////
// Interfaces
////////////////////////////////////////////////////////////////////////////////

// Interfaces provide a way to declare types that define only behavior. 
// This behavior can be implemented by concrete types, such as struct or named types, via methods.

// notifier is an interface that defines notification type behavior
type notifier interface {
	notify() //method
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(n notifier) {
	n.notify()
}

// methods in Go are about behaviors
// => polymorphism allowed: dont care about the type but the behavior (the fact that the object implements the interface)
// https://play.golang.org/p/Lo1ucf1e9d


////////////////////////////////////////////////////////////////////////////////
// Embedding / Composition
////////////////////////////////////////////////////////////////////////////////

// Embedding types allow us to share state or behavior between types.
// The inner type never loses its identity.
// !! This is not inheritance.
// Through promotion, inner type fields and methods can be accessed through the outer type.
// The outer type can override the inner type's behavior.

type user struct {
	name  string
	email string
}
type admin struct {
	user  // Embedded Type !!!
	level string
}


// => // We can acces the inner type's method & fields directly.
	ad.user.notify()

// => // The inner type's method & fields are promoted AT RUNTIME. 
	ad.notify()



////////////////////////////////////////////////////////////////////////////////
// Packaging
////////////////////////////////////////////////////////////////////////////////


package animals

type Dog struct {
	Name string // exportable field -> start with capital letter 'N'
	BarkStrength int // exportable field -> start with capital letter 'B'
	age int // !!!! NON exportable field -> start with lower case letter 'a'
}

// => age is unexported field from an exported struct



























