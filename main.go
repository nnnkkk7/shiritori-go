package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("start!!!!!!")
	firstword := "り"
	fmt.Printf("first word %s", firstword)
	t := input()
	nextlast := judge(firstword, t)
	tt := input()
	n := judge(nextlast, tt)
	tt2 := input()
	n2 := judge(n, tt2)
	tt3 := input()
	n3 := judge(n2, tt3)
	fmt.Println(n3)
}

func input() string {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	text := stdin.Text()
	return text
}

func judge(last, next string) string {
	fmt.Println(last)
	// nextのはじめの文字がlastかどうかの判定
	if strings.HasPrefix(next, last) {
		fmt.Println("ok")
	} else {
		fmt.Println("game over")
		os.Exit(1)
	}
	last2 := len(next)
	// 一文字３バイト
	prelast2 := len(next) - 3
	// 最後の文字
	laststrnext := next[prelast2:last2]
	return laststrnext
}

func getLastStr(word string) string {
	last := len(word)
	// 一文字３バイト
	prelast := len(word) - 3
	// 最後の文字
	laststr := word[prelast:last]
	return laststr
}
