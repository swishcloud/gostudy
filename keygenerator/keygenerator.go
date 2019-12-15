package keygenerator

import (
	"crypto/rand"
	"errors"
	"regexp"
)

func NewKey(length int, exculdeDigits, excludeUpperCaseLetters, excludeLowerCaseLetters, excludeSpecialSymbols bool) (string, error) {
	if exculdeDigits == true && excludeUpperCaseLetters == true && excludeLowerCaseLetters == true && excludeSpecialSymbols == true {
		return "", errors.New("can not exclude all type")
	}
	m := func(str string) bool {
		patterns := map[string]bool{}
		patterns[".*[\\d]"] = exculdeDigits
		patterns[".*[A-Z]"] = excludeUpperCaseLetters
		patterns[".*[a-z]"] = excludeLowerCaseLetters
		patterns[".*[^A-Za-z\\d]"] = excludeSpecialSymbols

		for pattern, exclude := range patterns {
			m, err := regexp.MatchString(pattern, str)
			if err != nil {
				panic(err)
			}
			if exclude && m {
				return false
			}
		}
		return true
	}

	var b []byte
	generatePwd(&b, length, m)
	key := string(b)
	return key, nil
}

func generatePwd(pwdbytes *[]byte, length int, m func(str string) bool) {
	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	for _, v := range b {
		if len(*pwdbytes) == length {
			return
		} else if v >= 33 && v <= 127 && m(string(v)) {
			*pwdbytes = append(*pwdbytes, v)
		}
	}
	generatePwd(pwdbytes, length, m)
}
