//获取虎扑nba版块的每日热帖链接
//写入当前目录下的news.txt
package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

const (
	URL        = "http://bbs.hupu.com/all-nba"
	outputPath = "./news.txt"
)

var (
	regExp1 = regexp.MustCompile("NBA论坛热帖.+b50000\">球队分区</span>")
	regExp2 = regexp.MustCompile("主版</span>.+</ul>")
	regExp3 = regexp.MustCompile("<ul>.+</ul>")

	linkReg = regexp.MustCompile("/(\\d*?).html")
)

func openFile(path string) *os.File {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0)
	if err != nil {
		if os.IsNotExist(err) {
			newFile, newErr := os.Create(path)
			if newErr != nil {
				log.Fatal(newErr)
			}
			return newFile
		} else {
			log.Fatal(err)
		}
	}
	return file
}

func writeToFile(links []string, file *os.File) {
	result := "虎扑nba每日热帖：\n"
	for i := 0; i < len(links); i++ {
		result += (URL + links[i] + "\n")
	}
	_, err := io.WriteString(file, result)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	resp, err := http.Get(URL)
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
	result := linkReg.FindAllString(regExp3.FindString(regExp2.FindString(regExp1.FindString(ress))), -1)

	//打开目标文件
	file := openFile(outputPath)
	defer func() {
		file.Close()
	}()

	//写入内容
	writeToFile(result, file)
}
