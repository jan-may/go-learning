package fp

func Map[TValue, TResult any](values []TValue, f func(TValue) TResult) []TResult {
	result := make([]TResult, len(values))
	for i, value := range values {
		result[i] = f(value)
	}
	return result
}