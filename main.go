package main

import (
	"fmt"
	"math"
	"time"
)


func main() {
	a := "mitcmusufhgr"
	b := "mtacnusguhgr"
	now := time.Now()
	LwstBT(a, b)
	fmt.Printf("%d, ", minDist)
	fmt.Println("耗时: ", time.Since(now))
	now = time.Now()
	fmt.Printf("%d, ", LwstBTDB(a, b))
	fmt.Println("耗时: ", time.Since(now))
}

func LwstBTDB(a, b string) int {
	// 创建动态矩阵
	minDist := make([][]int, len(a))
	for i := range a {
		minDist[i] = make([]int, len(b))
	}
	// 初始化第一行第一列
	for j := 0; j < len(minDist[0]); j++ {
		// 第一行
		if a[0] == b[j] {
			minDist[0][j] = j
		} else if j != 0 {
			minDist[0][j] = minDist[0][j-1]+1
		} else {
			minDist[0][j] = 1
		}
	}
	// 第一列
	for i := 0; i < len(minDist); i++ {
		// 第一列
		if b[0] == a[i] {
			minDist[i][0] = i
		} else if i != 0 {
			minDist[i][0] = minDist[i-1][0] + 1
		} else {
			minDist[i][0] = 1
		}
	}
	// 从第一行开始到结束填表
	for i := 1; i < len(minDist); i++ {
		for j := 1; j < len(minDist[0]); j++ {
			if a[i] == b[j] {
				minDist[i][j] = min(minDist[i][j-1]+1, minDist[i-1][j]+1, minDist[i-1][j-1])
			} else {
				minDist[i][j] = min(minDist[i][j-1]+1, minDist[i-1][j]+1, minDist[i-1][j-1]+1)
			}
		}
	}
	return minDist[len(minDist)-1][len(minDist[0])-1]
}

func min(x, y, z int) int {
	if x > y {
		x = y
	}
	if x > z {
		return z
	}
	return x
}

var minDist =  math.MaxInt64 // 存储结果

func LwstBT(a, b string) {
	minDist = math.MaxInt64
	f(a, b, 0, 0, 0)
}

func f(a, b string, i, j int, edit int) {
	// 判断是否结束
	if i == len(a) || j == len(b) {
		// 删除i剩下的字符
		if i < len(a) {
			edit += len(a)-i
		}
		if j < len(b) {
			edit += len(b)-j
		}
		if edit < minDist {
			minDist = edit
		}
		return
	}
	if a[i] == b[j] { // 相等不用增加
		f(a, b, i+1, j+1, edit)
	} else {
		// 删除a[i]或者b[j]前添加一个字符
		f(a, b, i+1, j, edit+1)
		// 删除b[j]或者a[i]前添加一个字符
		f(a, b, i, j+1, edit+1)
		// 将a[i]和b[j]替换为相同字符
		f(a, b, i+1, j+1, edit+1)
	}
}