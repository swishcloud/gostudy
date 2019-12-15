package main

import (
	"log"
)

func main() {
	Execute()
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func Error(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
