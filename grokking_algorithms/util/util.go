package util

// While equality works for interface{}, greater than or lower than doesn't.

// GreaterThan whether v1 is higher than v2.
func GreaterThan(v1 interface{}, v2 interface{}) bool {
	switch v1.(type) {
	case int:
		return v1.(int) > v2.(int)
	case string:
		return v1.(string) > v2.(string)
	default:
		return v1.(int) > v2.(int)
	}
}

// GreaterThanOrEqual whether v1 is greater than or equal to v2.
func GreaterThanOrEqual(v1 interface{}, v2 interface{}) bool {
	return GreaterThan(v1, v2) || v1 == v2
}

// LowerThan whether v1 is lower than v2.
func LowerThan(v1 interface{}, v2 interface{}) bool {
	switch v1.(type) {
	case int:
		return v1.(int) < v2.(int)
	case string:
		return v1.(string) < v2.(string)
	default:
		return v1.(int) < v2.(int)
	}
}

// LowerThanOrEqual whether v1 is lower than or equal to v2.
func LowerThanOrEqual(v1 interface{}, v2 interface{}) bool {
	return LowerThan(v1, v2) || v1 == v2
}
