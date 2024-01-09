package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	a := 200
	b := &a
	*b++
	c := &b
	**c++ // указатель на указатель
	fmt.Println(a)

}

func (b bill) format() string {
	fs := "Bill breakdown:\n"
	var total float64 = 0
	for k, v := range b.items {
		fs += fmt.Sprintf("%v ...$%v", k+":", v)
		total += v
	}
	fs += fmt.Sprintf("%v ...$%0.2f total:", total)
	return fs
}

func isFascinating(n int) bool {
	n = n*1000000 + (n*2)*1000 + (n * 3)
	sum := strconv.Itoa(n)
	byteSlice := []rune(sum)
	for i := 0; i < len(byteSlice); i++ {
		for j := i + 1; j < len(byteSlice); j++ {
			if byteSlice[i] == byteSlice[j] {
				return false
			}
		}
	}
	return true
}

func theMaximumAchievableX(num int, t int) int {
	return num + t*2
}

func divide(dividend int, divisor int) int {
	dividend /= divisor
	if dividend > math.MaxInt32 {
		return math.MaxInt32
	} else if dividend < math.MinInt32 {
		return math.MinInt32
	} else {
		return dividend
	}
}

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func removeElement(nums []int, val int) int {
	temp := 0
	var a []int
	for _, value := range nums {
		if value != val {
			temp++
			a = append(a, value)
		}
	}
	nums = a
	return temp
}

func maxArea(height []int) int {
	var max int
	if height[0] < height[1] {
		max = height[0]
	} else {
		max = height[1]
	}
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			if height[i] < height[j] {
				if height[i]*(j-i) > max {
					max = height[i] * (j - i)
				}
			} else {
				if height[j]*(j-i) > max {
					max = height[j] * (j - i)
				}
			}
		}
	}
	return max
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
