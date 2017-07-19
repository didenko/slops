package slops

// Exclude returns a new slice where strings in the regects slice
// removed from the src slice. Both slices expected to be sorted.
func Exclude(src, rejects []string) []string {

	filtered := make([]string, 0)

	for i, j := 0, 0; i < len(src); {

		if j >= len(rejects) || src[i] < rejects[j] {
			filtered = append(filtered, src[i])
			i++
			continue
		}

		if src[i] > rejects[j] {
			j++
			continue
		}

		// last option is src[i] == rejects[j]
		i++
	}

	return filtered
}
