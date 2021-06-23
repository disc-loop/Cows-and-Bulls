package response

import (
	"fmt"

	"github.com/Tom/CowsAndBulls/pkg/guess"
)

func ToString(accuracy guess.Accuracy) string {
	if accuracy == (guess.Accuracy{}) {
		return "No matches!"
	}

	return fmt.Sprintf("Bulls: %v, Cows: %v", accuracy.Exact, accuracy.Near)
}
