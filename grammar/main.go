package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func (u *User) Aging() {
	u.Age++
}

func main() {
	num, str := 1, "cat"
	fmt.Println(num, str)

	zahyo := struct {
		id   string
		x, y int
	}{
		id: "ABC",
		x:  1,
		y:  2,
	}
	fmt.Println(zahyo)

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	// slice = append(slice[:3], slice[4:]...)
	// slice = slice[:3+copy(slice[3:], slice[4:])]
	// fmt.Println(slice)
	size := 2
	dst := make([][]int, 0, (len(slice)+size-1)/size)
	for size < len(slice) {
		slice, dst = slice[size:], append(dst, slice[0:size:size])
	}
	dst = append(dst, slice)
	fmt.Println(dst)

	m := map[string]int{
		"John":    100,
		"Richard": 200,
	}
	age, ok := m["John"]
	fmt.Println(age, ok)

	delete(m, "John")
	_, ok = m["John"]
	fmt.Println(ok)

	u := &User{
		"Richard",
		15,
	}

	u.Aging()
	fmt.Println(u.Age)

	i := interface{}("Go expert")
	n, ok := i.([]byte)
	fmt.Println(n, ok)

}
