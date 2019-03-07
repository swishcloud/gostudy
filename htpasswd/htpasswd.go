package main

import (
	"bufio"
	"fmt"
	"github.com/foomo/htpasswd"
	"log"
	"os"
	"strings"
)

func main() {
	var filepath, newpwd string
	fmt.Print("please enter the file path:")
	fmt.Scanf("%s\n", &filepath)
	fmt.Print("new password:")
	fmt.Scanf("%s\n", &newpwd)
	users := ReadUsers(filepath)
	SetUserPassword(filepath, users, newpwd)
	fmt.Printf("the passwords of all users located in file '%s' have been successfully changed to '%s'\n", filepath, newpwd)
}
func ReadUsers(filepath string) ([]string) {
	var users []string
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	tempMap := map[string]int{}
	for scanner.Scan() {
		temp := strings.Split(scanner.Text(), ":")
		if len(temp) == 2 {
			user := temp[0]
			if tempMap[user] == 0 {
				users = append(users, user)
			}
			tempMap[user] = 1
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return users
}
func SetUserPassword(filepath string, users []string, pwd string) {
	rmErr := os.Remove(filepath)
	if rmErr != nil {
		log.Fatal(rmErr)
	}
	for _, v := range users {
		err := htpasswd.SetPassword(filepath, v, pwd, htpasswd.HashBCrypt)
		if err != nil {
			log.Fatal(err)
		}
	}
}
