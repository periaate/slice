package slice

import "sort"

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

			lastI = i
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
