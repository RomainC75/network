package cases

var Addresses = []struct {
	Address  string
	Expected []int
	Err      bool
}{
	{"212.0.2.43", []int{212, 0, 2, 43}, false},
	{"212.0.2.434", []int{212, 0, 2, 43}, true},
	{"212.0.2.", []int{212, 0, 2, 43}, true},
	{"212.0.2.43.2", []int{212, 0, 2, 43}, true},
}

var Compare = []struct {
	Min      []int
	Max      []int
	Expected bool
}{
	{[]int{12, 15, 42, 123}, []int{12, 14, 16, 123}, false},
	{[]int{12, 15, 42, 123}, []int{12, 16, 16, 123}, true},
	{[]int{12, 15, 42, 123}, []int{12, 15, 42, 123}, true},
}
