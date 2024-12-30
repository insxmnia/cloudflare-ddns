package utility

func SliceIncludes(slice []interface{}, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
