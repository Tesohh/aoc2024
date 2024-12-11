package aoc

func RepeatSlice[T any](v T, times int) []T {
	data := make([]T, times)
	for i := 0; i < times; i++ {
		data[i] = v
	}
	return data
}

// func sliceFilledWithString(size int, str string) []string {
// 	data := make([]string, size)
// 	for i := 0; i < size; i++ {
// 		data[i] = str
// 	}
// 	return data
// }
