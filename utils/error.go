package utils

func Unwrap[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func Assert(e error) {
	if e != nil {
		panic(e)
	}
}
