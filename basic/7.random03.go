package basic

import "math/rand"

// 问：你知道p1会以固定的概率返回0和1，但是p1的内容你看不到，你需要根据这个函数，写一个函数等概率返回0和1

func p1() int {
    if rand.Float64() < 0.84 {
        return 1
    } else {
        return 0
    }
}

// p2 等概率返回0和1
func p2() int {
    n := p1()
    m := p1()
    for (n == 1 && m == 1) || (n == 0 && m == 0) {
        n = p1()
        m = p1()
    }

    if n == 0 {
        return 0
    } else {
        return 1
    }
}

// 0 0 # 都返回0的概率为 a * a
// 1 1 # 都返回1的概率是 (1-a) * (1-a)
// 1 0 # 返回10的概率是 (1-a) * a
// 0 1 # 返回01的概率是 a * (1-a)
// 让返回00和11的都重做，这样返回10和01的概率都相同
