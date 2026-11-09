// Package main implements a program that reads an array of integers
// and finds the last even number in the array. If no even number is found,
// it returns -1.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// IO provides a fast, buffered input/output mechanism for reading from standard input
// and writing to standard output. It wraps bufio.Scanner and bufio.Writer.
type IO struct {
	s *bufio.Scanner
	w *bufio.Writer
}

// NewIO initializes and returns a new IO instance. The scanner is configured
// to split the input by words (spaces, newlines) to simplify parsing integers.
func NewIO() *IO {
	s := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	s.Split(bufio.ScanWords)

	return &IO{s: s, w: w}
}

// Flush writes any buffered data to the underlying io.Writer.
// It must be called before the program exits to ensure all output is printed.
func (io *IO) Flush() {
	io.w.Flush()
}

// ReadInput reads the number of elements (n) and the array of integers from standard input.
// It returns the size of the array and the slice containing the elements.
//
// Note: There is a logical error here. The slice 'nums' is created with size 0
// (because 'n' is initialized to 0 right before this line). This will cause a
// panic (index out of range) in the loop.
func (io *IO) ReadInput() (int, []int) {
	n := 0
	nums := make([]int, n) // BUG: creates a slice of length 0

	io.s.Scan()
	n, _ = strconv.Atoi(io.s.Text())

	for i := 0; i < n; i++ {
		io.s.Scan()
		nums[i], _ = strconv.Atoi(io.s.Text()) // Panics here because len(nums) == 0
	}

	return n, nums
}

// WriteOutput writes the integer result to the buffered writer.
func (io *IO) WriteOutput(res int) {
	fmt.Fprintf(io.w, "%v", res)
}

func main() {
	io := NewIO()
	defer io.Flush()

	n, nums := io.ReadInput()

	res := QualityControl(n, nums)

	io.WriteOutput(res)
}

// QualityControl iterates through the given slice of integers and finds
// the last even number. If no even number is present in the slice, it returns -1.
//
// Time Complexity: O(N) - where N is the number of elements in the slice.
// Space Complexity: O(1) - uses only a constant amount of extra space (the 'prev' variable).
func QualityControl(n int, nums []int) int {
	prev := -1

	for _, v := range nums {
		if v%2 == 0 {
			prev = v
		}
	}

	return prev
}
