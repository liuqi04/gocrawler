//获取虎扑nba版块的每日热帖链接
package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

const (
	URL1 = "https://bbs.hupu.com/all-nba"
	URL2 = "https://bbs.hupu.com"
)

var (
	regExp1 = regexp.MustCompile("NBA论坛热帖.+b50000\">球队分区</span>")
	regExp2 = regexp.MustCompile("主版</span>.+</ul>")
	regExp3 = regexp.MustCompile("<ul>.+</ul>")

	linkReg = regexp.MustCompile("/(\\d*?).html")

	result []string
)

func convToHTML(links []string) template.HTML {
	result := "<h1>虎扑nba每日热帖：</h1>"
	result += "<ul>"
	for i := 0; i < len(links); i++ {
		result += ("<li><a href=\"" + URL2 + links[i] + "\">" + URL2 + links[i] + "</a></li>")
	}
	result += "</ul>"
	return template.HTML(result)
}

func main() {
	resp, err := http.Get(URL1)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	//将html内容转换为字符串并去除所有空格和换行符
	ress := strings.ReplaceAll(string(body), " ", "")
	ress = strings.ReplaceAll(ress, "\n", "")

	//获取所有匹配的目标新闻子链接
	result = linkReg.FindAllString(regExp3.FindString(regExp2.FindString(regExp1.FindString(ress))), -1)

	//输出内容
	http.HandleFunc("/", displayHTML)
	fmt.Println("server started...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func displayHTML(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	content := convToHTML(result)
	fmt.Fprintf(w, string(content))
}
