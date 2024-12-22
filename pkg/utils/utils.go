package utils

func GetDiff(num1 int, num2 int) int {
	diff := num1 - num2
	if diff < 0 {
		diff = -diff
	}
	return diff
}
