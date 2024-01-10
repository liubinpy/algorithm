package basic

import "math"

// 位运算实现 加法、减法、乘法、除法，不使用+,-,*和/

// Add 加法
// 无进位相加(两个数进行异或运算) + 进位信息(两个数进行与运算，左移动一位)，一直计算，直到进位信息为0，则返回无进位相加的值
func Add(a, b int) int {
	sum := a ^ b
	carry := (a & b) << 1
	for carry != 0 {
		preSum := sum
		sum = sum ^ carry
		carry = (preSum & carry) << 1
	}
	return sum
}

// negNum 相反数
func negNum(n int) int {
	return Add(^n, 1)
}

// Sub 减法
// a - b 就是 a + b的相反数，b的相反数就是(b取反+1)
func Sub(a, b int) int {
	return Add(a, negNum(b))
}

// Multi 乘法
func Multi(a, b int) int {
	var sum int
	for b != 0 {
		if (b & 1) != 0 {
			sum = Add(sum, a)
		}
		a <<= 1
		b >>= 1
	}
	return sum
}

func IsNeg(n int) bool {
	return n < 0
}

// Div 除法
func Div(dividend int, divisor int) int {
	// 先处理特殊情况
	if divisor == math.MinInt32 && dividend == math.MinInt32 {
		return 1
	} else if divisor == math.MinInt32 {
		return 0
	} else if dividend == math.MinInt32 {
		if divisor == negNum(1) {
			return math.MaxInt32
		} else {
			res := div(Add(dividend, 1), divisor)
			return Add(res, div(Sub(dividend, Multi(res, divisor)), divisor))
		}
	}
	return div(dividend, divisor)
}

func div(a int, b int) int {
	x := a
	y := b
	if IsNeg(a) {
		x = negNum(a)
	}
	if IsNeg(b) {
		y = negNum(b)
	}
	var res int
	var i int = 30
	var flag int = 1
	for ; i >= 0; i = Sub(i, 1) {
		if (x >> i) >= y {
			res = res | (flag << i)
			x = Sub(x, y<<i)
		}
	}
	if IsNeg(a) && IsNeg(b) {
		return res
	} else if (!IsNeg(a)) && (!IsNeg(b)) {
		return res
	} else {
		return negNum(res)
	}
}

// 除法 https://leetcode.cn/problems/divide-two-integers/ 要求 不使用 乘法、除法和取余运算。
func divide(dividend int, divisor int) int {
	// 先处理特殊情况
	if divisor == math.MinInt32 && dividend == math.MinInt32 {
		return 1
	} else if divisor == math.MinInt32 {
		return 0
	} else if dividend == math.MinInt32 {
		if divisor == -1 {
			return math.MaxInt32
		} else {
			res := div1(dividend+1, divisor)
			return res + div1(dividend-res*divisor, divisor)
		}
	}
	return div1(dividend, divisor)
}

func div1(a int, b int) int {
	x := a
	y := b
	if a < 0 {
		x = ^a + 1
	}
	if b < 0 {
		y = ^b + 1
	}
	var res int
	var i int = 30
	var flag int = 1
	for ; i >= 0; i-- {
		if (x >> i) >= y {
			res = res | (flag << i)
			x -= y << i
		}
	}
	if a < 0 && b < 0 {
		return res
	} else if a > 0 && b > 0 {
		return res
	} else {
		return ^res + 1
	}
}
