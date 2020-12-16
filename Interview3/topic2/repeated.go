package topic2

/*

题目：不修改数组找出重复的数字
在一个长度为n+1的数组里的所有数字都在1~n的范围内，所以数组中至少有一个数字是重复的。请找出数组中任意一个重复的数字，但不能修改输入的数组。例如，如果输入长度为8的数组{2,3,5,4,3,2,5,7}，那么对应的输出是重复的数字2或3.

*/

/*
解法1：
由于不能修改原始数组，故无法采用interview3/topc1的解法3和4。我们可以创建一个长度为n+1的辅助数组，然后逐一把原始数组的每个元素复制到辅助数组。如果原始数组中被复制的数字是m，则把它复制到辅助数组中下标为m的位置。这样就很容易能发现那个数字是重复的。时间复杂度是O(n)，由于需要辅助数组，空间复杂度是O(n)。
*/

func solution1(nums []int) int {
	length := len(nums)
	tmp := make([]int, length)
	for _, n := range nums {
		if tmp[n] != 0 {
			return n
		}
		tmp[n] = n
	}

	return -1
}

/*
解法2：避免使用O(n)的辅助空间
假如没有重复的数字，那么在1~n的范围里只有n个数字。由于数组中存在重复数字，那么数字1~n在数组中出现的次数大于n。因此可以查看某范围的数字在数组中的出现次数。把1~n的数字从中间的数字m分为两部分。如果1~m的数字在数组出现超过了m次， 那么这一半区间内的数字一定在数组中重复出现了；否则，另一半m+1~n的区间内的数字一定包含重复的数字。我们继续把包含重复数字的区间一分为二，知道找到一个重复的数字。此过程和二分查找基本类似，只不过多了一个统计区间内数字出现次数的过程。该解法不需要辅助数组，空间复杂度为O(1)，二分查找的时间复杂度为O(logn)，而每次查找都需要遍历数组以统计数字的出现次数（时间复杂度为O(n)），因此总的时间复杂度为O(nlongn)。注意：此种解法只能找到一个重复的数字，即使数组中存在多个重复的数字。
*/

func solution2(nums []int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	start := 1
	end := length - 1
	for start <= end {
		middle := ((end - start) >> 1) + start
		count := count(nums, start, middle)
		if end == start {
			if count > 1 {
				return end
			}
			break
		}
		if count > end-start+1 {
			end = middle
		} else {
			start = middle + 1
		}
	}

	return -1
}

func count(nums []int, start, end int) int {
	count := 0
	for _, m := range nums {
		if m <= end && m >= start {
			count++
		}
	}

	return count
}
