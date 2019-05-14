package main

import (
	"html/template"
	"log"
	"net/http"
)


func login(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if req.Method==http.MethodGet{
		tmpl, err := template.ParseFiles("templates/login.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, nil)
	}else{
		account:=req.PostForm.Get("account")
		password:=req.PostForm.Get("password")
		if account=="123" && password=="456"{
			HandlerResult{}.Write(w)
		}else {
			HandlerResult{Error:"账号或密码有误",Data:"test"}.Write(w)
		}
	}
}