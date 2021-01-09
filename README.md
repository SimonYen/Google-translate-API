# Google-translate-API
Simple,no requirement.
***
## Install
`go get github.com/SimonYen/Google-translate-API`
***
## example
```
package main

import (
	"Google-translate-API/generater"
	"Google-translate-API/unmarshal"
	"fmt"
)

func main() {
	u := new(generater.URL)
	u.Ini(true, "ja", "hello") 
    /*If you in China,first parameter pls enter true.The second parameter is language code(iso-639-1 language code) that you wanana translated.The last parameter is text that you wanana translated.*/
	r := new(unmarshal.RawText)
	fmt.Println(r.Convert(u.Get())) //Convert json to string.
}
```
***
## Anthor
颜曾一(simonyen@foxmail.com)
Feel free to contract me if you need further help.