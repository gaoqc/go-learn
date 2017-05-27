package main

import (
	// "encoding/xml"
	"fmt"
	// "io/ioutil"
	"net/http"
	// "os"
	"regexp"
	"strings"
)

func print(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fmt.Println(req.Form)
	fmt.Println("path", req.URL.Path)
	fmt.Println("schem", req.URL.Scheme)
	fmt.Println(req.Form["url_long"])
	for k, v := range req.Form {
		fmt.Println("key:", k, ",val:", strings.Join(v, ""))
	}
	fmt.Fprintln(resp, "hello gaoqc")
}

type Servers struct {
	// XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []Server `xml:"server"`
	Description string   `xml:",innerxml"`
}
type Server struct {
	// XMLName    xmldd.Name `xml:"server"`
	ServerName string `xml:"serverName"`
	ServerIp   string `xml:"serverIp"`
}

func main() {
	str := "I am learning Go language"
	re, err := regexp.Compile("lang.*")
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	fmt.Println("orgin str:", str)
	fmt.Println("replace:", re.ReplaceAll(str, []byte("1111")))
}
