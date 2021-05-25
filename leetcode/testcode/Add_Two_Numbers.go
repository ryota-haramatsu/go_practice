package testcode

/*
2. Add Two Numbers
空ではない連結リストを与えられるので与えられた数字をそれぞれ桁ごとに足していき、
反転させて返すようなアルゴリズムを設計
例
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}


func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := new(ListNode)
	cur := head
	for l1 != nil || l2 != nil || carry != 0 {
		n1, n2 := 0, 0
		if l1 != nil {
			n1, l1 = l1.Val, l1.Next
		}
		if l2 != nil {
			n2, l2 = l2.Val, l2.Next
		}
		num := n1 + n2 + carry
		carry = num / 10
		cur.Next = &ListNode{Val: num % 10, Next: nil}
		cur = cur.Next
	}
	return head.Next
}

////////連結リスト
// セル?orノード?
type Cell struct {
	item int
	next *Cell // 自己参照構造体
}

// 連結リスト
type List struct {
	top *Cell
}

// セルの生成
func newCell(x int, cp *Cell) *Cell {
	newcp := new(Cell)
	newcp.item, newcp.next = x, cp
	return newcp
}

// 連結リストの生成
func newList() *List {
	lst := new(List)
	lst.top = new(Cell)
	return lst
}

// n番目のセルを求める処理
func (cp *Cell) nthCell(n int) *Cell {
	i := -1
	for cp != nil {
		if i == n {
			return cp
		}
		i++
		cp = cp.next
	}
	return nil
}

// n番目の要素を取得する
func (lst *List) nth(n int) (int, bool) {
	cp := lst.top.nthCell(n)
	if cp == nil {
		return 0, false
	}
	return cp.item, true
}
