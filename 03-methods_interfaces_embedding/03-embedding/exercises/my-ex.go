// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/c9Qrsq8QFe

// Copy the code from the template. Declare a new type called hockey
// which embeds the sports type. Implement the matcher interface for hockey.
// When implementing the Search method for hockey, call into the Search method
// for the embedded sport type to check the embedded fields first. Then create
// two hockey values inside the slice of matchers and perform the search.
package main

import (
	"fmt"
	"strings"
)

// matcher defines the behavior required for performing searches.
type matcher interface {
	Search(searchTerm string) bool
}

// sport represents a sports team.
type sport struct {
	team string
	city string
}

// Search checks the value for the specified term.
func (s sport) Search(searchTerm string) bool {
	return strings.Contains(s.team, searchTerm) || strings.Contains(s.city, searchTerm)
}

// Declare a struct type named hockey that represents specific
// hockey information. Have it embed the sport type first.
type hockey struct {
	sport
	country string
}

// Implement the matcher interface for hockey.
func (h hockey) Search(searchTerm string) bool {
	// Make sure you call into Search method for the embedded
	// sport type.
	return h.sport.Search(searchTerm) || strings.Contains(h.country, searchTerm)
}

// main is the entry point for the application.
func main() {
	// Define the term to search.
	searchTerm := "Canada"
	// Create a slice of matcher values to search.
	nhl := []hockey{
		hockey{sport{"Canucks", "Vancouver"}, "Canada"},
		hockey{sport{"Capitals", "Washington"}, "USA"},
	}
	// Display what we are searching for.

	// Range of each matcher value and check the search term.
	for _, nhl_team := range nhl {
		ok := nhl_team.Search(searchTerm)
		if ok {
			fmt.Printf("team %s of %s located in Canada\n", nhl_team.team, nhl_team.city)
		}

	}
}
