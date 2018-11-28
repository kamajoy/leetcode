package main

import (
	"fmt"
)

//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
//123456反转为654321

func main() {
	num := 123456
	rev := reverse(num)

	fmt.Println(rev)

}

var MIN int = 0x80000000
var MAX int = 0x7FFFFFFF

func reverse(x int) int {
	sum := 0
	for {
		leftDigits := x / 10
		lastDigit := x % 10
		x = leftDigits

		sum = sum*10 + lastDigit

		if 0 == leftDigits {
			break
		}
	}

	//如果溢出，返回0
	if sum < -MIN || sum > MAX {
		fmt.Println(sum)
		sum = 0
	}

	return sum
}
