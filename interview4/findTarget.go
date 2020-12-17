package interview4

/*

题目：二维数组中的查找
在一个二维数组中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

*/

/*

解法1：
首先选取矩阵右上角元素。如果该数字等于目标值，则查找过程结束；如果该数字大于目标值，那么该列的数字都大于目标值，则可以剔除该数字所在列；如果该数字小于目标值，那么该数字所在行的所有元素都小于目标值，则可以剔除该数字所在行。也即要查找的数字如果不在右上角，那么每一次都可以剔除一行或者一列，这样每一步都可以缩小查找范围，知道找到目标值或者查找完整个矩阵均未找到，最后查找结束。遍历完整个二维数组需要m+n趟（m，n分别为二维数组的行数和列数），故时间复杂度为O(m+n)。空间复杂度为O(1)。
注意：该解法还可以选取数组的左下角元素开始遍历，但不能选取左上角或者右下角元素开始。

*/

func searchMatrix(matrix [][]int, target int) bool {
	var found bool
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return found
	}

	for row, column := 0, len(matrix[0])-1; row < len(matrix) && column >= 0; {
		if matrix[row][column] == target {
			found = true
			break
		}
		if matrix[row][column] < target {
			row++
		} else {
			column--
		}
	}

	return found
}
