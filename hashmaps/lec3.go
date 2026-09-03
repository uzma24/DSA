/*
Question link: https://leetcode.com/problems/check-if-array-pairs-are-divisible-by-k/description/

(a) arr of integers and a number k, return true if you can divide the arr into n/2 pairs and each pair is divisible by k. 
(b) number can be -ve or +ve
(c) length of arr is even


Input: arr = [1,2,3,4,5,10,6,7,8,9], k = 5
Output: true
Explanation: Pairs are (1,9), (2,8), (3,7), (4,6), and (5,10).

Input: arr = [1,2,3,4,5,6], k = 7
Output: true
Explanation: Pairs are (1,6), (2,5), and (3,4).

Input: arr = [1,2,3,4,5,6], k = 10
Output: false
Explanation: You can try all possible pairs, but there is no way to divide arr into 3 pairs such that each pair’s sum is divisible by 10.



intutions: 
(a) find freqMap for all rem 
(b) for rem = 0 ---> even
    for rem = k/2 ---> even
    for each rem = x ------> (k-x) needs to be present in map
(c) prevent double count
    - when rem = 0, if count of 0 = 2, pair count is 1
    - when rem = k/2,count of k/2 = 2, pair count is 1
    - when rem = x, there has to be other k-x remainder, hence if we count for x = 1 pair, means we already counted its counter part k-x in that 1 pair, hence do not count k-x again. 
*/

func canArrange(arr []int, k int) bool {
    freqMap := make(map[int]int)
    expectedPairCount := len(arr)/2


    for i:=0;i<len(arr);i++{
        rem := (arr[i]%k + k )%k
        freqMap[rem]++ //creates reaminder frequency-map for each element when divisible by k
    }

    pairCount := 0
    for numKey, freqVal:= range freqMap{
        if numKey == 0 && isEven(freqVal){
            pairCount += freqVal/2 //prevent double count
        }else if numKey == k/2 && isEven(freqVal){
            pairCount += freqVal/2 //prevent double count
        }else if numKey < k-numKey{ //prevent double count
            if freqMap[k-numKey] == freqVal{
                pairCount += freqVal
            }
        }
    }

    if pairCount == expectedPairCount{
        return true
    }

    return false
}

func isEven(num int) bool{
    if num%2 == 0{
        return true
    }
    return false
}