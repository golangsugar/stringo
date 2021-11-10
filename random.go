package stringo

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// RandomInt returns a random integer within the given (inclusive) range
func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

// RandomReseed restarts the randonSeeder and returns a random integer within the given (inclusive) range
func RandomReseed(min, max int) int {
	x := time.Now().UTC().UnixNano() + int64(rand.Int())

	rand.Seed(x)

	return rand.Intn(max-min) + min
}
