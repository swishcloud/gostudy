package common

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

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
func SendRestApiRequest(method string, access_token string, urlPath string, body []byte, skip_tls_verify bool) []byte {
	headers := map[string][]string{
		"Content-Type":  []string{"application/x-www-form-urlencoded"},
		"Accept":        []string{"application/json"},
		"Authorization": []string{"Bearer " + access_token},
	}
	req, err := http.NewRequest(method, urlPath, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	req.Header = headers

	tlsConfig := tls.Config{}
	tlsConfig.InsecureSkipVerify = skip_tls_verify
	client := &http.Client{Transport: &http.Transport{TLSClientConfig: &tlsConfig}}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return b
}

type RestApiClient struct {
	client  *http.Client
	request *http.Request
}

func NewRestApiClient(method string, urlPath string, body []byte, skip_tls_verify bool) *RestApiClient {
	rac := new(RestApiClient)
	rac.client = http.DefaultClient
	if skip_tls_verify {
		tlsConfig := tls.Config{}
		tlsConfig.InsecureSkipVerify = skip_tls_verify
		rac.client.Transport = &http.Transport{TLSClientConfig: &tlsConfig}
	}
	if req, err := http.NewRequest(method, urlPath, bytes.NewBuffer(body)); err == nil {
		req.Header = map[string][]string{
			"Content-Type": []string{"application/x-www-form-urlencoded"},
			"Accept":       []string{"application/json"},
		}
		rac.request = req
	} else {
		panic(err)
	}
	return rac
}
func (rac *RestApiClient) UseToken(conf *oauth2.Config, token *oauth2.Token) *RestApiClient {
	c := conf.Client(oauth2.NoContext, token)
	rac.client = c
	return rac
}

func (rac *RestApiClient) SetHeader(key, value string) *RestApiClient {
	rac.request.Header.Set(key, value)
	return rac
}

func (rac *RestApiClient) Do() (*http.Response, error) {
	return rac.client.Do(rac.request)
}
func ReadAsMap(r io.Reader) map[string]interface{} {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	m := map[string]interface{}{}
	json.Unmarshal(b, &m)
	return m
}
