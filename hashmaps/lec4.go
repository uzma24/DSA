/*
Question link: https://www.geeksforgeeks.org/problems/count-distinct-elements-in-every-window/1

Given an integer array arr[] and a number k. 
Find the count of distinct elements in every window of size k in the array.

Example: 
	Input: arr[] = [1, 2, 1, 3, 4, 2, 3], k = 4
	Output: [3, 4, 4, 3]
	Explanation:
		First window is [1, 2, 1, 3], count of distinct numbers is 3.
		Second window is [2, 1, 3, 4], count of distinct numbers is 4.
		Third window is [1, 3, 4, 2], count of distinct numbers is 4.
		Fourth window is [3, 4, 2, 3], count of distinct numbers is 3.

	Input: arr[] = [1, 1, 1, 1, 1], k = 3
	Output: [1, 1, 1]
	Explanation: Every window of size 3 in the array [1, 1, 1, 1, 1], contains only the element 1, 
		so the number of distinct elements in each window is 1.

*/
func countDistinct(nums []int, k int) []int {
	n := len(nums)
	distinctCountResult := make([]int, 0)
	freqMap := make(map[int]int, k)

	for i:=0;i<k;i++{
		freqMap[nums[i]]++
	}
	append(distinctCountResult, len(freqMap)) //intilises distinctCountResult with first k distinct element
	

	start := 0
	end := k-1
	for end-start < k && start < n-k+1, end < n-1{
		releaseAtStart(freqMap, start)
		acquireAtEnd(nums, freqMap, end)
		append(distinctCountResult, len(freqMap))
		start++
		end++
	}

	return distinctCountResult
}

func releaseAtStart(freqMap map[int]int, start int){
	delete(freqMap, start)
}

func acquireAtEnd(nums []int, freqMap map[int]int, start int){
	freqMap[nums[end]]++
}