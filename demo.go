package main

import (
	"Google-translate-API/generater"
	"Google-translate-API/unmarshal"
	"fmt"
)

func main() {
	u := new(generater.URL)
	u.Ini(true, "ja", "hello")
	r := new(unmarshal.RawText)
	fmt.Println(r.Convert(u.Get()))
}
