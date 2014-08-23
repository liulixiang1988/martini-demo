package main

import (
	"github.com/go-martini/martini"
	"log"
	"net/http"
)

func main() {
	m := martini.Classic()
	//中间件处理器是工作于请求和路由之间的. 本质上来说和Martini其他的处理器没有分别.
	/*
		你可以通过Handlers函数对中间件堆有完全的控制. 它将会替换掉之前的任何设置过的处理器
		m.Handlers(
		  Middleware1,
		  Middleware2,
		  Middleware3,
		)
		中间件处理器可以非常好处理一些功能，像logging(日志), authorization(授权), authentication(认证), sessions(会话), error pages(错误页面),
		以及任何其他的操作需要在http请求发生之前或者之后的:
	*/
	// 验证api密匙
	// m.Use(func(res http.ResponseWriter, req *http.Request) {
	// 	if req.Header.Get("X-API-KEY") != "secret123" {
	// 		res.WriteHeader(http.StatusUnauthorized)
	// 	}
	// })

	//Context.Next()是一个可选的函数用于中间件处理器暂时放弃执行直到其他的处理器都执行完毕. 这样就可以很好的处理在http请求完成后需要做的操作.
	// log 记录请求完成前后  (*译者注: 很巧妙，掌声鼓励.)
	m.Use(func(c martini.Context, log *log.Logger) {
		log.Println("before a request")

		c.Next()

		log.Println("after a request")
	})

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

	//路由模型可能包含参数列表, 可以通过martini.Params服务来获取
	m.Get("/hello/:name", func(params martini.Params) string {
		return "Hello " + params["name"]
	})

	//路由匹配可以通过正则表达式或者glob的形式
	m.Get("/hello2/**", func(params martini.Params) string {
		log.Println("hello2")
		return "Hello " + params["_1"]
	})
	//m.Run()
	//修改端口
	http.ListenAndServe(":8080", m)
}
