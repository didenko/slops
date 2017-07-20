package slops

// Common gathers same entries from two sorted slices into
// a new slice. The order is preserved. The lesser number of
// duplicates is preserved
func Common(left, right []string) []string {

	common := make([]string, 0)

	for i, j := 0, 0; i < len(left) && j < len(right); {
		if left[i] < right[j] {
			i++
			continue
		}

		if left[i] > right[j] {
			j++
			continue
		}

		common = append(common, left[i])
		i++
		j++
	}

	return common
}
