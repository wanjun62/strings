package strings

import "sort"

// Strings provides convenience methods for string slices
type Strings []string

// Contains returns whether the slice contains a string
func (s Strings) Contains(a string) bool {
	if !sort.StringsAreSorted(s) {
		sort.Strings(s)
	}

	i := sort.SearchStrings(s, a)
	return i < len(s) && s[i] == a
}

// Add a string to the slice
func (s Strings) Add(a string) []string {
	return append(s, a)
}

// Remove a string from a slice
func (s Strings) Remove(a string) []string {
	i := sort.SearchStrings(s, a)
	if s[i] != a {
		return s
	}
	return append(s[:i], s[i+1:]...)
}
