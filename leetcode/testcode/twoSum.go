package testcode

func TwoSum(nums []int, target int) []int {
	var output []int // スライス初期化
CheckSum:
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				output = append(output, i, j) // 追加
				break CheckSum
			}
		}
	}
	return output
}
