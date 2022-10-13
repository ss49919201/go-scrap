package main

// 素数判定
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	// nの平方根までで割り切れるかを調べる
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
