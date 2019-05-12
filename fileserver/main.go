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
	mux.Handle("/filelist/", http.HandlerFunc(FileList))
	mux.Handle("/download", http.HandlerFunc(Download))
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
	cipher:=aesencryption.Encrypt("hello world!")
	plain:=aesencryption.Decrypt(cipher)
	fmt.Fprintln(w,cipher)
	fmt.Fprintln(w,plain)
}