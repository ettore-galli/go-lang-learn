package functools

func fmap[T any, U any](list []T, fu func(T) U) []U {
	transformed := make([]U, len(list))
	for i, v := range list {
		transformed[i] = fu(v)
	}
	return transformed
}
