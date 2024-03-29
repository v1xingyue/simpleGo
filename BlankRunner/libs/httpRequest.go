package libs

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func RequestGet(url string, host string) string {
	client := new(http.Client)
	reg, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("HTTP GET ERROR 1 : ", err)
		return ""
	}
	//reg.Header.Set(`HTTP`, `1.1`)
	resp, err := client.Do(reg)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("HTTP GET ERROR 2 :", err)
		return ""
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(body)
}

func RequestPost(url_string string, post_params string) string {
	v, err := url.ParseQuery(post_params)
	if err != nil {
		fmt.Println("HTTP POST ERROR 1 :", err)
		return ""
	}
	resp, err := http.PostForm(url_string, v)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("HTTP POST ERROR 2 :", err)
		return ""
	}
	return string(body)
}

var email_send_time = make(map[string]int)

// send email
// /usr/bin/curl -s -d"body=test" "http://172.16.30.169/wsend.php?id=007&to=yanqing4@staff.sina.com.cn&title=qr-err-ip:10.73.12.23"
func SendAlarmEmail(to string, body string, title string) {
	url := fmt.Sprintf("http://172.16.30.169/wsend.php?id=007&to=%s&title=%s", to, title)
	if email_send_time[url] == 0 || 600 < (email_send_time[url]-int(time.Now().Unix())) {
		post_data_string := fmt.Sprintf("body=%s", body)
		RequestPost(url, post_data_string)
		email_send_time[url] = int(time.Now().Unix())
	}
}

//func main() {
//	//requestget("http://baidu.com","baidu.com")
//	requestpost("http://10.13.0.41/e.php", "url=baidu.com")
//}
