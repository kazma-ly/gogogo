package main

import "fmt"

func main() {
	m := map[string]string{
		"c1": "c1v",
		"c2": "c2v",
	}

	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, v)
	}
}
