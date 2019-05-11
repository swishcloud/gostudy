package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"regexp"
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
		patterns:=[]string{".*[A-Z]",".*[a-z]",".*[^A-Za-z\\d]",".*\\d"}
		length := 16
		if len(os.Args) >= 3 {
			if l, err := strconv.Atoi(os.Args[2]); err == nil {
				length = l
			} else {
				fmt.Print("invalid argument:", os.Args[2])
			}

		}
		for{
			var pwdbytes []byte
			generatepwd(&pwdbytes, length)
			pwdstr:=string(pwdbytes)
			matched:=true
			for _,value:=range patterns{
				if m,err:=regexp.MatchString(value,pwdstr);!m{
					if err!=nil{
						println(err)
					}
					matched=false
					break
				}
			}
			if(matched){
				fmt.Printf("[%s]",pwdstr)
				break
			}
		}
	default:
		println("invalid command:", commandType)
	}

}

func generatepwd(pwdbytes *[]byte, length int) {
	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	for _, value := range b {
		if (value >= 33 && value <= 126) {
			if len(*pwdbytes) != length {
				*pwdbytes = append(*pwdbytes, value)
			} else {
				return
			}
		}
	}
	generatepwd(pwdbytes, length)
}
