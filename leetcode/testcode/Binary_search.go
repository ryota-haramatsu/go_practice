package testcode

/*
二分探索: O(log n)時間
線形探索: O/(n)時間
*/

// https://www.golangprograms.com/golang-program-for-implementation-of-binary-search.html
func BinarySearch(needle int, haystack []int) bool {
	low := 0                  // 最初
	high := len(haystack) - 1 // 最後

	for low <= high {
		median := (low + high) / 2 // 真ん中の要素

		// 対象が真ん中の要素より大きい時はlowをmedianより1つ大きい数字にする(右半分)
		if haystack[median] < needle {
			low = median + 1
		} else {
			// 対象が真ん中の要素より小さい時はheighをmedianより1つ小さい数字にする(左半分)
			high = median - 1
		}
	}
	// lowがスライスの要素数と同じ、もしくはlowの要素が対象と違う場合
	if low == len(haystack) || haystack[low] != needle {
		return false
	}
	return true
}

// func main() {
// 	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
// 	fmt.Println(BinarySearch(9, items))
// }
