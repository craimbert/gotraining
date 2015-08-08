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
// %T	a Go-syntax representation of the type of the value !!!!!
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

////////////////////////////////////////////////////////////////////////////////


// Instead of if / elsif / else block => use switch/case without any value -> truthness of expression

switch { // -> true implied
		case expr1:
			do_something_1

		case expr2:
			do_something_2
		}

// infinite loop == while true
for {
	// do stuff
	return
}



////////////////////////////////////////////////////////////////////////////////
// goroutines
////////////////////////////////////////////////////////////////////////////////

// Goroutines are functions that are created and scheduled to be run indenpently
//$ export GOMAXPROCS=4 => to enable more than 1 go routines to run concurrently

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

// Add a count of two goroutines allowed
wg.Add(2)

// Create a goroutine from the lowercase function.
go lowercase()

func lowercase() {
	// do stuff...

	wg.Done() // --> give back 1 mutex
}

// in the main or any function that needs to wait when go routines are getting triggered:
wg.Wait()


// Create a goroutine launching an anonymous function
go func() {
	// do something....

	// Decrements the count of the wait group.
	wg.Done()
}() // important () -> HAS TO BE a function call -> same way: object.Do()


////////////////////////////////////////////////////////////////////////////////
// goroutines - race conditions
////////////////////////////////////////////////////////////////////////////////

// A race condition is when two or more goroutines attempt to
// read and write to the same resource at the same time.

// Run: $ go run -race example1.go => add "-race" to pay attention to the race conditions

// mutex is used to define a critical section of code.
var mutex sync.Mutex

// Only allow one goroutine through this
// critical section at a time.
mutex.Lock()
{
	// do stuff
}
mutex.Unlock()
// Release the lock and allow any
// waiting goroutine through.

// Better declaration of mutex WITH the var which needs that mutex:
var (
	mutex sync.Mutex
	counter int // SHARED data
)

// distinguish Readers and Writers => more than 1 reader allowed at same time
var rwMutex sync.RWMutex 

// Capture the current read count.
// Keep this safe though we can due without this call.
atomic.LoadInt64(&readCount)

// Increment the read count value by 1.
atomic.AddInt64(&readCount, 1)


////////////////////////////////////////////////////////////////////////////////
// goroutines - channels
////////////////////////////////////////////////////////////////////////////////

// Channels are a reference type that provide a safe mechanism to share data between goroutines.
// a channel can have any nbr of senders and receivers
// unbuffered & buffered channels

// Create an unbuffered channel (accepting an int value).
channel := make(chan int)

// Send a value into that channel
channel <- val

// Spin off a new go routine
go doSomething(channel)
// with: 
func doSomething(channel chan int) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()
	// -> avoid code replication

	// Receive value
	value, ok := <- channel

	// Close the channel
	close(channel)

}

// same job with buffered or unbuffered channel

////////////////////////////////////////////////////////////////////////////////
// goroutines - errors
////////////////////////////////////////////////////////////////////////////////

type error interface {
	Error() string // an error 'returns' a string, always
}

// define new custom error type: has to contain a string & implement the interface "Error() string"
type MyError struct {
	s string
}

func (e *MyError) Error() string {
	return e.s
}

// Example from JSON package:

	// An UnmarshalTypeError describes a JSON value that was
	// not appropriate for a value of a specific Go type.
	type UnmarshalTypeError struct {
		Value string       // description of JSON value
		Type  reflect.Type // type of Go value it could not be assigned to
	}

	// Error implements the error interface.
	func (e *UnmarshalTypeError) Error() string {
		return "json: cannot unmarshal " + e.Value + " into Go value of type " + e.Type.String()
	}

	// Handle the error based on its type
	err := Unmarshal([]byte(`{"name":"bill"}`), u)
	if err != nil {
		switch e := err.(type) {
		case *UnmarshalTypeError:
			fmt.Printf("UnmarshalTypeError: Value[%s] Type[%v]\n", e.Value, e.Type)
		case *InvalidUnmarshalError:
			fmt.Printf("InvalidUnmarshalError: Type[%v]\n", e.Type)
		default: // generic default error handling
			fmt.Println(err)
	}

// Another way of checking the error		
e := err.(*CustomTypeError)
fmt.Printf("CustomTypeError: value [%s] Type[%v]\n", e.Value, e.Type)

// Create a basic error
errors.New("false flag")


////////////////////////////////////////////////////////////////////////////////
// goroutines - patterns
////////////////////////////////////////////////////////////////////////////////

// infinite loop on the channel as long as channel is opened
for dataReceived := range channel {
	dataReceived.DoSomething()
}
// => no need for "data, ok := <- channel" 



////////////////////////////////////////////////////////////////////////////////
// Testing
////////////////////////////////////////////////////////////////////////////////

// file names end in "_test.go" => mandatory to make compile aware and not compile it
func TestDownload(t *testing.T) {

} 
// => func name is "TestFunctionToTest"

// -> get object t:
t.Log ~ fmt.Println
t.Logf ~ fmt.Printf
t.Fatal -> abort test & print log
t.Errorf -> print error log but does NOT abort

test succeeds <=> test does not fail <=> no t.Fatal gets called

// Run test: 
$ go test // => NO log print except it test fails
$ go test -v // => force log print
$ go test -v ./... // include all files in the further subtree


/////// Example Tests
func ExampleSendJSON() { // tied with function SendJSON
}
// at the end, a comment with "Output:" and then expected result:

// Output:
// {Bill bill@ardanstudios.com}


$ go test -cover
       coverage: 83.3% of statements

$ go test -coverprofile cover.out
	PASS
	coverage: 83.3% of statements
	ok  	github.com/craimbert/gotraining/10-testing/01-testing/example4/handlers	0.011s
	-> generate cover.out

$ go tool cover -html cover.out -o coverage.html
-> generate coverage.html -> open the html page in web browser to see


////////////////////////////////////////////////////////////////////////////////
// Standard Library
////////////////////////////////////////////////////////////////////////////////

// Binary arguments:
	// init is called before main.
	func init() {
		if len(os.Args) != 2 {
			fmt.Fprintln(os.Stderr, "Usage: ./example2 <url>")
			os.Exit(2)
		}
	}

// Binary options:

// init is called before main.
func init() {
	// Let the flag package handle the options; -o for output and -s for silent.
	flag.StringVar(&Config.DestFile, "o", "", "output file")
	flag.BoolVar(&Config.Silent, "s", false, "silent (do not output to stdout)")
	flag.Parse()
	if len(flag.Args()) != 1 {
		fmt.Println("Usage: ./example3 [options] <url>")
		os.Exit(2)
	}
}

////////////////////////////////////////////////////////////////////////////////
// HTTP
////////////////////////////////////////////////////////////////////////////////

with req.URL.Query().Get("key")
$ curl "localhost:4000/charles?key=123"
Hello world 123

with req.URL.Query()
$ curl "localhost:4000/charles?key=123&test=charles"
Hello world map[test:[charles] key:[123]]



////////////////////////////////////////////////////////////////////////////////
// JSON
////////////////////////////////////////////////////////////////////////////////
// Decode JSON 
// https://github.com/craimbert/gotraining/blob/master/11-standard_library/02-encoding/example1/example1.go

CurrencyCode string `json:"currency_code"`
CurrencyCode int `json:"currency_code, string"`
CurrencyCode int `json:""` // => same json key
CurrencyCode int `json:",omitempty"`


// instead of Unmarshall:
r := strings.NewReader(document) // document is the raw string
dec := json.NewDecoder(r)
err := dec.Decode(&ucp) // ucp is the destination struct

// pretty print of jsons
data, err := json.MarshalIndent(&ucp, "", "\t")
fmt.Printf("%s\n", data)



