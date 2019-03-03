package main

import (
	"fmt"
	"github.com/foomo/htpasswd"
	"log"
)

func main() {
	var filepath,newpwd string
	fmt.Print("please enter the file path:")
	fmt.Scanf("%s\n",&filepath)
	fmt.Print("new password:")
	fmt.Scanf("%s\n",&newpwd)
	users:=ReadUsers(filepath)
	SetUserPassword(filepath,users,newpwd)
	fmt.Printf("the passwords of all users located in file '%s' have been successfully changed to '%s'\n",filepath,newpwd)
	fmt.Printf("press the enter key to terminate the console screen")
	fmt.Scanln()
}
func ReadUsers(filepath string) ([]string) {
	passwords, err1 := htpasswd.ParseHtpasswdFile(filepath)
	if err1!=nil{
		log.Fatal(err1)
	}
	var users []string
	for k,_:=range passwords{
		users=append(users,k)
	}
	return users
}
func SetUserPassword(filepath string,users[]string,pwd string){
 	for  _,v:=range users{
		err := htpasswd.SetPassword(filepath, v, pwd, htpasswd.HashBCrypt)
		if err!=nil{
			log.Fatal(err)
		}
	}
}