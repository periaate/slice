package slice

import "sort"

// SplitWithAll splits given string into an array, using all other `match` strings as
// delimiters. String is matched using the longest delimiter first.
// If no match strings are given, the original string is returned.
// If no matches are found, the original string is returned.
// Matched delimiters are not included in the result.
// If a found match would add a zero-length string to the result, it is ignored.
// Any consecutive matches are treated as one.
// If an empty match string is given (i.e. ""), every character is split.
func SplitWithAll(str string, match ...string) (res []string) {
	if len(match) == 0 || len(str) == 0 {
		return []string{str}
	}

	sort.SliceStable(match, func(i, j int) bool {
		return len(match[i]) > len(match[j])
	})

	var lastI int

	for i := 0; i < len(str); i++ {
		for _, pattern := range match {
			if i+len(pattern) > len(str) {
				continue
			}

			if str[i:i+len(pattern)] != pattern {
				continue
			}

			if len(str[lastI:i]) != 0 {
				res = append(res, str[lastI:i])
			}

			lastI = i + len(pattern)
			if len(pattern) != 0 {
				i += len(pattern) - 1
			}
			break
		}
	}

	if len(str[lastI:]) != 0 {
		res = append(res, str[lastI:])
	}

	if len(res) == 0 {
		return []string{str}
	}

	return res
}
