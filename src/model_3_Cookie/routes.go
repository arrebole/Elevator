package main

import (
	"fmt"
	"net/http"
)

type myCookie struct {
	Name  string
	Value string
}

func login(w http.ResponseWriter, r *http.Request) {

	// 生成cookie
	cookie := myCookie{
		Name:  "admin",
		Value: "123456",
	}
	// 登陆页面
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
			<form action="/tryLogin" method="post" enctype="application/x-www-form-urlencoded">
				<input type="text" name="userName" />
				<input type="password" name="password" />
				<input type="submit"/>
			</form>
		<body/>
	`

	// 验证是否已经登陆

	// 从请求头中获取 名字是admin的cookie
	cl, err := r.Cookie("admin")
	// 出现错误无法获取或者密码不正确 及没有登陆
	if err != nil {
		fmt.Fprintln(w, html)
	} else {
		fmt.Println(cl.Value)
		if cl.Value != cookie.Value {
			fmt.Fprintln(w, html)
		}
		fmt.Fprintln(w, `<h1>您已经登陆</h1><a href="/admin">进入</a>`)
	}
}

func tryLogin(w http.ResponseWriter, r *http.Request) {

	cookie := http.Cookie{
		Name:  "admin",
		Value: "123456",
	}

	if r.Method != "POST" {
		fmt.Fprintln(w, "not find")
	} else {
		r.ParseForm() // 解析post
		fmt.Println("tryLogin:post", r.PostForm["userName"][0], r.PostForm["password"][0])

		if r.PostForm["userName"][0] == "Bob" && r.PostForm["password"][0] == "123" {
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/admin", http.StatusTemporaryRedirect)
		} else {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		}
	}

}

func admin(w http.ResponseWriter, r *http.Request) {
	// 生成cookie
	cookie := myCookie{
		Name:  "admin",
		Value: "123456",
	}
	//获取cookie
	ck, err := r.Cookie("admin")
	//处理没有 admin的cookie cookie不合法
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
	} else {
		if ck.Value == cookie.Value {
			// 成功验证cookie
			w.Write([]byte("恭喜您 登陆成功"))

		} else {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		}
	}

}
