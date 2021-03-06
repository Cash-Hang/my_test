package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var (
	aciton       string
	filename     string
	h            string
	deviceRegexp = regexp.MustCompile("[a-f0-9]+")
)

func main() {
	//flag命令行
	flag.StringVar(&h, "help", "", "this help")
	flag.StringVar(&aciton, "a", "", "操作: add|del") //
	flag.StringVar(&filename, "f", "", "文件名：x.csv") //
	flag.Parse()

	if aciton != "add" && aciton != "del" {
		fmt.Println("action is not in [add,del] !")
		return
	}
	if filename == "" {
		fmt.Println("filename is null")
		return
	}
	// f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, os.ModeAppend)
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("打开文件" + filename + "失败：" + err.Error())
		return
	}
	// defer f.Close()

	// writer := csv.NewReader(f)
	// str, err := writer.ReadAll()
	// if err != nil {
	// 	log.Println("读取文件" + filename + "失败：" + err.Error())
	// 	return
	// }
	conStr := string(contents)
	fmt.Printf("%#v", conStr)
	str := strings.Split(strings.TrimSpace(conStr), "\r\n")
	if len(str) > 0 {
		//post请求
		// dataList := make([]string, 0)
		// for _, s := range str {
		// 	if len(s) == 1 {
		// 		if deviceRegexp.MatchString(s[0]) && !strings.Contains(s[0], " ") {
		// 			dataList = append(dataList, strings.TrimSpace(s[0]))
		// 		}
		// 		// if strings.HasPrefix(s[0], "00124b") && strings.Contains(s[0], substr string){
		// 		// }
		// 	}
		// }
		// fmt.Println(dataList)
		data := map[string][]string{
			// "deviceList": dataList,
			"deviceList": str,
		}
		fmt.Printf("%#v\n", data)
		url := "http://xxx/v1/backlist/" + aciton + "?key=xxx"
		bytedata, err := json.Marshal(data)
		if err != nil {
			log.Println("json.Marshal失败：" + err.Error())
			return
		}
		resp, err := http.Post(url, "application/json", bytes.NewReader(bytedata))
		if err != nil {
			return
		}
		d, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}
		fmt.Println(string(d))
	} else {
		log.Println("获取设备列表为空！")
		return
	}

}

// ---------------------------Usage-------------------------------
//./main -action=add -filename=gulou_devicelist.csv

//添加到黑名单，csv文件只有address(一列)
// ./main -a add -f 鼓楼大街店_copy.csv
//从黑名单移除数据
// ./main -a del -f 鼓楼大街店_copy.csv

//----------------链接-------------------
//golang读取文件
//https://www.jianshu.com/p/7790ca1bc8f6
