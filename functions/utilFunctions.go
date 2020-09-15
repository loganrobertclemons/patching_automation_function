package functions

/*
 * Return a string from a fact, with default "UNSET"
 */
func GetFactString(fact interface{}) string {
	if fact != nil {
		return fact.(string)
	}
	return "UNSET"
}

/*
 * Return an array of strings from a fact (using default values getFactString, which shouldn't occur)
 */
func GetFactArrayOfStrings(fact []interface{}) []string {
	result := make([]string, len(fact))
	for i, v := range fact {
		result[i] = GetFactString(v)
	}
	return result
}

/*
 * Return an integer from a fact, default 0
 */
func GetFactInt(fact interface{}) int {
	if fact != nil {
		return int(fact.(float64))
	}
	return 0
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
