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
	begin(firstword)
}

func begin(word string) {
	fmt.Printf("first word %s", word)
	text := input()
	judge(word, text)
}

func input() string {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	text := stdin.Text()
	return text
}

func judge(last, next string) {
	// nextのはじめの文字がlastかどうかの判定
	if strings.HasPrefix(next, last) {
		fmt.Println("ok")
	} else {
		fmt.Println("game over")
		os.Exit(1)
	}
	lastStrNext := getLastStr(next)
	judge(lastStrNext, input())
}

func getLastStr(word string) string {
	last := len(word)
	// 一文字３バイト
	prelast := len(word) - 3
	// 最後の文字を取得
	laststr := word[prelast:last]
	return laststr
}
