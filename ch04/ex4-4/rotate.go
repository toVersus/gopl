package rotate

func rotate(a []int) {
	first := a[0]
	copy(a, a[1:])
	a[len(a)-1] = first
}
