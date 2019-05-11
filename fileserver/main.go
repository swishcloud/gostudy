package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	Initialize()
	addr:= flag.String("addr", config.Host, "http service address")
	BindHandlers()
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func Initialize()  {
	config=ReadConfig()
}

func BindHandlers(){
	http.Handle("/filelist", http.HandlerFunc(FileList))
}

func FileList(w http.ResponseWriter, req *http.Request){
	fmt.Fprintln(w,config.FileLocation)
}