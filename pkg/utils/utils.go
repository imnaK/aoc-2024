package utils

func GetDiff(num1 int, num2 int) int {
	diff := num1 - num2
	if diff < 0 {
		diff = -diff
	}
	return diff
}

func ReverseArray[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func RemoveEmpty[T comparable](slice []T) []T {
	var result []T
	var empty T
	for _, item := range slice {
		if item != empty {
			result = append(result, item)
		}
	}
	return result
}

func ArrayRemove[T any](slice []T, idx int) []T {
	return append(slice[:idx], slice[idx+1:]...)
}

func ArrayContains[T comparable](slice []T, value T) bool {
	for _, elem := range slice {
		if elem == value {
			return true
		}
	}

	return false
}
