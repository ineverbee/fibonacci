package fibo

// Fibonacci numbers slice
func Fibo(n uint32) []uint32 {
	f := make([]uint32, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; i <= int(n); i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[:n+1]
}
