package genericsbench

func reduce[T, M any](s []T, f func(M, T) M, init M) M {
	accumulator := init
	for _, v := range s {
		accumulator = f(accumulator, v)
	}

	return accumulator
}

func contains[T comparable](s []T, toFind T) bool {
	for _, v := range s {
		if v == toFind {
			return true
		}
	}

	return false
}

func filter[T any](s []T, f func(T) bool) []T {
	var filtered []T
	for _, t := range s {
		if f(t) {
			filtered = append(filtered, t)
		}
	}

	return filtered
}
