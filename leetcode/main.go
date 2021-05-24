package main

import (
	"fmt"
)

func main() {
	strs := []string{"flower", "flow", "flight"}
	if len(strs) == 0 {
		fmt.Println("")
	}

	pre := []byte{}
	for j := 0; j < len(strs[0]); j++ {
		for i := 1; i < len(strs); i++ {
			if j >= len(strs[i]) || strs[0][j] != strs[i][j] {
				fmt.Println(string(pre))
			}
		}
		pre = append(pre, strs[0][j])
	}
	fmt.Println(string(pre))
}
