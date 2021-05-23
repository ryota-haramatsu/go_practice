package testcode

import "strconv"

// 文字列を反転するメソッド "こんにちは" -> "はちにんこ"
func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

func isPalindrome(x int) bool {
	// 最初と最後のみが同じ数字とは限らない
	// リバースして元の数字と同じになるかどうかで判定
	// string -> slice -> reverse -> string -> integer
	str_x := strconv.Itoa(x) // 文字列にする
	reverse_str := reverse(str_x)
	reverse_int, _ := strconv.Atoi(reverse_str)
	if x == reverse_int {
		return true
	} else {
		return false
	}

}
