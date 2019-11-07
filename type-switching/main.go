package main

/*
  ## Interface conversions and type assertions
  型のタイプに応じて、その型に合うように変換を行う。
*/

import (
	"fmt"
)

type Stringer interface {
	String() string
}

var value interface{}

func main() {
	switch str := value.(type) {
	case string:
		fmt.Print(str)
	case Stringer:
		fmt.Print(str.String())
	}
}
