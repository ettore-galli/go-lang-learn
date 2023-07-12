package executor

func Contains[T comparable](slice []T, element T) bool {
	for _, g := range slice {
		if g == element {
			return true
		}
	}
	return false
}
