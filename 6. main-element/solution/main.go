package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type IO struct {
	s *bufio.Scanner
	w *bufio.Writer
}

func NewIO() *IO {
	s := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	s.Split(bufio.ScanWords)

	return &IO{
		s: s,
		w: w,
	}
}

func (o *IO) Close() error {
	return o.w.Flush()
}

func (o *IO) ReadConditions() ([]int, int) {
	o.s.Scan()
	n, _ := strconv.Atoi(o.s.Text())

	elements := make([]int, n)

	for i := 0; i < n; i++ {
		o.s.Scan()
		elements[i], _ = strconv.Atoi(o.s.Text())
	}

	return elements, n
}

func (o *IO) WriteResult(v int) {
	fmt.Fprintln(o.w, v)
}

func main() {
	o := NewIO()
	elements, n := o.ReadConditions()

	// Cases:
	// res := MainElement([]int{1, 2, 3, 4, 5}, 5); fmt.Println(res)
	// res := MainElement([]int{1, 1, 1, 4, 5}, 5); fmt.Println(res)
	// res := MainElement([]int{}, 0); fmt.Println(res)

	res := MainElement(elements, n)
	o.WriteResult(res)
}

func MainElement(elements []int, n int) int {
	if n == 0 {
		return -1
	}

	seen := make(map[int]int)

	for _, element := range elements {
		seen[element]++
	}

	for element, count := range seen {
		if count > n/2 {
			return element
		}
	}

	return -1
}
