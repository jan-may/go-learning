package fp

func Reduce[TValue, TResult any](values []TValue, f func(TResult, TValue) TResult, initialValue TResult) TResult {
	result := initialValue
	for _, value := range values {
		result = f(result, value)
	}
	return result
}