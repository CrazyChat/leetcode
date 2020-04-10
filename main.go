package main

import "fmt"

func main() {
	T := 0	// T 组测试数据
	// 每次更新
	n, m, c := 0, 0, 0	// n*m 方格大小，c种颜色
	// 赋值
	_, _ = fmt.Scanf("%d", &T)
	for {
		n, _ = fmt.Scanf("%d %d %d", &n, &m, &c)
		if n == 0 {
			break
		}
		num := make([]int, c)	// 染料的数量
		// 矩阵
		trc := make([][]int, n)
		for i := range trc {
			trc[i] = make([]int, m)
		}
		// 填充
		for i := 0; i < len(trc); i++ {
			for j := 0; j < len(trc); j++ {
				// 有染料

				// 无染料，退出

				fmt.Println("NO")
				return
			}
		}
	}

}




//func main() {
//	n := 0
//	// 赋值
//	_, _ = fmt.Scanf("%d", &n)
//	house := []int{}
//	housey := 0
//	for i := 0; i < n; i++ {
//		_, _ = fmt.Scanf("%d %d", &house[i], housey)
//	}
//	// 动态规划表, arr[0][j] = house[j], arr[i][j] = house[j]
//	arr := make([][]int, n)
//	// 初始化第一行和第一列
//	arr[0][0] = 0
//	for j := 1; j < len(arr[0]); j++ {
//		arr[0][j] =  f(house[j], house[0]) + arr[0][j-1]
//	}
//	for i := 1; i < len(arr); i++ {
//		arr[i][0] = arr[0][i]
//	}
//	// 逐行填表
//	for i := 1; i < len(arr); i++ {
//		for j := 1; j < len(arr[0]); j++ {
//			arr[i][j] = f(house[i], house[j]) + arr[i][j-1]
//		}
//		if arr[i][len(arr[0])-1] > arr[i][len(arr[0])-1] {
//			arr[i][len(arr[0])-1] = arr[i][len(arr[0])-1]
//		}
//	}
//	fmt.Println(arr[len(arr)-1][len(arr[0])-1])
//}
//
//func f(x, y int) int {
//	if x > y {
//		return x - y
//	}
//	return y - x
//}

/*
a:=0
	b:=0
	for {
		n, _ := fmt.Scan(&a,&b)
		if n == 0 {
			break
		} else {
			fmt.Printf("%d\n",a+b)
		}
	}
 */

/*
n:=0
    ans:=0

    fmt.Scan(&n)
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            x:=0
            fmt.Scan(&x)
            ans = ans + x
        }
    }
    fmt.Printf("%d\n",ans)
 */