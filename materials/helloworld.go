package main

func main() {
	a := make(map[string]struct{ d, b, c int })
	a["123"] = struct{ d, b, c int }{1, 2, 3}
	a["123"].d = 5
}
