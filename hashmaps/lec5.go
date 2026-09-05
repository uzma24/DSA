/*
Link to question: https://www.geeksforgeeks.org/problems/largest-subarray-with-0-sum/1 

Given an array arr[] containing both positive and negative integers, the task is to find the length of the longest subarray with a sum equals to 0.

Example
Input: arr[] = [15, -2, 2, -8, 1, 7, 10, 23]
Output: 5
Explanation: The longest subarray with sum equals to 0 is [-2, 2, -8, 1, 7].

Input: arr[] = [2, 10, 4]
Output: 0
Explanation: There is no subarray with a sum of 0.

Input: arr[] = [1, 0, -4, 3, 1, 0]
Output: 5
Explanation: The longest subarray with sum equals to 0 is [0, -4, 3, 1, 0]
*/


func maxLength(nums []int) int{
	cummulativeSumIndexMap := make(map[int]int)
	resLargestSubarrayLen := 0

	cummulativeSumIndexMap[0] = -1 //sum at index -1 = 0
	for i:=0;i<len(nums);i++{
		sum := sum + nums[i]

		if !exists := cummulativeSumIndexMap[sum]{
			cummulativeSumIndexMap[sum]= i
		}else{
			resLargestSubarrayLen = max(resLargestSubarrayLen, i - cummulativeSumIndexMap[sum])
		}
	}
	return resLargestSubarrayLen
}