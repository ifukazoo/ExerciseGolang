package rotate

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return
}

func rotateSimple(s []int, n int) {
	reverse(s[:n])
	reverse(s[n:])
	reverse(s)
}

func rotateFor(s []int, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < len(s)-1; j++ {
			s[j], s[j+1] = s[j+1], s[j]
		}
	}
}
