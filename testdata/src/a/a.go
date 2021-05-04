package a

func f() {

	test := []int{1, 2, 3, 4, 5}

	a := make([]int, 5)
	for _, tt := range test {
		a = append(a, tt) // want "sleuth detects illegal"
	}

	for _, tt := range test {
		a = append(a, tt) // ok
	}

	b := make([]int, 5)
	c := make([]int, 5)

	for _, tt := range test {
		b = append(b, tt) // want "sleuth detects illegal"
	}

	for _, tt := range test {
		c = append(c, tt) // want "sleuth detects illegal"
	}

	var d []int

	for _, tt := range test {
		d = append(d, tt) // ok
	}
}
