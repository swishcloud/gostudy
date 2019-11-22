package main

import (
	"archive/zip"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func writeArchive(w *zip.Writer, archiveFolder string, directory string) {
	fileInfos, err := ioutil.ReadDir(directory)
	handleErr(err)
	for _, fileInfo := range fileInfos {
		full_path := directory + "\\" + fileInfo.Name()
		if fileInfo.IsDir() {
			newArchiveFolder := archiveFolder + fileInfo.Name() + "/"
			w.Create(newArchiveFolder)
			fmt.Println("writing directory:", full_path)
			writeArchive(w, newArchiveFolder, full_path)
		} else {
			writeArchiveFile(w, archiveFolder, full_path)
		}
	}

}

func writeArchiveFile(w *zip.Writer, archiveFolder string, file string) {
	f, err := w.Create(archiveFolder + filepath.Base(file))
	handleErr(err)
	fmt.Println("writing file:", file)
	b, err := ioutil.ReadFile(file)
	handleErr(err)
	_, err = f.Write(b)
	handleErr(err)
}

func handleErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func main() {
	compress_f := `C:\Users\allen\Desktop\compress-folder`
	archive_file_path := ""

	flag.StringVar(&compress_f, "f", "", "the folder/file to compress")
	flag.StringVar(&archive_file_path, "o", "", "the path of archive output file")

	flag.Parse()

	if compress_f == "" {
		fmt.Println("please specify folder/file path to compress using -f option")
		os.Exit(1)
	}

	stat, err := os.Stat(compress_f)
	handleErr(err)
	if archive_file_path == "" {
		archive_file_path = stat.Name() + ".zip"
	}

	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)

	f, err := os.Create(archive_file_path)
	handleErr(err)

	if !stat.IsDir() {
		writeArchiveFile(w, "", compress_f)
	} else {
		writeArchive(w, "", compress_f)
	}
	w.Close()
	n, err := f.Write(buf.Bytes())
	handleErr(err)
	err = f.Close()
	handleErr(err)
	fmt.Println("saved archive successfully with file size", n, "bytes at", archive_file_path)
}
