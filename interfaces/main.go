package main

/*
  ## Interfaces
  Goのインターフェースはオブジェクトの動作を定義する。

  ここでは、Len(),Less(i,j int),Swap(i,j int)を含むsort.interfaceを実装している。
  つまり、sortパッケージを用いてソートができる。
*/

import (
	"fmt"
	"sort"
)

// Sequence []int
type Sequence []int

// Len returns length of sequence
func (s Sequence) Len() int {
	return len(s)
}

func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Sequence) Copy() Sequence {
	copy := make(Sequence, 0, len(s))
	return append(copy, s...)
}

// String は リストの中身を文字列として出力する関数。
// 文字列として出力する前にソートしておく必要がある。
func (s Sequence) String() string {
	s = s.Copy()
	sort.Sort(s)
	str := "["
	for i, elem := range s {
		if i > 0 {
			str += " "
		}
		str += fmt.Sprint(elem)
	}
	return str + "]"
}

func main() {
	s := Sequence{1, 2, 3}
	fmt.Printf("len : %v\n", s.Len())

	fmt.Println("show Sequence as string.")

	fmt.Printf("string: %s\n", s.String())
}
