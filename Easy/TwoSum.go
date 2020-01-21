/*给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
示例:
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]*/
package main

import "fmt"

func main() {

	arr := TwoSumOneHash()
	//arr := TwoSumTwoHash()
	//arr := TwoSum()
	fmt.Println(arr)
}

//一遍hash 时间复杂度：O(n)
func TwoSumOneHash() []int {
	target := 9
	nums := [...]int{1, 7, 11, 15, 3, 5, 6}

	len := len(nums)
	hashList := map[int]int{}
	for i := 0; i < len; i++ {
		comp := target - nums[i]
		if _, err := hashList[comp]; err {
			if hashList[comp] != i {
				return []int{hashList[comp], i}
			}
		} else {
			hashList[nums[i]] = i
		}
	}
	return []int{}
}

//两遍hash，时间复杂度：O(n)
func TwoSumTwoHash() []int {
	target := 9
	nums := [...]int{2, 7, 11, 15}

	hashList := map[int]int{}
	for k, v := range nums {
		hashList[v] = k
	}

	len := len(nums)
	for i := 0; i < len; i++ {
		comp := target - nums[i]
		if _, err := hashList[comp]; err {
			if hashList[comp] != i {
				return []int{i, hashList[comp]}
			}
		}
	}
	return []int{}
}

//暴力破解,时间复杂度 O(n)^2
func TwoSum() []int {
	target := 9
	nums := [...]int{2, 7, 11, 15}

	var x, y int = 0, 0
	len := len(nums)
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			if nums[i]+nums[j] == target {
				x, y = i, j
				break
			}
		}
	}
	return []int{x, y}
}
