package fileserver

import (
	"flag"
	"fmt"
	"gostudy/aesencryption"
	"log"
	"net/http"
	"os"
)

const SessionName = "session"

func main() {
	addr := flag.String("addr", config.Host, "http service address")
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func init() {
	config = ReadConfig()
	BindHandlers()
}

func BindHandlers() {
	mux := http.NewServeMux()
	http.Handle("/", mux)
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

type FileListModel struct {
	Path  string
	Files []os.FileInfo
}

func Download(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie(SessionName)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	plain, err := aesencryption.Decrypt([]byte("key123"), cookie.Value)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, plain)
}
