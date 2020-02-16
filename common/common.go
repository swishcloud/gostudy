package common

import (
	"bytes"
	"context"
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"golang.org/x/oauth2"
)

func StringLimitLen(str string, maxLen int) string {
	runes := []rune(str)
	if len(runes) > maxLen {
		runes = runes[:maxLen]
		return string(runes)

	}
	return str
}
func Md5Check(hashedStr string, plain string) bool {
	return Md5Hash(plain) == hashedStr
}
func Md5Hash(plain string) string {
	sb := []byte(plain)
	b := md5.Sum(sb)
	return hex.EncodeToString(b[:])
}

type RestApiClient struct {
	client *http.Client
}

func NewRestApiClient(skip_tls_verify bool) *RestApiClient {
	rac := new(RestApiClient)
	rac.client = new(http.Client)
	tlsConfig := tls.Config{}
	tlsConfig.InsecureSkipVerify = skip_tls_verify
	rac.client.Transport = &http.Transport{TLSClientConfig: &tlsConfig, DisableKeepAlives: true}
	return rac
}

func (rac *RestApiClient) Do(rar *RestApiRequest) (*http.Response, error) {
	if rar.token != nil {
		ts := rar.conf.TokenSource(context.Background(), rar.token)
		new_token, err := ts.Token()
		if err != nil {
			return nil, err
		}
		new_token.SetAuthHeader(rar.Request)
	}
	return rac.client.Do(rar.Request)
}

type RestApiRequest struct {
	Request *http.Request
	conf    *oauth2.Config
	token   *oauth2.Token
}

func NewRestApiRequest(method string, urlPath string, body []byte) *RestApiRequest {
	if req, err := http.NewRequest(method, urlPath, bytes.NewBuffer(body)); err != nil {
		panic(err)
	} else {
		req.Header = map[string][]string{
			"Content-Type": []string{"application/x-www-form-urlencoded"},
			"Accept":       []string{"application/json"},
		}
		rar := new(RestApiRequest)
		rar.Request = req
		return rar
	}
}
func (rar *RestApiRequest) UseToken(conf *oauth2.Config, token *oauth2.Token) *RestApiRequest {
	rar.conf = conf
	rar.token = token
	return rar
}
func (rar *RestApiRequest) SetAuthHeader(token *oauth2.Token) *RestApiRequest {
	token.SetAuthHeader(rar.Request)
	return rar
}
func ReadAsMap(r io.Reader) (map[string]interface{}, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	m := map[string]interface{}{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func FormatByteSize(n int64) (string, string) {
	unit := "kb"
	size := float64(n) / 1024
	if size >= 1024 {
		unit = "mb"
		size = size / 1024
	}
	s := fmt.Sprintf("%.2f", size)
	regex, err := regexp.Compile("0+$")
	if err != nil {
		panic(err)
	}
	s = regex.ReplaceAllString(s, "")
	s = strings.TrimSuffix(s, ".")
	return s, unit
}

type FileInfoWrapper struct {
	Fi   os.FileInfo
	Path string
}

func ReadAllFiles(path string, items *[]*FileInfoWrapper) error {
	fi, err := os.Stat(path)
	if err != nil {
		return err
	}
	*items = append(*items, &FileInfoWrapper{Fi: fi, Path: path})
	//check if it is directory
	if fi.IsDir() {
		if dir_items, err := ioutil.ReadDir(path); err != nil {
			return err
		} else {
			for i := 0; i < len(dir_items); i++ {
				err := ReadAllFiles(path+"/"+dir_items[i].Name(), items)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
