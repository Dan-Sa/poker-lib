package shared

func SliceMap[T, U any](
	slice []T,
	callback func(value T, index int) U,
) []U {
	mappedSlice := make([]U, len(slice))

	for i, v := range slice {
		mappedSlice[i] = callback(v, i)
	}

	return mappedSlice
}
