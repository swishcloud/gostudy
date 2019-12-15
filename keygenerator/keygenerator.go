package keygenerator

import (
	"crypto/rand"
	"regexp"
)

func NewKey(length int, requireUpperCaseLetters, requireLowerCaseLetters, requireSpecialSymbols bool) string {
	patterns := map[string]bool{}
	patterns[".*[\\d]"] = true
	patterns[".*[A-Z]"] = requireUpperCaseLetters
	patterns[".*[a-z]"] = requireLowerCaseLetters
	patterns[".*[^A-Za-z\\d]"] = requireSpecialSymbols
	for {
		var pwdbytes []byte
		generatePwd(&pwdbytes, length)
		pwdstr := string(pwdbytes)
		matched := true
		for pattern, required := range patterns {
			m, err := regexp.MatchString(pattern, pwdstr)
			if err != nil {
				panic(err)
			}
			if m != required {
				matched = false
				break
			}
		}
		if matched {
			return pwdstr
		}
	}
}

func generatePwd(pwdbytes *[]byte, length int) {
	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	for _, value := range b {
		if value >= 33 && value <= 126 {
			if len(*pwdbytes) != length {
				*pwdbytes = append(*pwdbytes, value)
			} else {
				return
			}
		}
	}
	generatePwd(pwdbytes, length)
}
