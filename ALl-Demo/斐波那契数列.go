package main

import (
	"fmt"
)

//斐波那契数列-递归 1、1、2、3、5、8、13、21、34

func main() {
	n := 8
	fmt.Println(fib(n))

	fmt.Println(fibFoo(n))
}

//递归，会有严重的性能问题，递归的最大弊端是对已经计算的不能重复利用
func fib(n int) int {
	if n < 3 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

//循环-比较优
func fibFoo(n int) int {
	arr := make(map[int]int)
	for i := 1; i <= n; i++ {
		if i < 3 {
			arr[i] = 1
		} else {
			arr[i] = arr[i-1] + arr[i-2]
		}
	}

	return arr[n]
}
