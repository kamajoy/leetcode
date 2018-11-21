//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的 两个 整数。


//判断一个元素是否在切片中
func is_array(tg int,nums []int)bool{
    for _,v := range nums{
        if v == tg {
            return true
        }
    }
    
    return false
}

//一遍hash表
func twoSum(nums []int, target int) []int {
    ret := make([]int,2)
	len := len(nums)

	for i := 0; i < len; i++ {
        if is_array(nums[i],nums){
            ret = append(ret,i)
            ret = append(ret,i)
            return ret
        }
	}
    
    return ret
    
}

//暴力破解
func twoSum2(nums []int, target int) []int {
    tg := make([]int,3)
	len := len(nums)

	for i := 0; i < len-1; i++ {
		for j := i + 1; j < len; j++ {
			if nums[i]+nums[j] == target {
                tg = append(tg,i)
                tg = append(tg,j)
			}
		}
	}
    
    return tg
    
}
