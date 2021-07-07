// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package slices

// ConvertStringToInterface generics helper.
func ConvertStringToInterface(src []string) []interface{} {
	dst := make([]interface{}, len(src))
	for i, value := range src {
		dst[i] = value
	}
	return dst
}

// AppendStringToInterface generics helper.
func AppendStringToInterface(dst []interface{}, src ...string) []interface{} {
	for _, value := range src {
		dst = append(dst, value)
	}
	return dst
}

// ConvertIntToInterface generics helper.
func ConvertIntToInterface(src []int) []interface{} {
	dst := make([]interface{}, len(src))
	for i, value := range src {
		dst[i] = value
	}
	return dst
}

// AppendIntToInterface generics helper.
func AppendIntToInterface(dst []interface{}, src ...int) []interface{} {
	for _, value := range src {
		dst = append(dst, value)
	}
	return dst
}

// ConvertInt64ToInterface generics helper.
func ConvertInt64ToInterface(src []int64) []interface{} {
	dst := make([]interface{}, len(src))
	for i, value := range src {
		dst[i] = value
	}
	return dst
}

// AppendInt64ToInterface generics helper.
func AppendInt64ToInterface(dst []interface{}, src ...int64) []interface{} {
	for _, value := range src {
		dst = append(dst, value)
	}
	return dst
}

// ConvertUint64ToInterface generics helper.
func ConvertUint64ToInterface(src []uint64) []interface{} {
	dst := make([]interface{}, len(src))
	for i, value := range src {
		dst[i] = value
	}
	return dst
}

// AppendUint64ToInterface generics helper.
func AppendUint64ToInterface(dst []interface{}, src ...uint64) []interface{} {
	for _, value := range src {
		dst = append(dst, value)
	}
	return dst
}
