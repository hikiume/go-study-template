package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadLine(prompt string) string {
	fmt.Print(prompt)

	scanner := bufio.NewScanner(os.Stdin)
	
	scanner.Scan()

	input := scanner.Text()

	return strings.TrimSpace(input)

}
