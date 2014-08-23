package main

import (
	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()
	m.Get("/", func() (int, string) {
		return 418, "I'm a teaport" //状态码418
	})
	m.Run()
}
