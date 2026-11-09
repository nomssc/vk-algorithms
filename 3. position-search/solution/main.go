// Package main implements a program that searches for a target value in a sorted array.
// If the target is found, it returns its index. If not, it returns the index where
// the target should be inserted to maintain the array's sorted order.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// IO provides a fast buffered input/output mechanism for reading from standard input
// and writing to standard output. It is designed to minimize execution time
// for competitive programming tasks.
type IO struct {
	s *bufio.Scanner
	w *bufio.Writer
}

// NewIO initializes and returns a new IO instance. It wraps os.Stdin and os.Stdout
// with bufio.Scanner and bufio.Writer respectively. The scanner is configured
// to split input by words (spaces, newlines).
func NewIO() *IO {
	s := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	s.Split(bufio.ScanWords)

	return &IO{
		s: s,
		w: w,
	}
}

// Close flushes the buffered writer, ensuring that all buffered output is
// written to standard output. It should typically be deferred at the start of main.
func (o *IO) Close() error {
	return o.w.Flush()
}

// ReadConditions reads the problem constraints and input data from standard input.
// It returns the size of the array (n), the slice of array elements (nums),
// and the target integer to search for (target).
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

// WriteResult writes the final integer result to the buffered writer.
func (o *IO) WriteResult(res int) {
	fmt.Fprintf(o.w, "%v", res)
}

func main() {
	o := NewIO()
	defer o.Close()

	n, nums, target := o.ReadConditions()

	res := PositionSearch(n, nums, target)

	o.WriteResult(res)
}

// PositionSearch performs a classic iterative Binary Search on a sorted array
// to find the target value.
//
// If the target is found, it returns its index. If the target is not found,
// it returns the index where the target should be inserted to maintain the
// sorted order (equivalent to the lower_bound in C++ or sort.SearchInts in Go).
//
// Time Complexity: O(log N) - the search space is halved on each iteration.
// Space Complexity: O(1) - uses only a few variables for pointers.
func PositionSearch(n int, nums []int, target int) int {
	left := 0
	right := n - 1

	for left <= right {
		// Calculate the middle index. Using 'left + (right-left)/2' instead of
		// '(left + right) / 2' prevents potential integer overflow in languages
		// with fixed-size integers (though less of a concern in Go due to int size).
		middle := left + (right-left)/2

		if nums[middle] == target {
			return middle
		}

		if nums[middle] < target {
			// The target must be in the right half.
			left = middle + 1
		} else {
			// The target must be in the left half.
			right = middle - 1
		}
	}

	// If the loop finishes without finding the target, 'left' will point to
	// the exact index where the target should be inserted to keep the array sorted.
	return left
}
