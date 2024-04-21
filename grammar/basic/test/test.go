package main

import "fmt"

/*
[問題]
1から20まで昇順で表現されたsliceA、
人の名前が記述されたsliceB、
点の名前と座標の位置が表現されたsliceC
が存在します。
これらはまとめてlistというインターフェースに属します。
すべてのスライスはtestSliceとしてさらにsliceとしてまとめられています。
型スイッチすることにより、処理を分岐させ以下の要件を満たしてください。
(1) 1から20のうち21から30を追加し、また、11から20を削除してください。
(2) 同じ名前は削除して、ユニークなsliceを作成してください。
(3) すべての座標を(+10, +10)してください。
*/

type Point struct {
	Name string
	x, y int
}

func (p *Point) add() {
	p.x += 10
	p.y += 10
}

func main() {
	sliceA := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	sliceB := []string{"John", "Bob", "John", "Alice", "Bob", "John", "Daniel", "Alice"}
	sliceC := []Point{{"first", 1, 2}, {"second", 10, 30}, {"third", -100, -20}}
	list := []interface{}{sliceA, sliceB, sliceC}

	for i, slice := range list {
		switch s := slice.(type) {
		case []int:
			add := []int{21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
			list[i] = append(s[:10], add...)
		case []string:
			m := make(map[string]struct{})
			unique := []string{}
			for _, v := range s {
				if _, ok := m[v]; ok {
					continue
				}
				unique = append(unique, v)
				m[v] = struct{}{}
			}
			list[i] = unique
		case []Point:
			for i := range s {
				s[i].add()
			}
		}
	}

	fmt.Println(list)
}
