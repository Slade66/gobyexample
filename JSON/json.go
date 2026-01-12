package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Person1 struct {
	Name string
	Age  int
}

type Person2 struct {
	Name string `json:"姓名"`
	Age  int    `json:"年龄"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(1.1)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("lyz")
	fmt.Println(string(strB))

	slcD := []string{"a", "b", "c"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"a": 1, "b": 2}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	lyz := Person1{"lyz", 24}
	lyzB, _ := json.Marshal(lyz)
	fmt.Println(string(lyzB))

	lx := Person2{"lx", 21}
	lxB, _ := json.Marshal(lx)
	fmt.Println(string(lxB))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]any
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]any)
	fmt.Println(strs)
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"name": "lyz", "age": 24}`
	var p1 Person1
	json.Unmarshal([]byte(str), &p1)
	fmt.Println(p1)
	fmt.Println(p1.Age)

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

	dec := json.NewDecoder(strings.NewReader(str))
	var p2 Person1
	dec.Decode(&p2)
	fmt.Println(p2)
}
