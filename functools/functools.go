package functools

func Fmap[T any, R any](iterable []T, processor func(T) R) []R {
	var result []R
	for _, t := range iterable {
		result = append(result, processor(t))
	}
	return result
}

func Ffilter[T any](iterable []T, predicate func(T) bool) []T {
	var result []T
	for _, t := range iterable {
		if predicate(t) {
			result = append(result, t)
		}
	}
	return result
}

func Freduce[T any, K any](iterable []T, reducer func(K, T, int) K, initialValue K) K {
	var result K = initialValue
	for idx, t := range iterable {
		result = reducer(result, t, idx)
	}
	return result
}
