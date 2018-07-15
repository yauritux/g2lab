package benchmarking

func FibNew(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	default:
		return FibNew(n-1) + FibNew(n-2)
	}
}
