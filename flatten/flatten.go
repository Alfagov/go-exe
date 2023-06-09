package flatten

// Generic flatten function
func Flatten(input interface{}) []interface{} {

	result := []interface{}{}

	switch t := input.(type) {
	case []interface{}:
		for _, v := range t {
			result = append(result, Flatten(v)...)
		}

	case interface{}:
		result = append(result, input)
	}

	return result
}
