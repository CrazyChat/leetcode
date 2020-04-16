package main

/*
360 - 最后赢家
 */

//var saveMap = make(map[float32]float32)
//
//func main() {
//	saveMap = make(map[float32]float32)
//	n, a0 := 0, 0
//	_, _ = fmt.Scanf("%d %d", &n, &a0)
//	if n == 1 {
//		var molecule float32 = 1
//		fmt.Printf("%.5f\n", molecule/float32(a0+1))
//		return
//	}
//	arr := make([]float32, n)
//	arr[0] = float32(a0)
//	fmt.Printf("%.5f\n", f(arr, n-1))
//}
//
//func f(arr []float32, n int) float32 {
//	if n == 1 {
//		return 1 / (arr[0]+1)
//	} else {
//		if v, ok := saveMap[arr[n-1]]; ok {
//			return v
//		}
//		var temp float32
//		cur := arr[n-1]
//		denominator := 1 /cur
//		var end float32 = 1
//		for denominator-1 >= end {
//			temp += cur
//			cur = 1 / (denominator-1)
//			denominator--
//		}
//		// 存储数据
//		saveMap[arr[n]] = arr[n-1] * temp
//		return arr[n-1] * temp
//	}
//}

