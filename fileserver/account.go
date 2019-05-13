package main

import (
	"html/template"
	"log"
	"net/http"
)

type LoginModel struct{
	Account string
}

func login(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if req.Method==http.MethodGet{
		login_get(w,req,LoginModel{})
	}else{
		account:=req.PostForm.Get("account")
		//password:=req.PostForm.Get("password")
		login_get(w,req,LoginModel{Account:account})
	}
}
func login_get(w http.ResponseWriter, req *http.Request,data LoginModel) {
		tmpl, err := template.ParseFiles("templates/login.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, data)
}
