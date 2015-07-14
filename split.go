package strings

import "strings"

// SplitAndPreserveQuotes splits a string by spaces, but keeps quoted content together
func SplitAndPreserveQuotes(data string) (set []string) {
	var (
		quoted     bool
		start, pos int
	)
	for pos = 0; pos < len(data); pos++ {
		switch data[pos] {
		case '"':
			quoted = !quoted
		case ' ', '\t':
			if quoted {
				break
			}
			if str := strings.TrimSpace(data[start:pos]); len(str) > 0 {
				set = append(set, strings.Replace(str, "\"", "", -1))
			}
			start = pos + 1
		}
	}
	if start < pos {
		if str := strings.TrimSpace(data[start:pos]); len(str) > 0 {
			set = append(set, strings.Replace(str, "\"", "", -1))
		}
	}
	return
}
