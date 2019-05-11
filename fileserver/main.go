package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
	http.Handle("/filelist", http.HandlerFunc(FileList))
}

type FileListModel struct{
	Path string
	Files []os.FileInfo
}

func FileList(w http.ResponseWriter, req *http.Request){
	var data FileListModel
	data.Path="/root"

	files, err := ioutil.ReadDir(config.FileLocation)
	if err != nil {
		log.Fatal(err)
	}
	data.Files=files

	tmpl,err:=template.ParseFiles("templates/filelist.html")
	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w,data)
}