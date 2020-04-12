package sortMethod

func ShellSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}
	for gap := length/2; gap >= 1; gap /= 2 {
		// i = gap，i从gap开始就可以了，即每个分组的第二个元素
		for i := gap; i < length; i++ {
			shellInsert(arr, gap, i)
		}
	}
}

/*
arr: 排序数组
gap: 希尔排序步长
i: 当前排序元素
shellInsert 对 i 元素执行它所在分组的插入排序
 */
func shellInsert(arr []int, gap int, i int) {
	value := arr[i]
	j := i - gap
	for ; j >= 0; j = j - gap {
		if arr[j] > value {
			arr[j+gap] = arr[j]
		} else {
			break
		}
	}
	arr[j+gap] = value
}
