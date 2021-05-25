package testcode

import "fmt"

// 任意の英語の文章(すべて小文字)で、各文字が出現する回数を出力するプログラムを作成しましょう。
// ex) hello world
func Index(word string) {
	var elements []string  // 文字を入れるスライス
	var elementCount []int //文字の出現回数を入れるスライス

	for _, w := range word {
		letter := fmt.Sprintf("%c", w)
		// 初めて出現したときの処理
		if !isExist(letter, elements) {
			elements = append(elements, letter)
			elementCount = append(elementCount, count(letter, word))
		}
	}

	for i := 0; i < len(elements); i++ {
		fmt.Printf("letter: %s count: %d\n", elements[i], elementCount[i])
	}
}

// すでにカウントした文字かどうかを判定
func isExist(letter string, elements []string) bool {
	for _, e := range elements {
		if letter == e {
			return true
		}
	}
	return false
}

// wordのなかに該当の文字が何回出現するのかをカウント
func count(letter string, word string) int {
	count := 0
	for _, w := range word {
		l := fmt.Sprintf("%c", w)
		if l == letter {
			count++
		}
	}
	return count
}
