package main

import "fmt"

func Filter(p func(int) bool) func(...int) []int {
	return func(args ...int) []int {
		values := []int{}
		for _, v := range args {
			if p(v) {
				values = append(values, v)
			}
		}

		return values
	}
}

func Mapper(f func(int) int) func(...int) []int {
	return func(args ...int) []int {
		values := []int{}
		for _, v := range args {
			values = append(values, f(v))
		}

		return values
	}
}

func Reducer(initial int, f func(int, int) int) func(...int) int {
	current := initial
	return func(args ...int) int {
		for _, v := range args {
			current = f(current, v)
		}

		return current
	}
}

func MapReducer(initial int, mapper func(int) int, reducer func(int, int) int) func(...int) int {
	current := initial
	return func(args ...int) int {
		for _, v := range args {
			current = reducer(current, mapper(v))
		}

		return current
	}
}

func main() {
	odds := Filter(func(x int) bool { return x%2 == 1 })
	evens := Filter(func(x int) bool { return x%2 == 0 })
	fmt.Printf("%v\n", odds(1, 2, 3, 4, 5))   // [1 3 5]
	fmt.Printf("%v\n", evens(1, 2, 3, 4, 5))  // [2 4]
	fmt.Printf("%v\n", odds(6, 7, 8, 9, 10))  // [7 9]
	fmt.Printf("%v\n", evens(6, 7, 8, 9, 10)) // [6 8 10]

	double := Mapper(func(a int) int { return 2 * a })
	fmt.Printf("%v\n", double(1, 2, 3)) // [2 4 6]
	fmt.Printf("%v\n", double(4, 5, 6)) // [8 10 12]

	sum := Reducer(0, func(a, b int) int { return a + b })
	fmt.Printf("%v\n", sum(1, 2, 3))       // 6
	fmt.Printf("%v\n", sum(5))             // 11
	fmt.Printf("%v\n", sum(100, 101, 102)) // 314

	powerSum := MapReducer(
		0,
		func(v int) int { return v * v },
		func(a, v int) int { return a + v },
	)
	fmt.Printf("%v\n", powerSum(1, 2, 3, 4)) // 30
	fmt.Printf("%v\n", powerSum(1, 2, 3, 4)) // 60
	fmt.Printf("%v\n", powerSum())           // 60
}
