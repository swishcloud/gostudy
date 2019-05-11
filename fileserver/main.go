package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

func FileList(w http.ResponseWriter, req *http.Request){
	fmt.Fprintln(w,config.FileLocation)
	files, err := ioutil.ReadDir(config.FileLocation)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Fprintln(w,f.Name())
	}
}