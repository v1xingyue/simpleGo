package main

import (
	"io/ioutil"
    "net/http"
	"net/url"
	"fmt"
	"os/exec"
	"strings"
	"bytes"
)

func RequestGet(url string,host string) string {
	client := new(http.Client)
    reg, err := http.NewRequest("GET",url, nil)
    if err != nil {
        fmt.Println("Error1:", err)
        return ""
    }
    //reg.Header.Set(`HTTP`, `1.1`)
    //reg.Header.Set(`Accept`, `*/*`)
    //reg.Header.Set(`Host`, host)
    //reg.Header.Set(`Connection`, `keep-alive`)
    //resp, err := client.Do(reg)
    defer resp.Body.Close()
    if err != nil {
        fmt.Println("Error2:", err)
        return ""
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return ""
    }
    fmt.Printf(string(body))
	return string(body)
}

func RequestPost(url_string string,post_params string) string {
	v , err := url.ParseQuery(post_params)
	if err != nil {
        return ""
    }
	resp, err := http.PostForm(url_string,v)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
        return ""
    }
    fmt.Println(string(body))
	return ""
}


func downloadFile(url string,downloadFile string, dirName string){
	cmd := exec.Command("/data0/myget012/bin/mytget","-d",dirName,"-f",downloadFile,url)
	cmd.Stdin = strings.NewReader("")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%#v",err)
	}
}

func main(){
	downloadFile("http://baidu.com","a.html","/data0/shareGo/src/leafApi")
}
