package fileserver

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"regexp"

	"github.com/swishcloud/gostudy/aesencryption"
)

func index(w http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("templates/layout.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, map[string]string{"Body": "hello world"})
}
func FileList(w http.ResponseWriter, req *http.Request) {
	cookie := http.Cookie{Name: SessionName, Value: aesencryption.Encrypt([]byte("key123"), "hello world!"), Path: "/"}
	http.SetCookie(w, &cookie)

	var data FileListModel
	re := regexp.MustCompile(`/filelist/`)
	data.Path = re.ReplaceAllString(req.RequestURI, "")
	path := filepath.Join(config.FileLocation, data.Path)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		http.NotFound(w, req)
		return
	}
	data.Files = files

	tmpl, err := template.ParseFiles("templates/layout.html", "templates/filelist.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
}
