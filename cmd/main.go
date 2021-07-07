package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Tom/CowsAndBulls/pkg/generate"
	"github.com/Tom/CowsAndBulls/pkg/guess"
	"github.com/Tom/CowsAndBulls/pkg/input"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	secret := generate.Secret(2)

	fmt.Println("[== Cows and Bulls ==]")

	for {
		fmt.Print("Enter your guess: ")

		i, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		i = strings.ReplaceAll(i, "\n", "")
		g, err := input.ValidateGuess(secret, i)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		result, err := guess.Compare(secret, g)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Println(result.ToString())
		}
	}
}
