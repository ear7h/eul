package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ear7h/eul/grammar"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		res, _ := grammar.Parse("stdin", []byte(scanner.Text()))
		fmt.Println(res.(grammar.Scalar).Val())
	}
}
