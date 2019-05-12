package keygenerator

import (
	"crypto/rand"
	"fmt"
	"regexp"
)

func NewKey(length int)(string) {
	patterns := []string{".*[A-Z]", ".*[a-z]", ".*[^A-Za-z\\d]", ".*\\d"}
	for {
		var pwdbytes []byte
		generatePwd(&pwdbytes, length)
		pwdstr := string(pwdbytes)
		matched := true
		for _, value := range patterns {
			if m, err := regexp.MatchString(value, pwdstr); !m {
				if err != nil {
					println(err)
				}
				matched = false
				break
			}
		}
		if (matched) {
			return pwdstr
		}
	}
}

func generatePwd(pwdbytes *[]byte, length int) {
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
	generatePwd(pwdbytes, length)
}
