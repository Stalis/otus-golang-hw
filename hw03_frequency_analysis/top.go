package hw03frequencyanalysis

import (
	"regexp"
	"sort"
)

func TakeStrings(slice []string, count int) []string {
	l := len(slice)
	if l > count {
		l = count
	}

	return slice[:l]
}

func Top10(input string) []string {
	r := regexp.MustCompile(`\S+`)
	matches := r.FindAllString(input, -1)

	entries := make(map[string]int, len(matches))
	for _, v := range matches {
		entries[v] += 1
	}

	keys := make(sort.StringSlice, 0, len(entries))
	for k := range entries {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i int, j int) bool {
		iEntry := entries[keys[i]]
		jEntry := entries[keys[j]]
		if iEntry == jEntry {
			return sort.StringsAreSorted([]string{keys[i], keys[j]})
		}
		return iEntry > jEntry
	})

	return TakeStrings(keys, 10)
}
