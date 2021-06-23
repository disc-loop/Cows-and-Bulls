package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Cows and Bulls")
	fmt.Println("Enter your guess:")
	for {
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
	}
}
