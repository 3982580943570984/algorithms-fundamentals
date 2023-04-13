package main

import (
	"fmt"
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

func generateRandomNumbers(n int) []int {
	result := make([]int, n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result[i] = r.Intn(n)
	}
	return result
}

func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

/* Простая сортировка двухфазным слиянием */
func twoPhaseMerge(left, right []int, step int) (int, int, []int) {
	result := make([]int, 0, len(left)+len(right))

	permutations, comparisons := 0, 0
	for len(left) > 0 || len(right) > 0 {
		for j := 0; j < step; j++ {
			sub_left, sub_right := left[:min(step, len(left))], right[:min(step, len(right))]
			left, right = left[min(step, len(left)):], right[min(step, len(right)):]

			for len(sub_left) > 0 || len(sub_right) > 0 {
				if len(sub_left) == 0 {
					result = append(result, sub_right...)
					permutations += len(sub_right)
					sub_right = sub_right[:0]
					break
				}
				if len(sub_right) == 0 {
					result = append(result, sub_left...)
					permutations += len(sub_left)
					sub_left = sub_left[:0]
					break
				}

				if sub_left[0] < sub_right[0] {
					result = append(result, sub_left[0])
					sub_left = sub_left[1:]
				} else {
					result = append(result, sub_right[0])
					sub_right = sub_right[1:]
				}

				permutations++
				comparisons++
			}
		}
	}

	return comparisons, permutations, result
}

func twoPhaseMergeSort(array []int) (int, int, int64) {
	start := time.Now()

	comparisons, permutations := 0, 0
	step := 1
	for step <= len(array) {
		b, c := []int{}, []int{}
		for i := 0; i < len(array); i += step * 2 {
			left, right, end := i, min(i+step, len(array)), min(i+step*2, len(array))

			b = append(b, array[left:right]...)
			c = append(c, array[right:end]...)
		}
		comparisons, permutations, array = twoPhaseMerge(b, c, step)
		step *= 2
	}
	fmt.Println(array)

	end := time.Now()

	return comparisons, permutations, int64(end.Sub(start).Nanoseconds())
}

/* Простая сортировка однофазным слиянием */
func onePhaseMerge(left, right []int, step int) ([]int, []int, int, int) {
	b, c := make([]int, 0), make([]int, 0)

	permutations, comparisons := 0, 0

	counter := 0
	for len(left) > 0 || len(right) > 0 {
		sub_left, sub_right := left[:min(step, len(left))], right[:min(step, len(right))]
		left, right = left[min(step, len(left)):], right[min(step, len(right)):]

		if counter&1 == 0 {
			for i := 0; i <= step*2; i++ {
				if len(sub_left) == 0 {
					b = append(b, sub_right...)
					break
				}

				if len(sub_right) == 0 {
					b = append(b, sub_left...)
					break
				}

				if sub_left[0] < sub_right[0] {
					b = append(b, sub_left[0])
					sub_left = sub_left[1:]
				} else {
					b = append(b, sub_right[0])
					sub_right = sub_right[1:]
				}
			}
		} else {
			for i := 0; i <= step*2; i++ {
				if len(sub_left) == 0 {
					c = append(c, sub_right...)
					break
				}

				if len(sub_right) == 0 {
					c = append(c, sub_left...)
					break

				}

				if sub_left[0] < sub_right[0] {
					c = append(c, sub_left[0])
					sub_left = sub_left[1:]
				} else {
					c = append(c, sub_right[0])
					sub_right = sub_right[1:]
				}
			}
		}

		permutations++
		comparisons++

		counter++
	}

	return b, c, comparisons, permutations
}

func onePhaseMergeSort(array []int) (int, int, int64) {
	start := time.Now()

	comparisons, permutations := 0, 0

	b, c := []int{}, []int{}
	for i := 0; i < len(array); i++ {
		if i&1 != 1 {
			b = append(b, array[i])
		} else {
			c = append(c, array[i])
		}
	}

	step := 1
	for step <= len(array) {
		b, c, comparisons, permutations = onePhaseMerge(b, c, step)
		step *= 2
	}

	array = b
	array = append(array, c...)
	fmt.Println(isSorted(array))

	end := time.Now()

	return comparisons, permutations, int64(end.Sub(start).Nanoseconds())
}

/* Естественная сортировка двухфазным слиянием */
func naturalTwoPhaseSplit(items []int) [][]int {
	splits := [][]int{}
	currentSplit := []int{items[0]}

	for i := 1; i < len(items); i++ {
		if items[i] >= items[i-1] {
			currentSplit = append(currentSplit, items[i])
		} else {
			splits = append(splits, currentSplit)
			currentSplit = []int{items[i]}
		}
	}

	return append(splits, currentSplit)
}

func naturalTwoPhaseMergeSort(array []int) []int {
	merge := func(left, right []int) []int {
		result := make([]int, len(left)+len(right))
		for len(left) > 0 && len(right) > 0 {
			if left[0] <= right[0] {
				result = append(result, left[0])
				left = left[1:]
			} else {
				result = append(result, right[0])
				right = right[1:]
			}
		}
		result = append(result, left...)
		result = append(result, right...)
		return result
	}

	if len(array) <= 1 {
		return array
	}

	splits := naturalTwoPhaseSplit(array)
	for len(splits) > 1 {
		newSplits := [][]int{}
		for i := 0; i < len(splits); i += 2 {
			if i == len(splits)-1 {
				newSplits = append(newSplits, splits[i])
			} else {
				newSplits = append(newSplits, merge(splits[i], splits[i+1]))
			}
		}

		splits = newSplits
	}

	return splits[0]
}

/* Естественная сортировка однофазным слиянием */
func naturalOnePhaseMerge() {}

func naturalOnePhaseMergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	return []int{}
}

func mergeInsertionSort() {}

func main() {
	for i := 0; i < 1000; i++ {
		array := generateRandomNumbers(i)
		fmt.Println(isSorted(naturalTwoPhaseMergeSort(array)))
	}
}
