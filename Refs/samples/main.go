package main

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"
)

func main() {
	var sample map[string]string
	dec := json.NewDecoder(os.Stdin)
	dec.Decode(&sample)
	enc := json.NewEncoder(os.Stdout)
	for i := 0; i < 10; i++ {
		output := make(map[string]string, len(sample))
		for k := range sample {
			output[k] = getRandomColor()
		}
		enc.Encode(output)
	}

}

// getRandomColor returns a random color from the specified options.
func getRandomColor() string {
	rand.Seed(time.Now().UnixNano()) // Initialize the random number generator.
	colorNumber := rand.Intn(4)      // Generate a random number between 0 and 3.

	switch colorNumber {
	case 0:
		return "green"
	case 1:
		return "orange"
	case 2:
		return "red"
	default:
		return "nocolor"
	}
}
