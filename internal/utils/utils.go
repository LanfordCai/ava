package utils

// Contains - a string in a []string or not
func Contains(sli []string, target string) bool {
	for _, s := range sli {
		if s == target {
			return true
		}
	}

	return false
}
