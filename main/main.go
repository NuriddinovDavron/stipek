package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	for i := 0; i < len(nums2); i++ {
		nums1[len(nums1)+i] = nums2[i]
	}
	if len(nums1)%2 == 0 {
		return float64(nums1[len(nums1)/2-1]+nums1[len(nums1)/2]) / 2
	} else {
		return float64(nums1[len(nums1)/2])
	}
}
