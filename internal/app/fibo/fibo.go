package fibo

// Fibonacci numbers slice
func Fibo(m, n uint32, f *[]uint32) {
	arr := *f
	for i := m; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
}
