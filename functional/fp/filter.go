package fp

func Filter[TValue any](values []TValue, f func(TValue) bool) []TValue {
	result := make([]TValue, 0)
	for _, value := range values {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}