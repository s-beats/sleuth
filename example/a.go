package main

func main() {
	test := []int{1, 2, 3, 4, 5}
	a := make([]int, 5)
	for _, tt := range test {
		a = append(a, tt) // want "sleuth detects illegal"
	}
}
