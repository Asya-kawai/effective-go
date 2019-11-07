package main

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

// 上は sort.Interface に実装されているメソッドなので、
// sortを利用したい場合には必要

func (s Sequence) Copy() Sequence {
	copy := make(Sequence, 0, len(s))
	return append(copy, s...)
}

// String は リストの中身を文字列として出力する関数。
// 文字列として出力する前にソートしておく必要がある。
func (s Sequence) String() string {
	s = s.Copy()
	sort.Sort(s)
	// ../interfaces/main.go のロジックを Sprintにて書き換え
	return fmt.Sprint([]int(s))
}

func main() {
	s := Sequence{1, 2, 3}
	fmt.Println("show Sequence as string.")
	fmt.Printf("string: %s\n", s.String())
}
