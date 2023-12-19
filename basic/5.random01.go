package basic

import "math/rand"

//问：有一个函数f()可以获取到随机数1-5，在不使用其他的系统内置的随机函数之前，获取1-7的随机数。

func f() int {
    return rand.Intn(5) + 1
}

// zeroOneGenerator 使用f()函数生成一个0和1的等概率发生器
func zeroOneGenerator() int {
    i := f()
    for i == 3 {
        i = f()
    }

    if i < 3 {
        return 0
    } else {
        return 1
    }

}

// random0to7 得到0-7的等概率随机数生成器
func random0to7() int {
    return zeroOneGenerator()<<2 + zeroOneGenerator()<<1 + zeroOneGenerator()
}

// random0to6 得到0-6的等概率随机数生成器
func random0to6() int {
    n := random0to7()
    for n == 7 {
        n = random0to7()
    }
    return n
}

// random1to7 最终获取1-7的等概率随机数生成器
func random1to7() int {
    return random0to6() + 1
}
