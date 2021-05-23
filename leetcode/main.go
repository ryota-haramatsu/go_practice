package main

import (
	"container/list"
	"go_practice/leetcode/testcode"
)

func main() {
	// #1
	// nums := []int{2, 7, 11, 15}
	// target := 9
	// testcode.TwoSum(nums, target)
	// fmt.Println(testcode.TwoSum(nums, target))
	l1 := list.New()
	l2 := list.New()
	testcode.AddTwoNumbers(l1, l2)
}
