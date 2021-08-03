package utils

// SplitSlice splits input slice into chunks of specified size.
func SplitSlice(slice []int, chunkSize int) [][]int {
	if len(slice) == 0 || chunkSize <= 0 {
		return nil
	}

	count := (len(slice) + chunkSize - 1) / chunkSize
	res := make([][]int, count)

	start := 0
	i := 0
	for ; i < count-1; i++ {
		res[i] = slice[start : start+chunkSize]
		start += chunkSize
	}
	res[i] = slice[start:] // last chunk
	return res
}

// ReverseMap returns map with keys and values exchanged.
func ReverseMap(m map[string]int) map[int]string {
	if len(m) == 0 {
		return nil
	}

	res := make(map[int]string, len(m))
	for k, v := range m {
		res[v] = k
	}
	return res
}

// FilterSlice removes elements specified in filter list.
func FilterSlice(slice []int, filter []int) []int {
	if len(slice) == 0 {
		return nil
	}

	flt := make(map[int]bool, len(filter))
	for _, v := range filter {
		flt[v] = true
	}

	res := make([]int, 0, len(slice))
	for _, v := range slice {
		if !flt[v] {
			res = append(res, v)
		}
	}
	return res
}
