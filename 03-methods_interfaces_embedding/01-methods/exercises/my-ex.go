// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/Rj0QfwVPhX

// Declare a struct that represents a baseball player. Include name, atBats and hits.
// Declare a method that calculates a players batting average. The formula is Hits / AtBats.
// Declare a slice of this type and initialize the slice with several players. Iterate over
// the slice displaying the players name and batting average.
package main

// Add imports.
import "fmt"

type player struct {
	name   string
	atBats int
	hits   int
}

// Declare a struct that represents a ball player.
// Include field called name, atBats and hits.

// Declare a method that calculates the batting average for a batter.
func (p player) average() float64 {
	return float64(p.atBats) / float64(p.hits)
}

// main is the entry point for the application.
func main() {
	// Create a slice of players and populate each player
	// with field values.
	players := []player{
		{"joe1", 10, 5},
		{"joe2", 12, 5},
		{"joe3", 15, 5},
	}
	// Display the batting average for each player in the slice.
	for _, player := range players {
		fmt.Printf("%v\n", player.average())
	}
}
