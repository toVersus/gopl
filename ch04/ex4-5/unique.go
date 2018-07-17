// ex4-5 eliminates a duplicated string.
package unique

func unique(strings []string) []string {
	seen := make(map[string]struct{}, len(strings))
	j := 0
	for _, str := range strings {
		if _, ok := seen[str]; ok {
			continue
		}
		seen[str] = struct{}{}
		strings[j] = str
		j++
	}
	return strings[:j]
}
