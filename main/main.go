package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"unicode"
)

func main() {

	fmt.Println(myAtoi("42"))

	//fmt.Println(isPalindrome(122))

	//fmt.Println(reverse(1534236469))

	//var nums []int = []int{2, 7, 11, 15}
	//fmt.Println(twoSum(nums, 9))

	//var nums []int = []int{0, 1, 0, 3, 2, 3}
	//fmt.Println(lengthOfLIS(nums))

	//	var nums1 []int = []int{1, 3}
	//	var nums2 []int = []int{2}
	//	fmt.Println(findMedianSortedArrays(nums1, nums2))

}

func myAtoi(s string) int {
	num := 0
	for i := range s {
		if unicode.IsDigit(rune(i)) {
			fmt.Println(i)
			n, _ := strconv.Atoi(string(rune(i)))
			num = num*10 + n
		}
	}
	return num
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	current := 0
	temp := x
	for x > 0 {
		current = current*10 + x%10
		x /= 10
	}
	if current == temp {
		return true
	} else {
		return false
	}
}

func reverse(x int) int {
	current := 0
	if x >= 0 {
		for x > 0 {
			current = current*10 + x%10
			x /= 10
		}
	} else {
		x *= -1
		for x > 0 {
			current = current*10 + x%10
			x /= 10
		}
		current *= -1
	}
	if current < math.MinInt32 || current > math.MaxInt32 {
		return 0
	}
	return current
}

func twoSum(nums []int, target int) []int {
	var newNums []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				newNums = append(newNums, i)
				newNums = append(newNums, j)
				return newNums
			}
		}
	}
	newNums = append(newNums, nums[len(nums)-2])
	newNums = append(newNums, nums[len(nums)-1])
	return newNums
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	for i := 0; i < len(nums2); i++ {
		nums1 = append(nums1, nums2[i])
	}
	sort.Ints(nums1)
	fmt.Println(nums1)
	if len(nums1)%2 == 0 {
		return float64(nums1[len(nums1)/2-1]+nums1[len(nums1)/2]) / 2
	} else {
		return float64(nums1[len(nums1)/2])
	}
}

func lengthOfLIS(nums []int) int {
	min := nums[0]
	index := 0
	for i, value := range nums {
		if value < min {
			min = value
			index = i
		}
	}
	var newNums []int
	newNums = append(newNums, nums[index])
	j := 0
	for i := index; i < len(nums); i++ {
		if newNums[j] < nums[i] {
			newNums = append(newNums, nums[i])
			j++
		}
	}
	return len(newNums)
}
