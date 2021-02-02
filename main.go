package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	firstword := "り"
	t := input()
	judge(firstword, t)

}

func input() string {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	text := stdin.Text()
	return text
}

func judge(word, next string) {
	last := len(word)
	prelast := len(word) - 3
	laststr := word[prelast:last]

	fmt.Println(laststr)
	if strings.HasSuffix(next, laststr) {
		fmt.Println("ok")
	} else {
		fmt.Println("game over")
		os.Exit(1)
	}
}
