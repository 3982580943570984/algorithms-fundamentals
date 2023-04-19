package main

import (
	"math"
	"sort"
	"time"
)

/* Простая сортировка двухфазным слиянием */
func TwoPhaseMerge(left, right []int, step int) (int, int, []int) {
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

func TwoPhaseMergeSort(array []int) (int, int, int64, bool) {
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
		comparisons, permutations, array = TwoPhaseMerge(b, c, step)
		step *= 2
	}

	return comparisons, permutations, int64(time.Since(start)), IsSorted(array)
}

/* Простая сортировка однофазным слиянием */
func OnePhaseMerge(left, right []int, step int) ([]int, []int, int, int) {
	b, c := make([]int, 0), make([]int, 0)

	permutations, comparisons := 0, 0

	counter := 0
	for len(left) > 0 || len(right) > 0 {
		sub_left, sub_right := left[:min(step, len(left))], right[:min(step, len(right))]
		left, right = left[min(step, len(left)):], right[min(step, len(right)):]

		// Четное -> в 'b', Нечетное -> в 'c'
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

func OnePhaseMergeSort(array []int) (int, int, int64, bool) {
	start := time.Now()

	comparisons, permutations := 0, 0

	// Определяем четные элементы в массив 'b', нечетные - в массив 'c'
	b, c := []int{}, []int{}
	for i := 0; i < len(array); i++ {
		if i&1 == 0 {
			b = append(b, array[i])
		} else {
			c = append(c, array[i])
		}
		permutations++
		comparisons++
	}

	step := 1
	for step <= len(array) {
		cmps, prms := 0, 0
		b, c, cmps, prms = OnePhaseMerge(b, c, step)
		comparisons += cmps
		permutations += prms
		step *= 2
	}

	array = b
	array = append(array, c...)

	return comparisons, permutations, int64(time.Since(start)), IsSorted(array)
}

/* Естественная сортировка двухфазным слиянием */
func NaturalTwoPhaseSplit(items []int) ([][]int, [][]int) {
	b, c := [][]int{}, [][]int{}

	counter := 0
	currentSplit := []int{items[0]}
	for i := 1; i < len(items); i++ {
		if items[i] >= items[i-1] {
			currentSplit = append(currentSplit, items[i])
		} else {
			if counter&1 == 0 {
				b = append(b, currentSplit)
			} else {
				c = append(c, currentSplit)
			}
			counter++

			currentSplit = []int{items[i]}
		}
	}

	if counter&1 == 0 {
		b = append(b, currentSplit)
	} else {
		c = append(c, currentSplit)
	}

	return b, c
}

func NaturalTwoPhaseMergeSort(array []int) (int, int, int64, bool) {
	start := time.Now()
	comparisons, permutations := 0, 0

	merge := func(left, right []int) []int {
		result := []int{}
		permutations += len(left) + len(right)
		for len(left) > 0 && len(right) > 0 {
			if left[0] <= right[0] {
				result = append(result, left[0])
				left = left[1:]
			} else {
				result = append(result, right[0])
				right = right[1:]
			}
			comparisons++
		}
		result = append(result, left...)
		result = append(result, right...)
		return result
	}

	if len(array) <= 1 {
		return comparisons, permutations, int64(time.Since(start)), IsSorted(array)
	}

	// Разбиваем массив на отсортированные массивы
	b, c := NaturalTwoPhaseSplit(array)
	for len(b) > 1 {
		array = []int{}
		i := 0
		for ; i < min(len(b), len(c)); i++ {
			comparisons++
			permutations += len(b[i]) + len(c[i])
			array = append(array, merge(b[i], c[i])...)
		}

		if i < len(b) {
			array = append(array, b[i]...)
		} else if i < len(c) {
			array = append(array, c[i]...)
		}

		b, c = NaturalTwoPhaseSplit(array)
	}

	return comparisons, permutations, int64(time.Since(start)), IsSorted(b[0])
}

/* Естественная сортировка однофазным слиянием */
func NaturalOnePhaseSplit(items []int) [][]int {
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

func NaturalOnePhaseMergeSort(array []int) (int, int, int64, bool) {
	start := time.Now()
	comparisons, permutations := 0, 0

	merge := func(left, right []int) []int {
		result := []int{}
		permutations += len(left) + len(right)
		for len(left) > 0 && len(right) > 0 {
			if left[0] <= right[0] {
				result = append(result, left[0])
				left = left[1:]
			} else {
				result = append(result, right[0])
				right = right[1:]
			}
			comparisons++
		}
		result = append(result, left...)
		result = append(result, right...)
		return result
	}

	if len(array) <= 1 {
		return comparisons, permutations, int64(time.Since(start)), IsSorted(array)
	}

	splits := NaturalOnePhaseSplit(array)
	for len(splits) > 1 {
		newSplits := [][]int{}
		for i := 0; i < len(splits); i += 2 {
			if i == len(splits)-1 {
				permutations++
				newSplits = append(newSplits, splits[i])
			} else {
				permutations += len(splits[i]) + len(splits[i+1])
				newSplits = append(newSplits, merge(splits[i], splits[i+1]))
			}
			comparisons++
		}

		splits = newSplits
	}

	return comparisons, permutations, int64(time.Since(start)), IsSorted(splits[0])
}

/* Сортировка поглощением */
func MergeInsertionSort(array []int, blockSize int) (int, int, int64, bool) {
	start := time.Now()
	comparisons, permutations := 0, 0

	// Функция слияния
	merge := func(subArray []int, start int) {
		i, j := 0, start+len(subArray)
		for i < len(subArray) && j < len(array) {
			if subArray[i] < array[j] {
				array[start] = subArray[i]
				i++
			} else {
				array[start] = array[j]
				j++
			}
			start++
			comparisons++
			permutations++
		}

		for i < len(subArray) {
			array[start] = subArray[i]
			i++
			start++
			permutations++
		}
	}

	// Сортируем последний блок данных
	lastBlock := array[len(array)-blockSize:]
	sort.Slice(lastBlock, func(i, j int) bool { return lastBlock[i] < lastBlock[j] })
	permutations += len(lastBlock)

	// Перезаписываем последний блок
	copy(array[len(array)-blockSize:], lastBlock)
	permutations += len(lastBlock)

	// Определяем количество блоков
	blocksAmount := int(math.Ceil(float64(len(array))/float64(blockSize))) - 2

	// Постоянный массив для блоков
	block := make([]int, blockSize)
	for i := blocksAmount; i >= 0; i-- {
		// Копируем элементы в массив для блоков
		copy(block, array[i*blockSize:(i+1)*blockSize])
		// Сортируем блок
		sort.Ints(block)
		// Производим слияние с основным массивом
		merge(block, i*blockSize)

		permutations += len(block)
	}

	return comparisons, permutations, int64(time.Since(start)), IsSorted(array)
}
