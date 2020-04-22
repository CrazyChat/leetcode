package interview

import (
	"fmt"
	"strconv"
)

var wayMap = make(map[string]int)

func WaysCount() {
	Way := 0
	// 初始化	2 <= n <= 1000, 1 <= a[i], h0 <= n - 1
	n, h0 := 0, 0
	_, _ = fmt.Scanf("%d %d", &n, &h0)
	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		_, _ = fmt.Scanf("%d", &arr[i])
	}
	// 如果加上h/arr[i]后等于到达终点，只需要走一条
	if h0 == arr[1] {
		Way = dp(arr, 1, h0)
	} else {
		// 否则走两条路
		Way = dp(arr, 1+h0, h0) + dp(arr, 1+arr[1], arr[1])
	}
	fmt.Println(Way%998244353)
}

func dp(arr []int, index, h int) int {
	if index > len(arr)-1 {
		return 0
	}
	if n, ok := wayMap[getMapKey(index, h)]; ok {
		return n
	}
	if index == len(arr)-1 {
		wayMap[getMapKey(index, h)] = 1
		return 1
	}
	temp := dp(arr, index+h, h)
	wayMap[getMapKey(index, h)] = temp
	if h == arr[index] {
		return temp
	}
	temp2 := dp(arr, index+arr[index], arr[index])
	wayMap[getMapKey(index, h)] = temp + temp2
	return temp + temp2
}

func getMapKey(i, j int) string {
	return strconv.Itoa(i) + strconv.Itoa(j)
}
