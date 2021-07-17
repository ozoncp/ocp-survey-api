package utils

// SplitSlice splits input slice into chunks of specified size.
func SplitSlice(s []int, chunkSize uint) [][]int {
	if s == nil || len(s) == 0 || chunkSize <= 0 {
		return nil
	}

	size := int(chunkSize)
	count := (len(s) + size-1) / size
	os := make([][]int, count)

	start := 0
	i := 0
	for ; i < count-1; i++ {
		os[i] = s[start:start+size]
		start += size
	}
	os[i] = s[start:]	// last chunk
	return os
}

// ReverseMap returns map with keys and values exchanged.
func ReverseMap(m map[string]int) map[int]string {
	if m == nil || len(m) == 0 {
		return nil
	}

	om := make(map[int]string, len(m))
	for k, v := range m {
		om[v] = k
	}
	return om
}

// FilterSlice removes elements specified in hardcoded list.
func FilterSlice(s []int) []int {
	if s == nil || len(s) == 0 {
		return nil
	}

	filter := [...]int{-1, 0, 1}	// hardcoded list
	os := make([]int, 0, len(s))

	loop:
	for _, v := range s {
		for _, f := range filter {
			if v == f {
				continue loop
			}
		}
		os = append(os, v)
	}
	return os
}
