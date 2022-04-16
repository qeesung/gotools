// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package slices

// ConvertIntToAny generics helper.
func ConvertIntToAny(src []int) []any {
	dst := make([]any, len(src))
	for i, value := range src {
		dst[i] = value
	}
	return dst
}

// AppendIntToAny generics helper.
func AppendIntToAny(dst []any, src ...int) []any {
	for _, value := range src {
		dst = append(dst, value)
	}
	return dst
}
