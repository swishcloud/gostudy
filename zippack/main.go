package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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
	compress_f := ""
	archive_file_path := ""

	flag.StringVar(&compress_f, "f", "", "the folder/file paths to compress,separated by semicolons")
	flag.StringVar(&archive_file_path, "o", "", "the path of archive output file")

	flag.Parse()

	if compress_f == "" {
		fmt.Println("please specify folder/file paths to compress which separated by semicolons using -f option")
		os.Exit(1)
	}

	if archive_file_path == "" {
		fmt.Println("please specify archive file output path using -o option")
		os.Exit(1)
	}

	f, err := os.Create(archive_file_path)
	handleErr(err)

	w := zip.NewWriter(f)

	paths := strings.Split(compress_f, ";")
	for i := 0; i < len(paths); i++ {
		stat, err := os.Stat(paths[i])
		handleErr(err)
		if stat.IsDir() {
			writeArchive(w, "", paths[i])
		} else {
			writeArchiveFile(w, "", paths[i])
		}
	}

	w.Close()
	err = f.Close()
	handleErr(err)
	fi, err := os.Stat(archive_file_path)
	handleErr(err)
	fmt.Println("saved archive successfully with file size", fi.Size(), "bytes at", archive_file_path)
}
