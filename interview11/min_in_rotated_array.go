package interview11

import (
	"math"
)

// 题目：
// 把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。
// 例如，数组[3,4,5,1,2]为[1,2,3,4,5]的一个旋转，该数组的最小值为1

// 解法：考虑到输入是一个递增排序的数组的旋转，该输入数组的两个子数组也是递增排序的。在一个排序的数组中寻找最值问题的最理想解法就是二分查找。
// 根据二分查找的一般思路，用两个指针分别指向数组的第一个index1后最后一个元素index2，根据题意，第一个元素应该大于或等于最后一个元素。
// 然后找到中间位置Mid的元素，如果该元素位于前面的递增子数组，那么该元素一定大于等于index1指向的元素，此时数组中最小的元素应该位于该元素Mid的后面。
// 我们就可以把第一个指针index1指向该位置Mid。
// 如果该元素位于后面的递增子数组，那么该元素一定小于等于index2指向的元素，此时数组中最小的元素应该位于该元素Mid的前面，我们就可以把第二个指针index2
// 指向该Mid。
// 如此，不管是移动第一个指针还是第二个指针，查找范围都会缩小为原来的一半。另外，第一个指针总是会指向前面递增四数组的元素，第二个指针总是会指向后面递
// 增子数组的元素。最终，第一个指针index1会指向前面的递增子数组的最后一个元素，而第二个指针index2会指向后面递增子数组的第一个元素。也即他们会指向两个
// 相邻的元素，而第二个指针则刚好指向最小的元素。

// Min ...
func Min(arr []int) int {
	if len(arr) == 0 {
		return math.MinInt64
	}

	index1, index2 := 0, len(arr)-1
	mid := index1
	for arr[index1] >= arr[index2] {
		if index2-index1 == 1 {
			mid = index2
			break
		}
		mid = (index2 + index1) / 2

		// 注意：如果index1和index2以及mid指向的元素三者相等，则只能进行顺序查找
		if arr[index2] == arr[index1] && arr[index1] == arr[mid] {
			return MinInOrder(arr, index1, index2)
		}

		if arr[index1] <= arr[mid] {
			index1 = mid
		} else if arr[mid] <= arr[index2] {
			index2 = mid
		}
	}

	return arr[mid]
}

func MinInOrder(arr []int, index1, index2 int) int {
	res := arr[index1]
	for i := index1 + 1; i <= index2; i++ {
		if arr[i] < res {
			res = arr[i]
		}
	}

	return res
}
