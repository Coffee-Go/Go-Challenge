package fib

var fib map[int]int

func init() {
	fib = make(map[int]int)
}

func Fib(n int) int {
	if n <= 2 {
		return 1
	}
	if f, ok := fib[n]; ok {
		return f
	}

	f := Fib(n-1) + Fib(n-2)
	fib[n] = f
	return f
}
