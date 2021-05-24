package testcode

import "strings"

/*
binary search 二分探索
二分探索とは
https://wa3.i-3-i.info/word1614.html
- 順番に並んだデータを真ん中で半分にして候補を絞っていく検索方法
*/

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var sb strings.Builder // 文字列生成 ゼロ値
	for i := 0; ; i++ {
		for _, str := range strs {
			if i == len(str) || strs[0][i] != str[i] {
				return sb.String()
			}
		}
		sb.WriteByte(strs[0][i])
	}
}
