package main

import (
	"time"
)

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
					break
				}
				if len(sub_right) == 0 {
					result = append(result, sub_left...)
					permutations += len(sub_left)
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

func twoPhaseMergeSort(array []int) (int, int, int64, bool) {
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
		var cmp, prmt int
		cmp, prmt, array = twoPhaseMerge(b, c, step)
		comparisons += cmp
		permutations += prmt
		step *= 2
	}

	end := time.Now()

	return comparisons, permutations, int64(end.Sub(start).Nanoseconds()), isSorted(array)
}

func onePhaseMerge(left, right []int, step int) ([]int, []int, int, int) { // {{{
	b, c := make([]int, 0), make([]int, 0)

	permutations, comparisons := 0, 0

	counter := 0
	for len(left) > 0 || len(right) > 0 {
		sub_left, sub_right := left[:min(step, len(left))], right[:min(step, len(right))]
		left, right = left[min(step, len(left)):], right[min(step, len(right)):]

		if counter&1 == 0 {
			for i := 0; i < step*2; i++ {
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

				permutations++
				comparisons++
			}
		} else {
			for i := 0; i < step*2; i++ {
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

				permutations++
				comparisons++
			}
		}

		counter++
	}

	return b, c, comparisons, permutations
} // }}}

func onePhaseMergeSort(array []int) (int, int, int64, bool) { // {{{
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

	end := time.Now()

	return comparisons, permutations, int64(end.Sub(start).Nanoseconds()), isSorted(array)
} // }}}

func twoPhaseNaturalMergeSort(array []int) (int, int, int64) {
	start := time.Now()
	end := time.Now()
	return 0, 0, end.Sub(start).Nanoseconds()
}

func onePhaseNaturalMergeSort(array []int) (int, int, int64) {
	start := time.Now()
	end := time.Now()
	return 0, 0, end.Sub(start).Nanoseconds()
}

func absorptionSort(array []int) (int, int, int64) {
	start := time.Now()
	end := time.Now()
	return 0, 0, end.Sub(start).Nanoseconds()
}
