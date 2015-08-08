// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/MxcJ581bt6

// Download any document from the web and display the content in
// the terminal and write it to a file at the same time.
package main

// Add imports.
import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// main is the entry point for the application.
func main() {
	// Download the RSS feed for "http://www.goinggo.net/feeds/posts/default".
	// Check for errors.
	resp, err := http.Get("http://www.goinggo.net/feeds/posts/default")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Arrange for the response Body to be Closed using defer.
	defer resp.Body.Close()
	// Declare a slice of io.Writer interface values.
	var writers []io.Writer
	// Append stdout to the slice of writers.
	writers = append(writers, os.Stdout)
	// Open a file named "goinggo.rss" and check for errors.
	file, err := os.Create("goinggo.rss")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Close the file when the function returns.
	defer file.Close()
	// Append the file to the slice of writers.
	writers = append(writers, file)
	// Create a MultiWriter interface value from the writers
	// inside the slice of io.Writer values.
	multiWriters := io.MultiWriter(writers...)
	// Write the response to both the stdout and file using the
	// MultiWriter. Check for errors.
	_, err = io.Copy(multiWriters, resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
}
