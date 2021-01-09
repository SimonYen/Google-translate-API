package generater

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//URL fuck golint.
type URL struct {
	prefixURL   string
	srcLanguage string
	desLanguage string
	text        string
}

//Ini initialize the struct.
func (u *URL) Ini(inChina bool, dl, text string) {
	if inChina == true {

		u.prefixURL = "http://translate.google.cn/translate_a/single?client=gtx&dt=t&dj=1&ie=UTF-8"
	} else {

		u.prefixURL = "http://translate.google.com/translate_a/single?client=gtx&dt=t&dj=1&ie=UTF-8"
	}
	u.srcLanguage = "auto"
	u.desLanguage = dl
	u.text = text

}

//Get get json.
func (u *URL) Get() []byte {
	parm := url.Values{}
	parm.Set("sl", u.srcLanguage)
	parm.Set("tl", u.desLanguage)
	parm.Set("q", u.text)
	U, err := url.ParseRequestURI(u.prefixURL)
	if err != nil {
		panic(err)
	}
	U.RawQuery = parm.Encode()
	fmt.Println(U.String())
	resp, err := http.Get(U.String())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return []byte(b)
}
