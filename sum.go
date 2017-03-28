package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please two integer args.")
		os.Exit(1)
	}

	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Please two integer args.")
		os.Exit(1)
	}

	b, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Please two integer args.")
		os.Exit(1)
	}

	fmt.Println(Sum(a, b))
}

func Sum(a int, b int) int {
	return a + a
}
