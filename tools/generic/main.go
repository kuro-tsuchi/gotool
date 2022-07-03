package main

import "fmt"

type MySlice[T int | float32] []T

func (s MySlice[T]) Sum() T {
	var sum T
	for _, value := range s {
		sum += value
	}
	return sum
}
func main() {
	var s MySlice[int] = []int{1, 2, 3}
	sum := s.Sum()
	fmt.Println(sum)

	var s2 MySlice[float32] = []float32{1.1, 2.2, 3.3}
	t := s2.Sum()
	fmt.Println(t)

}
