/*---------------------------------------------

			http静态服务器模板
			update 2018/9/4

-----------------------------------------------*/
package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	html := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>Document</title>
	</head>
	<body>
		<div id="app"></div>
		<script type="text/javascript" src="/static/app.js" defer></script>
	</body>
	</html>`
	fmt.Fprintln(w, html)
}

func main() {

	// mux是多路复用处理器 类似总闸
	mux := http.NewServeMux()

	// file是 文件处理器 静态文件处理
	file := http.FileServer(http.Dir("./public"))

	// mux 绑定处理器函数或者处理器
	// handle 是处理器 结构体类型
	// handleFunc 是处理器函数 函数类型
	mux.Handle("/static/", file)
	mux.HandleFunc("/", index)

	/*----------------------------------------------------------*/

	// 配置文件
	service := &http.Server{
		Addr: "127.0.0.1:8080",
		// 交给mux处理器处理
		Handler: mux,
	}

	// 启动监听
	fmt.Println("service runing http://127.0.0.1:8080")
	service.ListenAndServe()
}
