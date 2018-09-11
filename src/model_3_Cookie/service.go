/*----------------------------------------------------

	使用cookie模拟登陆

------------------------------------------------------*/

package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/", login)
	http.HandleFunc("/login", login)
	http.HandleFunc("/admin", admin)
	http.HandleFunc("/tryLogin", tryLogin)

	service := &http.Server{
		Addr: "127.0.0.1:8080",
	}
	service.ListenAndServe()

}
