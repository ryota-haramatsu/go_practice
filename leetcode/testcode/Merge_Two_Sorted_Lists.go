package testcode

// 連結リストと配列について
// 配列
var a = [4]int{1, 2, 3, 4}
// 配列要素の中間に別の要素が必要になっても、固定長だからできず
// size5の配列を新しく作る必要がある

// しかし連結リスト(Linked List)であれば
// リンク情報を書き換えるだけで 1 2 5 3 4 のように2と3の間に5を入れることが可能


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//  連結リスト
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

// 	// var mergedList []int
// 	// if l1 != nil || l2 != nil {
// 	// 	mergedList = append(l1, l2...) // l1とl2をマージ
// 	// 	sort.Slice(mergedList, func(i, j int) bool {
// 	// 		return mergedList[i] < mergedList[j]
// 	// 	})
// 	// }

// }
