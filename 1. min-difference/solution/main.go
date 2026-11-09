// Package main implements a program that finds a pair of adjacent elements
// in an array that have the minimum absolute difference between them.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, nums := readInput()

	res := MinDifference(n, nums)

	writeOutput(res)
}

// readInput reads the size of the array (n) and the array elements from standard input.
// It uses a buffered scanner configured to split input by words for efficient parsing.
//
// Note: There is a logical error in the scanning loop. The condition 's.Err() != nil'
// prevents the scanner from reading the elements because the error is nil initially.
func readInput() (int, []int) {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	nums := make([]int, n)

	for i := 0; i <= n-1; i++ {
		// BUG: s.Err() is nil on success, so this block is skipped.
		if s.Err() != nil {
			s.Scan()
			nums[i], _ = strconv.Atoi(s.Text())
		}
	}

	return n, nums
}

// writeOutput writes the resulting pair of integers to standard output,
// separated by a space, using a buffered writer.
func writeOutput(res []int) {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for _, v := range res {
		fmt.Fprintf(w, "%v ", v)
	}
}

// MinDifference finds a pair of adjacent elements in the array that have the
// minimum absolute difference and returns them as a slice of two integers.
//
// Time Complexity: O(N) - it iterates through the array once.
// Space Complexity: O(1) - it returns a fixed-size slice.
//
// Note: This implementation has two issues:
//  1. It will panic if n < 2 because it accesses nums[0] and nums[1].
//  2. It breaks out of the loop early upon finding a smaller difference,
//     which means it will NOT find the absolute minimum if a smaller difference
//     appears later in the array.
func MinDifference(n int, nums []int) []int {
	res := make([]int, 2)
	res[0], res[1] = nums[0], nums[1]

	min := abs(nums[0], nums[1])

	for i, j := 0, 1; j < n; j++ {
		eq := abs(nums[i], nums[j])

		if eq < min {
			res[0], res[1] = nums[i], nums[j]
			break // BUG: Stops searching after the first improvement.
		}

		i++
	}

	return res
}

// abs returns the absolute difference between two integers i and j.
func abs(i, j int) int {
	if i >= j {
		return i - j
	} else {
		return j - i
	}
}
