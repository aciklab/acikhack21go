package main

import "fmt"

func main() {
	for i := 1; i < 5; i++ {
		fmt.Println("Satir")
	}

	// while-ish loop
	n := 1
	for n < 5 {
		n *= 2
	}
	fmt.Println(n)

	// infinite loop
	sum := 0
	for {
		if sum == 3 {
			break
		}
		sum++ // repeated forever
	}
	fmt.Println(sum)

	// range loop
	strings := []string{"hello", "world"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	// continue
	sum = 0
	for i := 1; i < 5; i++ {
		if i%2 != 0 { // skip odd numbers
			continue
		}
		sum += i
	}
	fmt.Println(sum) // 6 (2+4)

	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", sumof(i, i+1))
	}
	fmt.Println("")

	s := sumoff()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", s(i, i+1))
	}
	fmt.Println("")

	f := fib()
	for i := 0; i < 25; i++ {
		fmt.Printf("%v ", f())
	}
}

func sumof(a int, b int) int {
	return a + b
}

func sumoff() func(a int, b int) int {
	return func(a int, b int) int {
		return a + b
	}
}

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
