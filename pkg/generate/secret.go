package generate

import (
	"fmt"
	"math/rand"
)

// Generates a random 4 digit number based on a seed. 

const limit = 9999

func Secret(s int64) string {
	rand.Seed(s)
	num := rand.Intn(limit)
	secret := fmt.Sprintf("%04d", num)

	return secret
}
