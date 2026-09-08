/*
Link to question: https://www.geeksforgeeks.org/problems/zero-sum-subarrays1825/1

Given an integer array `nums`, find the total number of continuous subarrays whose elements sum up to exactly `0`. 
Note: A subarray is defined as any contiguous (ordered and unbroken) sequence of elements inside the main array.

Example 1
Input: nums = [1, 2, -3]
Output: 1
Explanation: The entire array `[1, 2, -3]` sums up to `1 + 2 + (-3) = 0`. Thus, there is only 1 valid subarray.

Example 2
Input:nums = [6, -1, -3, 4, -2, 2]
Output: `3`
Explanation: The 3 valid subarrays that sum up to 0 are:
  1. `[-1, -3, 4]` (at indices 1 to 3) -> -1 + (-3) + 4 = 0
  2. `[-2, 2]` (at indices 4 to 5) -> -2 + 2 = 0
  3. `[-1, -3, 4, -2, 2]` (at indices 1 to 5) -> -1 + (-3) + 4 + (-2) + 2 = 0

Example 3
Input: nums = [0, 0, 0]
Output:6
Explanation: Every individual element and combination sums to 0. The 6 valid subarrays are:
  - Single elements: 3 subarrays of single length.
  - Pairs: 2 [(0,0)+ (0,0)]
  - Complete array: 1 [(0,0,0)]
  - Total = 3 + 2 + 1 = 6.

*/


func find0SumSubarrayCount(nums []int) int{
	cummulativeSumIndexMap := make(map[int]int)
	zeroSumSubarrayCount := 0

	cummulativeSumIndexMap[0] = 1 //initially, freq of this sum 0 = 1
	sum := 0 
	for i:=0;i<len(nums);i++{
		sum = sum + nums[i]
		zeroSumSubarrayCount = zeroSumSubarrayCount+ cummulativeSumIndexMap[sum]
		cummulativeSumIndexMap[sum]++
	}
	return zeroSumSubarrayCount
}