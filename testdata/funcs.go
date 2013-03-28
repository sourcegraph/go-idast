package vars

func A(b, c string, d int) (w, x float, y bool) {
	return len(b+c) + d
}
