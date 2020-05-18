package mix

import (
	"math"
)

/*
如何编程实现“求一个数的平方根”？要求精确到小数点后 6 位
*/

/*
解法：
因为要精确到后六位，可以先用二分查找出整数位，然后再二分查找小数第一位，第二位，到第六位。
整数查找很简单，判断当前数小于+1后大于即可找到，
小数查找举查找小数后第一位来说，从x.0到(x+1).0，查找终止条件与整数一样，当前数小于，加0.1大于，
后面的位数以此类推，可以用x*10^(-i)通项来循环或者递归，终止条件是i>6，
想了一下复杂度，每次二分是logn，包括整数位会查找7次，所以时间复杂度为7logn。空间复杂度没有开辟新的储存空间，空间复杂度为1。
*/

// precision 要求的精度
func SquareRoot(num, precision int) float64 {
	integer := IntegerSqrt(num)
	numFloat := float64(num)
	define := float64(integer)
	var decimal float64 = 1
	for precision > 0 {
		decimal = decimal * 0.1
		define = decimalSqrt(numFloat, define, decimal)
		if define == numFloat {
			return define
		}
		precision--
	}
	return define
}

// IntegerSqrt 求整数部分的开方值
func IntegerSqrt(x int) int {
	l, r := 0, x
	result := -1
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			result = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return result
}

// decimalSqrt 求小数部分的开方值
func decimalSqrt(x float64, define float64, decimal float64) float64 {
	var l, r, result float64
	l = define
	r = define + decimal*9
	result = -1
	for l <= r {
		mid := l + (r-l)/2.0
		mid = TruncateNaive(mid, decimal) // 截断，只考虑当前小数位
		if mid*mid <= x {
			result = mid
			l = mid + decimal
		} else {
			r = mid - decimal
		}
	}
	decimal = decimal * 0.1
	return result
}

// TruncateNaive 截断float，不会四舍五入
// uint 代表小数位数，格式位 0.000001 如果是几位就指定为几位
func TruncateNaive(f float64, unit float64) float64 {
	return math.Trunc(f/unit) * unit
}
