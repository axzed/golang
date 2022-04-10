package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 2.解析模板
	files, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 3.渲染模板
	name := "薛文朝"
	err = files.Execute(w, name)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}