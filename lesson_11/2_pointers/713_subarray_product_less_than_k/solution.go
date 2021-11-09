package subarray_product_less_than_k

/*
	https://leetcode.com/problems/subarray-product-less-than-k

   2 pointers f & s at index 0
   iterate f & s to the end of nums
   f != 0 && nums[f] < k => res++
   subarray's product < k
       res++
       f++
   ortherwise
       s++

   return resuul

   the complexity: O(n)
*/

func numSubarrayProductLessThanK(nums []int, k int) int {
	res, s := 0, 0
	if k < 2 {
		return 0
	}
	for f := 0; f < len(nums); f++ {
		if nums[f] < k {
			prod := getProduct(nums[s : f+1])
			//fmt.Println(prod) //76,8,75,22
			for prod >= k {
				s++
				prod = getProduct(nums[s : f+1])
			}
			res += f - s + 1
		}
	}
	return res
}

func getProduct(nums []int) int {
	r := 1
	for _, n := range nums {
		r *= n
	}
	return r
}
