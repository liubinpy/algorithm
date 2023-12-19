package basic

import (
    "math/rand"
)

//问：有一个函数f1()可以获取到随机数3-19，在不使用其他的系统内置的随机函数之前，获取20-56的随机数。

func f1() int {
    return rand.Intn(17) + 3
}

// f2 0,1等概率随机数生成器
func f2() int {
    n := f1()
    for n == 11 {
        n = f1()
    }
    if n < 11 {
        return 0
    } else {
        return 1
    }
}

// f3 0-64等概率随机数生成器
func f3() int {
    return f2()<<5 + f2()<<4 + f2()<<3 + f2()<<2 + f2()<<1 + f2()
}

// f4 0-36等概率随机数生成器
func f4() int {
    n := f3()
    for n > 36 {
        n = f3()
    }
    return n
}

// f5 得到20-56的等概率随机数生成器
func f5() int {
    return f4() + 20
}
