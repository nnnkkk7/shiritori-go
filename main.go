package main

import (
	"bufio"
	"os"
)

func main() {
	firtword := "り"
	input()
}

func input(word string) {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	text := stdin.Text()
	return text
}
