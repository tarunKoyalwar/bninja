package internal

import "strings"

// Split : Similar to Strings.Feilds with Custom separator
func Split(s string, delim rune) []string {

	// Must trim the string first
	s = strings.TrimSpace(s)

	arr := []string{}

	var sb strings.Builder

	for _, v := range s {
		if v != delim {
			sb.WriteRune(v)
		} else {
			if sb.Len() != 0 {
				arr = append(arr, sb.String())
				sb.Reset()
			}
		}
	}

	if sb.Len() != 0 {
		arr = append(arr, sb.String())
		sb.Reset()
	}

	return arr
}
