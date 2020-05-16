package sortMethod

func CountingSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	// 查找数据范围(最大值)
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	// 创建计数数组c
	c := make([]int, max+1)
	// 计算每个元素的个数，存入c中，此时c表示各个元素的个数
	for i := 0; i < len(arr); i++ {
		c[arr[i]]++
	}
	// 依次累加，变成各个元素的前缀和，由该数组可推算元素在排序后数组的位置
	for i := 1; i <= max; i++ {
		c[i] = c[i-1] + c[i]
	}
	// 临时数组，存储排序结果，最后再赋值给arr
	r := make([]int, len(arr))
	// 计数排序
	for i := len(arr) - 1; i >= 0; i-- {
		index := c[arr[i]] - 1
		r[index] = arr[i]
		c[arr[i]]--
	}
	// 拷贝回原数组arr
	for i, v := range r {
		arr[i] = v
	}
}
