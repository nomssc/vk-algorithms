// Package main implements a solution for finding the majority element in an array.
// A majority element is defined as an element that appears strictly more than n/2 times.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// IO is a wrapper around bufio.Scanner and bufio.Writer to provide fast I/O.
// This is crucial in competitive programming to prevent Time Limit Exceeded (TLE)
// errors caused by slow standard input/output operations.
type IO struct {
	s *bufio.Scanner
	w *bufio.Writer
}

// NewIO initializes and returns a new IO instance.
func NewIO() *IO {
	s := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	// Configure the scanner to split input by words (spaces, newlines, etc.).
	// This makes it easy to read integers regardless of how they are spaced on the line.
	s.Split(bufio.ScanWords)

	return &IO{
		s: s,
		w: w,
	}
}

// Close flushes any buffered output to the standard output.
// It must be called before the program exits to ensure all results are printed.
func (o *IO) Close() error {
	return o.w.Flush()
}

// ReadConditions reads the array size (n) and the array elements from standard input.
// Returns the slice of elements and the integer n.
func (o *IO) ReadConditions() ([]int, int) {
	o.s.Scan()
	n, _ := strconv.Atoi(o.s.Text()) // Assuming valid input as per contest constraints

	elements := make([]int, n)

	for i := 0; i < n; i++ {
		o.s.Scan()
		elements[i], _ = strconv.Atoi(o.s.Text())
	}

	return elements, n
}

// WriteResult writes the final integer result to the buffered writer.
func (o *IO) WriteResult(v int) {
	fmt.Fprintln(o.w, v)
}

func main() {
	// Initialize fast I/O
	o := NewIO()
	// Ensure the buffer is flushed when main returns
	defer o.Close()

	elements, n := o.ReadConditions()

	// Cases for manual testing (uncomment to run):
	// res := MainElement([]int{1, 2, 3, 4, 5}, 5); fmt.Println(res)
	// res := MainElement([]int{1, 1, 1, 4, 5}, 5); fmt.Println(res)
	// res := MainElement([]int{}, 0); fmt.Println(res)

	// Find the majority element and write it to output
	res := MainElement(elements, n)
	o.WriteResult(res)
}

// MainElement finds the element that appears more than n/2 times in the array.
// It uses a hash map to count frequencies.
//
// Time Complexity: O(N) - where N is the number of elements.
// We iterate through the array once to build the map, and potentially iterate
// through the map (which has at most N unique elements).
//
// Space Complexity: O(N) - In the worst case (all elements are unique),
// the map will store N key-value pairs.
func MainElement(elements []int, n int) int {
	// Handle edge case: empty array cannot have a majority element.
	if n == 0 {
		return -1
	}

	// Map to store the frequency of each element.
	// Key: element value, Value: count of occurrences.
	seen := make(map[int]int)

	// Count the occurrences of each element.
	for _, element := range elements {
		seen[element]++
	}

	// Iterate through the frequencies to find the majority element.
	// A majority element appears strictly more than n/2 times.
	for element, count := range seen {
		if count > n/2 {
			return element
		}
	}

	// Return -1 if no majority element is found.
	return -1
}
