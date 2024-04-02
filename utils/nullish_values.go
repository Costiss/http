package utils

func IfEmptyInt(value int, defaultValue int) int {
	if value == 0 {
		return defaultValue
	}
	return value
}
