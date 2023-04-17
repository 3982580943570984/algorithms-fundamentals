package main

import (
	"math/rand"
	"time"
)

/* Утилиты */
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func GenerateRandomNumbers(n int) []int {
	result := make([]int, n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result[i] = r.Intn(n)
	}
	return result
}

func IsSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}
