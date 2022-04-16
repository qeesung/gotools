// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package slices

// ConvertStringToAny generics helper.
func ConvertStringToAny(src []string) []any {
	dst := make([]any, len(src))
	for i, value := range src {
		dst[i] = value
	}
	return dst
}

// AppendStringToAny generics helper.
func AppendStringToAny(dst []any, src ...string) []any {
	for _, value := range src {
		dst = append(dst, value)
	}
	return dst
}
