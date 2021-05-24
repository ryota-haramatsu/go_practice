package testcode

import (
	"fmt"
	"strings"
)

// スタック
func IsValid(s string) bool {
	// スタックを利用
	// pop push
	// () {} [] のペアを作成
	// [ (  ) [  ] {  } ]
	// "]"など一文字のみのときも想定
	// s := ")(){}"

	pairs := map[string]string{"(": ")", "{": "}", "[": "]"}
	arr := []string{}
	arr = strings.Split(s, "")
	fmt.Println(arr)
	stack_list := []string{}
	var check bool

	if len(arr) > 1 {
		for i := 0; i < len(arr); i++ {
			for key, val := range pairs {
				// pop スタックのスライスに1つ以上あり、スタックの最後の要素がkeyと同じかつ対象要素がvalと同じ時
				if len(stack_list) > 0 && stack_list[len(stack_list)-1] == key && arr[i] == val {
					stack_list = stack_list[:len(stack_list)-1] // pop処理 最後の要素を取り除く
					break
				} else {
					// push pop条件以外で対象要素がkeyかvalと等しければpushしてループを抜ける
					if arr[i] == key || arr[i] == val {
						stack_list = append(stack_list, arr[i])
						break
					}
				}
			}
		}
		if len(stack_list) == 0 {
			check = true
		}
	} else {
		check = false
	}
	// fmt.Println(check)
	// Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Parentheses.
	// Memory Usage: 2.8 MB, less than 8.06% of Go online submissions for Valid Parentheses.
	return check
}
