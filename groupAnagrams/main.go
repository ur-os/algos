package groupAnagrams

import "sort"

func groupAnagrams(strs []string) [][]string {
	matches := map[string][]string{}

	for _, str := range strs {
		anagram := sortString(str)
		if len(matches[anagram]) < 0 {
			matches[anagram] = []string{str}
		} else {
			matches[anagram] = append(matches[anagram], str)
		}
	}

	result := make([][]string, 0)
	for _, strs := range matches {
		result = append(result, strs)
	}

	return result
}

func sortString(w string) string {
	s := []rune(w)
	sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
	return string(s)
}
