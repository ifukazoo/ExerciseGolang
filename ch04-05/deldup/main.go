package deldup

func deldup(strings []string) []string {
	i := 0
	prev := ""
	for _, s := range strings {
		if s != prev {
			strings[i] = s
			i++
			prev = s
		}
	}
	return strings[:i]
}
