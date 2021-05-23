package testcode

import (
	"fmt"
	"strings"
)

// ローマ数字の文字列の合計値を整数型として出力
func RomanToInteger(s string) int {
	summe := 0
	i := 0
	s = strings.ToLower(s)
	fmt.Println(len(s))
	// ローマ数字のmap
	dickey := map[byte]int{'i': 1, 'v': 5, 'x': 10, 'l': 50, 'c': 100, 'd': 500, 'm': 1000}
	for i < len(s) {
		if i+1 < len(s) {
			if dickey[s[i]] < dickey[s[i+1]] {
				summe += dickey[s[i+1]] - dickey[s[i]]
				i += 2
				continue
			} else {
				summe += dickey[s[i]]
			}
		} else {
			summe += dickey[s[i]]

		}
		i += 1
	}
	return summe
}
