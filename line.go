package strings

import "regexp"

// Line provides a way to move through a string slice
type Line []string

// Trim removes a left-side prefix from the first element, returning whether it happened
func (c Line) Trim(prefix string) bool {
	if len(c) == 0 || len(c[0]) < len(prefix) {
		return false
	}

	c[0] = c[0][len(prefix):]
	return true
}

// Len returns the number of elements
func (c Line) Len() int {
	return len(c)
}

// Match returns whether the current element matches the string
func (c Line) Match(s string) bool {
	return len(c) > 0 && c[0] == s
}

// MatchRe returns whether the current element matches the regexp
func (c Line) MatchRe(re *regexp.Regexp) bool {
	return len(c) > 0 && re.MatchString(c[0])
}

// Next advances the iterator and returns the Line and if it worked
func (c Line) Next() (Line, bool) {
	if len(c) <= 1 {
		return c, false
	}

	return Line(c[:0+copy(c[0:], c[1:])]), true
}
