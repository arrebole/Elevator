package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		<form action="/post/" method="post" enctype="multipart/form-data">
			<input type="text" name="first_name" />
			<input type="text" name="last_name" />
			<input type="file" name="uploaded" />
			<input type="submit"/>
		</form>
	<body/>
	`
	fmt.Fprintln(w, html)
}

// 解析post
func post(w http.ResponseWriter, r *http.Request) {
	// 解析 上传文件
	r.ParseMultipartForm(1024)
	fileHeader := r.MultipartForm.File["uploaded"][0]
	file, err := fileHeader.Open()
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, r.Form["first_name"], r.Form["last_name"], "\n", string(data))
		}

	}
}

// 重定向
func redirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://bilibili.com")
	w.WriteHeader(302) //必须设置状态码
}

type user struct {
	Name string
	Age  int
}

// 输出json
func sendJSON(w http.ResponseWriter, r *http.Request) {
	// 设置输出格式
	w.Header().Set("Content-Type", "application/json")

	data := &user{
		Name: "hello",
		Age:  0,
	}
	jso, _ := json.Marshal(data)
	w.Write(jso)
}
