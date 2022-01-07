// 排序算法
package sortalgo

// 归并排序-分治法
func MergeSort(nums []int) []int {
    return mergeSort(nums)
}

func mergeSort(nums []int) []int {
    if len(nums) <= 1 {
	return nums
    }
    // 分治法
    mid := len(nums) / 2
    left := mergeSort(nums[:mid])
    right := mergeSort(nums[mid:])
    // 合并
    result := merge(left, right)
    return result
}

func merge(left, right []int) (result []int) {
    leftIndex := 0
    rightIndex := 0
    for leftIndex < len(left) && rightIndex < len(right) {
	if left[leftIndex] > right[rightIndex] {
	    result = append(result, right[rightIndex])
	    rightIndex++
	} else {
	    result = append(result, left[leftIndex])
	    leftIndex++
	}
    }
    // 剩余部分合并
    result = append(result, left[leftIndex:]...)
    result = append(result, right[rightIndex:]...)
    return
}
