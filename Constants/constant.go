package main

import "math"

const s string = "constant"

func main() {
	println(s)

	const n = 5_0000_0000

	const d = 3e20 / n
	println(d)

	println(int64(d))

	println(math.Sin(n))
}