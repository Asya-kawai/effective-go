package main

/*
  ## インターフェースとメソッド
  httpパッケージの Handlerインターフェースを例に見ていく。

  httpパッケージのHandlerインターフェースの実装は以下のとおり。
  type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
  }

  この実装があれば、HTTPリクエストをカスタマイズできる。
*/

import (
	"fmt"
	"net/http"
)

type Counter struct {
	n int
}

func (c *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c.n++
	// ignore error
	fmt.Fprintf(w, "counter = %d\n", c.n)
}

func main() {
	c := new(Counter)
	http.Handle("/counter", c)
	// ignore error
	http.ListenAndServe(":80", nil)
}
