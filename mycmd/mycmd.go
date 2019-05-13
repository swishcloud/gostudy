package main

import (
	"fmt"
	"github.com/github-123456/gostudy/keygenerator"
	"os"
	"strconv"
)

func main() {
	fmt.Println(len(os.Args), os.Args)
	if len(os.Args) < 2 {
		return
	}
	commandType := os.Args[1]
	switch commandType {
	case "generatepwd":
		length := 16
		if len(os.Args) >= 3 {
			if l, err := strconv.Atoi(os.Args[2]); err == nil {
				length = l
			} else {
				fmt.Print("invalid argument:", os.Args[2])
			}

		}
		println(keygenerator.NewKey(length))
	default:
		println("invalid command:", commandType)
	}

}