package keygenerator

import (
	"crypto/rand"
	"regexp"
)

func NewKey(length int, requireUpperCaseLetters, requireLowerCaseLetters, requireSpecialSymbols bool) string {
	m := func(str string) bool {
		patterns := map[string]bool{}
		patterns[".*[\\d]"] = true
		patterns[".*[A-Z]"] = requireUpperCaseLetters
		patterns[".*[a-z]"] = requireLowerCaseLetters
		patterns[".*[^A-Za-z\\d]"] = requireSpecialSymbols

		for pattern, required := range patterns {
			m, err := regexp.MatchString(pattern, str)
			if err != nil {
				panic(err)
			}
			if m != required {
				return false
			}
		}
		return true
	}

	var b []byte
	generatePwd(&b, length, m)
	key := string(b)
	return key
}

func generatePwd(pwdbytes *[]byte, length int, m func(str string) bool) {
	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	for _, value := range b {
		if value >= 33 && value <= 126 {
			if len(*pwdbytes) == length {
				return
			} else if m(string(value)) {
				*pwdbytes = append(*pwdbytes, value)
			}
		}
	}
	generatePwd(pwdbytes, length, m)
}
