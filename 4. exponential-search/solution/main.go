// Package main implements a program that reads a sorted array and a target value,
// and uses Exponential Search to find the bounds (range) where the target
// could be located.
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

// ReadConditions reads the array size (n), the sorted array elements,
// and the target value from standard input.
func (o *IO) ReadConditions() (int, []int, int) {
	o.s.Scan()
	n, _ := strconv.Atoi(o.s.Text())

	nums := make([]int, n)

	for i := 0; i < n; i++ {
		o.s.Scan()
		nums[i], _ = strconv.Atoi(o.s.Text())
	}

	o.s.Scan()
	target, _ := strconv.Atoi(o.s.Text())

	return n, nums, target
}

// WriteResult prints the resulting bounds (slice of integers) separated by spaces.
func (o *IO) WriteResult(res []int) {
	for _, v := range res {
		fmt.Fprintf(o.w, "%v ", v)
	}
}

func main() {
	o := NewIO()
	defer o.Close()

	n, nums, target := o.ReadConditions()

	// Find the bounds using Exponential Search
	res := ExponentialSearch(n, nums, target)

	o.WriteResult(res)
}

// ExponentialSearch finds the range [left, right] within a sorted array
// where a target value could potentially exist. It does this by doubling
// the search index until it finds an element greater than the target.
//
// Time Complexity: O(log i), where i is the index of the target element.
// In the worst case, it is O(log N).
// Space Complexity: O(1) - It uses a constant amount of extra space.
func ExponentialSearch(n int, nums []int, target int) []int {
	// Edge case: Empty array
	if n == 0 {
		return []int{-1, -1}
	}

	// Edge case: The target is smaller than the very first element.
	// In a sorted array, this means the target can only be at index 0 (or doesn't exist).
	if nums[0] > target {
		return []int{0, 0}
	}

	var i int

	// Phase 1: Finding the range (Galloping)
	// Start with a step size of 1 and double it (1, 2, 4, 8, ...)
	// until we find an element that is strictly greater than the target,
	// or we reach the end of the array.
	for i = 1; i < n; i *= 2 {
		if nums[i] > target {
			break
		}
	}

	// Phase 2: Determining the bounds
	// If our doubled index 'i' exceeded the array length,
	// the target (if it exists) must be in the second half of the array.
	// The lower bound is i/2 (the last index we know was <= target),
	// and the upper bound is the end of the array.
	if i >= n {
		return []int{i / 2, n - 1}
	} else {
		// Otherwise, the target is bounded between the previous index (i/2)
		// and the current index (i).
		return []int{i / 2, i}
	}
}
