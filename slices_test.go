package genericsbench

import (
	"testing"
)

func BenchmarkReduceGeneric_10elements_IntegerSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		_ = reduce(numbers, func(acc, current int) int {
			return acc + current
		}, 0)
	}
}

func BenchmarkReduceLoop_10elements_IntegerSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		sum := 0
		for _, n := range numbers {
			sum += n
		}
	}
}

func BenchmarkContainsGeneric_10elements_ContainsLastElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		toFind := 10
		_ = contains(numbers, toFind)
	}
}

func BenchmarkContainsLoop_10elements_ContainsLastElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		toFind := 10
		_ = func() bool {
			for _, n := range numbers {
				if n == toFind {
					return true
				}
			}
			return false
		}()
	}
}

func BenchmarkFilterGeneric_10elements_ReturnEvenIntegers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		_ = filter(numbers, func(n int) bool { return n%2 == 0 })
	}
}

func BenchmarkFilterLoop_10elements_ReturnEvenIntegers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		var even []int
		for _, n := range numbers {
			if n%2 == 0 {
				even = append(even, n)
			}
		}
	}
}
