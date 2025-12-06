package main

func main() {
	i := 1
	for i <= 3 {
		println(i)
		i = i + 1
	}

	for i := 0; i < 3; i++ {
		println(i)
	}

	for i := range 3 {
		println(i)
	}

	for {
		println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		println(n)
	}
}
