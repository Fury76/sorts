package main

import (
"fmt"
	"os"
)

func main() {
	arr := []int{10, 9, 5, 7, 3, 5, 2, 9, 4, 6, 10}

	res := SelectSortByFury(arr)// 选择排序
	//res := InsertionSort(arr) // 插入排序
	//res := InsertSortByFury(arr) // 插入排序优化版
	//res := BubbleSortByFury(arr) // 冒泡排序
	//res := MergeSort(arr) // 归并排序
	//res := QuickSort(arr) // 快速排序
	fmt.Print(res)
	os.Open()
}

//选择排序
//思路：每次循环找出最小的数，跟数组第一个数交换顺序，接下来在剩余的数里重复以上逻辑
func SelectionSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			//只要找到比要比较的值小的值，就更新min的位置，循环一圈就能找到最小的值的位置
			if arr[j] < arr[min] {
				min = j
			}
		}
		//交换最小值与这一次循环最左边值的位置
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

func SelectSortByFury(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		min := i
		for j := i+1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}

		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

//插入排序，类似扑克牌起牌，将未排序的数据插入到已排序的数据中
func InsertionSort(arr []int) []int {
	length := len(arr)
	for i := 1; i < length; i++ {
		for j := i; j > 0; j-- {
			//如果要比较的数据小于左边的数据，则交换位置
			if arr[j-1] > arr[j] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}
	return arr
}

func InsertSortByFury(arr []int) []int {
	length := len(arr)
	for i := 1; i < length; i++ {
		for j := i; j > 0 ; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			} else {
				break
			}
		}
	}

	return arr
}

//插入排序优化版，用赋值代替交换操作
func InsertionSortPro(arr []int) []int {
	length := len(arr)
	for i := 1; i < length; i++ {
		temp := arr[i] //复制一份待比较的值
		var j int
		for j = i; j > 0; j-- {
			//如果左边数据大于待比较待值，则将左边数据赋值给右边的（往右挪一位），否则停止比较
			if arr[j-1] > temp {
				arr[j] = arr[j-1]
			} else {
				break
			}
		}
		arr[j] = temp //找到合适的位置了（左边不再比该值大），将刚刚待比较的值赋值给这个元素
	}
	return arr
}

//冒泡排序，每次和相邻的元素比较，内层每循环一次会把最大的循环到最后
func BubbleSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		//j < length -i -1 原因：每循环一次，最后一位数已排好，不用再比
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func BubbleSortByFury(arr []int) []int {
	length := len(arr)
	for i := 0; i < length - 1; i++ {
		over := false
		for j := 0; j < length - i - 1; j++ {
			if arr[j] > arr[j+1] {
				over = true
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		if !over {
			break
		}
	}

	return arr
}

//冒泡排序优化版，如果某次循环发现没有需要交换的元素，则认为整个排序已完成
func BubbleSortPro(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		over := false
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				over = true
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		if over == false {
			break
		}
	}
	return arr
}


//归并排序
func MergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	i := length / 2
	left := MergeSort(arr[0:i])
	right := MergeSort(arr[i:])
	res := merge(left, right)
	return res
}

//合并数组
func merge(left, right []int) []int {
	result := make([]int, 0)
	m, n := 0, 0
	l, r := len(left), len(right)
	//比较两个数组，谁小把元素值添加到结果集内
	for m < l && n < r {
		if left[m] > right[n] {
			result = append(result, right[n])
			n++
		} else {
			result = append(result, left[m])
			m++
		}
	}
	//如果有一个数组比完了，另一个数组还有元素的情况，则将剩余元素添加到结果集内
	result = append(result, right[n:]...)
	result = append(result, left[m:]...)
	return result
}

//快排，以第一个值为标准，小于此值的放左边，大于此值放右边，将第一个值放中间，在分好的数组里如此往复
func QuickSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	p := 0
	res := quickSort(arr, p, length-1)
	return res
}

func QuickSortByFury(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}

	p := 0
	return quickSortByFury(arr, p, length-1)
}

//递归方法
func quickSort(arr []int, p int, r int) []int {
	if p >= r {
		return arr
	}
	q := partition(arr, p, r)
	quickSort(arr, p, q-1)
	quickSort(arr, q+1, r)
	return arr
}

func quickSortByFury(arr []int, p int, r int) []int {
	if p >= r {
		return arr
	}

	return arr
}
//排序并返回pivot
func partition(arr []int, p int, r int) int {
	k := arr[p]
	j := p
	for i := p; i < r; i++ {
		if k > arr[i] {
			arr[i], arr[j] = arr[j], arr[i]
			j++
		}
	}
	arr[r], arr[j] = arr[j], arr[r]
	return j
}
