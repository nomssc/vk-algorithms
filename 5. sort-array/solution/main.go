// Package main implements a program that reads an array of integers,
// sorts it using the QuickSort algorithm, and prints the sorted result.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// IO is a wrapper around bufio.Scanner and bufio.Writer to provide fast I/O.
// Using buffered I/O prevents Time Limit Exceeded (TLE) errors on large inputs.
type IO struct {
	s *bufio.Scanner
	w *bufio.Writer
}

// NewIO initializes and returns a new IO instance.
func NewIO() *IO {
	s := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	// Split input by words (spaces, newlines) for easy integer parsing.
	s.Split(bufio.ScanWords)

	return &IO{
		s: s,
		w: w,
	}
}

// Close flushes any buffered output to standard output.
// Must be called before the program exits.
func (o *IO) Close() error {
	return o.w.Flush()
}

// ReadConditions reads the array size and elements from standard input.
// It returns a slice of integers containing the elements.
func (o *IO) ReadConditions() []int {
	o.s.Scan()
	n, _ := strconv.Atoi(o.s.Text())

	nums := make([]int, n)

	for i := 0; i < n; i++ {
		o.s.Scan()
		nums[i], _ = strconv.Atoi(o.s.Text())
	}

	return nums
}

// WriteResult prints the sorted slice elements, separated by spaces.
func (o *IO) WriteResult(res []int) {
	for _, v := range res {
		fmt.Fprintf(o.w, "%v ", v)
	}
}

func main() {
	o := NewIO()
	defer o.Close()

	nums := o.ReadConditions()

	// Sort the array in-place
	QuickSort(0, len(nums)-1, nums)

	o.WriteResult(nums)
}

// QuickSort sorts the given slice of integers using the Divide and Conquer strategy.
// It uses the Lomuto partition scheme, choosing the last element as the pivot.
//
// Time Complexity:
//   - Average/Best case: O(N log N)
//   - Worst case: O(N^2) - Occurs when the array is already sorted or reverse sorted.
//
// Space Complexity:
//   - Average case: O(log N) due to recursion stack.
//   - Worst case: O(N) due to unbalanced recursion tree.
func QuickSort(low int, high int, nums []int) {
	// Base case: subarray of size 1 (or 0) is already sorted
	if low >= high {
		return
	}

	// Choose the rightmost element as the pivot
	pivot := nums[high]

	// i acts as a boundary pointer for elements smaller than the pivot.
	// Everything before index 'i' will be strictly less than 'pivot'.
	i, j := low, low

	// Partitioning loop: j scans through the subarray
	for j < high {
		// If the current element is smaller than the pivot,
		// swap it to the left side of the boundary and expand the boundary.
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}

	// Place the pivot in its correct final position
	// (right after all elements smaller than it).
	nums[i], nums[j] = nums[j], nums[i]

	// Recursively sort the left and right partitions.
	// Note: 'i' is now the correct position of the pivot.
	QuickSort(low, i-1, nums)
	QuickSort(i+1, high, nums)
}
