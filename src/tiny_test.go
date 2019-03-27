package src

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

func Test_tiny(t *testing.T) {
	fmt.Println(5 &^ 4) //与非！！卧草！！就是去除1的而已！！我日！装尼玛！
	fmt.Println(round(7777, 512))
}

func round(n, a uintptr) uintptr { //使得n为a的倍数！
	return (n + a - 1) &^ (a - 1)
}

func roundm(n, a uintptr) uintptr {
	return (n + a - 1) &^ (a - 1) //n + a -1 ，即如果 n & (a - 1)  不等于0 n 的 a 位必须 加1  来达到 a的倍数的效果!!!向上进 a 一位。 然后地位全部清零!!>> 进位 清零
}

const size = 5 * 1024 * 1024 * 1024 / 8

func test() {
	fmt.Println(os.Getpid())
	var a [size]int //5g
	for i := 0; i < len(a); i++ {
		a[i] = i
	}
	for {
		a[rand.Intn(size)] = a[2] //减少被置换!!
	}
}
