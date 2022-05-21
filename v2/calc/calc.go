package calc

func Add(a ...int) int {
	s := 0
	for _, val := range a {
		s += val
	}
	return s
}

func Sub(a, b int) int {
	return a - b
}
