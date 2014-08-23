package main

import (
	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()

	//路由匹配的顺序是按照他们被定义的顺序执行的. 最先被定义的路由将会首先被用户请求匹配并调用
	m.Get("/", func() (int, string) {
		return 418, "I'm a teaport" //状态码418
	})

	m.Patch("/", func() {
		// 更新
	})

	m.Post("/", func() {
		// 创建
	})

	m.Put("/", func() {
		// 替换
	})

	m.Delete("/", func() {
		// 删除
	})

	m.Options("/", func() {
		// http 选项
	})

	m.NotFound(func() {
		// 处理 404
	})

	m.Run()
}
