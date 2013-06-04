package vars

func A(b, c string, d int) (w, x int, y bool) {
	return len(b+c) + d, 0, true
}

func B(_, _ string, _ int) (_, _ string, _ int) {
	return
}
