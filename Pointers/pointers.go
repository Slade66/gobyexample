package main

import "fmt"

func fv(a int) {
	a = 0
}

func fp(a *int) {
	*a = 0
}

func main() {

	i := 1

	fv(i)
	fmt.Println(i)

	fp(&i)
	fmt.Println(i)

	fmt.Println(&i)

}
