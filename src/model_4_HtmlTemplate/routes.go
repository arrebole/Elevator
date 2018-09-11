package main

import (
	"net/http"
	"text/template"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 读取并解析 渲染模板
		t, _ := template.ParseFiles("./public/helloWorld.html")
		// 向模板中传入数据
		t.ExecuteTemplate(w, "helloWorld.html", "hello world")
	} else {
		w.WriteHeader(404)
	}

}

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(404)
		return
	}
	t, _ := template.ParseFiles("./public/layout.html", "./public/header.html")
	t.ExecuteTemplate(w, "layout.html", map[string]interface{}{
		"name": "Bob",
		"age":  16,
		"time": time.Now().String(),
	})
}
