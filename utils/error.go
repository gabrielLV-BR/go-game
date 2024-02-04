package utils

func Unwrap[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
