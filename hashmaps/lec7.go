/*
Link to question: https://www.naukri.com/code360/problems/length-of-the-largest-subarray_2825087

Given an array(arr) of integers. Values may be duplicated.
You have to find the length of the largest subarray with contiguous elements.
Note - The contiguous elements can be in any order(not necessarily in increasing order).


GO does not have built-in hashSet to put elements in set, 
so developers can use hashmap and set value to true, inorder to use it as set. 

Intution: 
(1) max-min == i-j , then we have found subarray with continuous elements
(2) The above intution fails , if array has duplicate elements
    Ex: 9,8,8,6 max-min == i-j ---> 9-6 == 3-0 , its wrong, 7 is not present hence need to check for duplicacy alsways.
Note: to pass the testcase, you can make totalCount := 1, to consider the single element as well. 
      and consider starting j from same element as i, to consider length = 1 always.
*/


func maxLength(nums []int)int{
	duplicacySet := make(map[int]bool)
	totalCount := 0

	minn := math.IntMin
	maxx := math.IntMax
	for i:=0;i<len(nums);i++{ //need to go till len-1, as j traverses from i+1 to end
		for j:=i; j<len(nums); j++{
			if _, exists := duplicacySet[nums[i]]; exists{
				break
			}else{
				duplicacySet[nums[i]] = true
			}
			minn = min(minn, nums[i])
			maxx = max(maxx, nums[i])
			if maxx-minn == i-j{
				totalCount += i-j+1
			}
		}
	}
	return totalCount
}