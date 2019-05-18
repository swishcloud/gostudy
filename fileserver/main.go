package main

import (
	"flag"
	"fmt"
	"github.com/github-123456/gostudy/aesencryption"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

const SessionName  ="session"

func main() {
	addr:= flag.String("addr", config.Host, "http service address")
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func init()  {
	config=ReadConfig()
	BindHandlers()
}



func BindHandlers(){
	mux:=http.NewServeMux()
	http.Handle("/",mux)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.Handle("/filelist/", http.HandlerFunc(FileList))
	mux.Handle("/download", http.HandlerFunc(Download))
	mux.Handle("/login", http.HandlerFunc(login))
	mux.Handle("/index", http.HandlerFunc(index))
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		fmt.Fprintf(w, "Welcome to the home page!")
	})
}

type FileListModel struct{
	Path string
	Files []os.FileInfo
}

func FileList(w http.ResponseWriter, req *http.Request){
	cookie:=http.Cookie{Name:SessionName,Value:aesencryption.Encrypt("hello world!"),Path:"/"}
	http.SetCookie(w,&cookie)

	var data FileListModel
	re:=regexp.MustCompile(`/filelist/`)
	data.Path=re.ReplaceAllString(req.URL.Path,"")
	path:=filepath.Join(config.FileLocation,data.Path)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		http.NotFound(w,req)
		return
	}
	data.Files=files

	tmpl,err:=template.ParseFiles("templates/filelist.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w,data)
}

func Download(w http.ResponseWriter, req *http.Request){
	cookie,err:=req.Cookie(SessionName)
	if err!=nil{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	plain:=aesencryption.Decrypt(cookie.Value)
	fmt.Fprintln(w,plain)
}