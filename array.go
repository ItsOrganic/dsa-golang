package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{0, 5, 0, 6, 8, 9}

	arr1 := []int{1, 5, 9, 12, 15}
	arr2 := []int{9, 12, 15, 27, 52}
	fmt.Print(arr)
	rev := reverse(arr)
	fmt.Println(rev)
	fmt.Println("Second largest number in an array :", secondLargest(arr))
	fmt.Println("Checking if the array is Pallindrome or not :", Pallindrome(arr))
	fmt.Println("Create a program that moves all zeroes in an array to the end, preserving the order of non-zero elements. solve this using golang", zeroLast(arr))
	fmt.Println("Design a function to find the common elements in two sorted arrays:", commonElement(arr1, arr2))
}
func reverse(arr []int) []int {
	n := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		n[len(arr)-1-i] = arr[i]
	}
	return n
}

func secondLargest(arr []int) int {
	first := arr[0]
	secondLarge := math.MinInt
	for _, value := range arr {
		if value > first {
			secondLarge = first
			first = value
		} else if value < secondLarge && value != first {
			secondLarge = value
		}
	}
	return secondLarge
}

// Unoptimiezed is commented
func Pallindrome(arr []int) bool {
	length := len(arr) - 1
	for i := 0; i < len(arr); i++ {
		if arr[i] != arr[length-i] {
			return false
		}
	}
	return true
}
func zeroLast(arr []int) []int {
	read, write := 0, 0
	for read < len(arr) {
		if arr[read] != 0 {
			arr[write], arr[read] = arr[read], arr[write]
			write++
		}
		read++
	}
	return arr
}
func commonElement(arr1 []int, arr2 []int) []int {
	result := []int{}
	hashMap := make(map[int]bool)

	for _, value := range arr1 {
		hashMap[value] = true
	}
	for _, value := range arr2 {
		if hashMap[value] {
			result = append(result, value)
			hashMap[value] = false
		}
	}
	return result
}
