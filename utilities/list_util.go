package utils

func RemoveValue(slice []string, value string) []string {
	var index int
	for i, v := range slice {
		if v == value {
			index = i
			break
		}
	}

	// If the value was found, remove it
	if index < len(slice) {
		slice = append(slice[:index], slice[index+1:]...)
	}

	return slice
}

func Contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
