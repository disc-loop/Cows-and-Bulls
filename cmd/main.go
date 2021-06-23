package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Tom/CowsAndBulls/pkg/generate"
	"github.com/Tom/CowsAndBulls/pkg/guess"
	"github.com/Tom/CowsAndBulls/pkg/response"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	secret := generate.Secret(2)

	fmt.Println("[== Cows and Bulls ==]")

	for {
		fmt.Print("Enter your guess: ")
		input, _ := reader.ReadString('\n')
		input = strings.ReplaceAll(input, "\n", "")
		result := guess.Compare(secret, input)
		fmt.Println(response.ToString(result))
	}
}
