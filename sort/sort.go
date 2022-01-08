// 排序算法
package sort

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
    lIndex := 0
    rIndex := 0
    for lIndex < len(left) && rIndex < len(right) {
	if left[lIndex] > right[rIndex] {
	    result = append(result, right[rIndex])
	    rIndex++
	} else {
	    result = append(result, left[lIndex])
	    lIndex++
	}
    }
    // 剩余部分合并
    result = append(result, left[lIndex:]...)
    result = append(result, right[rIndex:]...)
    return
}

// 快速排序
func QuickSort(nums []int) []int {
    quickSort(nums, 0, len(nums)-1)
    return nums
}

func quickSort(nums []int, start, end int) {
    if start < end {
	pivot := partition(nums, start, end)
	quickSort(nums, start, pivot-1)
	quickSort(nums, pivot+1, end)
    }
}
func partition(nums []int, start, end int) int {
    p := nums[end]
    i := start
    for j := start; j < end; j++ {
	if nums[j] < p {
	    nums[i], nums[j] = nums[j], nums[i]
	    i++
	}
    }
    nums[i], nums[end] = nums[end], nums[i]
    return i 
}
