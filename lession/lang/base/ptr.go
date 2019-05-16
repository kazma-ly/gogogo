package main

import "fmt"

type (
	Person struct {
		name string
		sex  string
	}
)

// 只有值传递， 传递 指针的值
func main() {
	v := 10
	passByRef(&v)
	fmt.Println(v)

	p1 := &Person{name: "小泽", sex: "1"}
	p1.change("yq")
	fmt.Println(p1)

	p2 := &Person{}
	p2.init("haha", "cool")
	fmt.Println(p2)
}

func passByRef(v *int) {
	*v += 1
}

func (p *Person) change(name string) {
	p.name = name
}

func (p *Person) init(name, sex string) {
	p.name = name
	p.sex = sex
}
