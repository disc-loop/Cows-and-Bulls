package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Tom/CowsAndBulls/pkg/generate"
	"github.com/Tom/CowsAndBulls/pkg/guess"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	secret := generate.Secret(2)

	fmt.Println("[== Cows and Bulls ==]")

	for {
		fmt.Print("Enter your guess: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		input = strings.ReplaceAll(input, "\n", "")

		result, err := guess.Compare(secret, input)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Println(result.ToString())
		}
	}
}
