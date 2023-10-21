package main

import "fmt"

func BinarySearch(arr []int, num int) int {
	return search(arr, 0, len(arr), num)
}

func search(arr []int, lo int, hi int, num int) int {
	mid := lo + (hi-lo)/2
	val := arr[mid]
	if lo > hi {
		return -1
	}
	if val == num {
		return mid
	} else if val > mid {
		return search(arr, mid+1, hi, num)
	} else {
		return search(arr, lo, mid, num)
	}
}

func BubbleSort(arr *[]int) {
	for i := 0; i < len(*arr)-1; i++ {
		for j := i; j < len(*arr); j++ {
			curr := (*arr)[i]
			next := (*arr)[i+1]
			if curr > next {
				(*arr)[i] = next
				(*arr)[i+1] = curr
			}
		}
	}
}
