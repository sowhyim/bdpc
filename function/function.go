package function

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetNumber(in []string, mode int) []string {
	var info bd_index_info
	var Number []string
	v, err := ioutil.ReadFile("c:/users/hemin/Desktop/src/zaqizaba/bd_index_info.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(v, &info)
	for _, val := range in {
		if info.Provinces[val] == "" {
			fmt.Println("输入的省份信息有误")
			time.Sleep(time.Second * 10)
		}
		Number = append(Number, info.Provinces[val])
	}
	fmt.Println(Number)
	//判断是否需要按城市获取数据
	if mode == 1 {
		var change []string
		Number, change = change, Number
		for _, val := range change {
			for _, valu := range info.CityShip[val] {
				if valu["value"] == "" {
					fmt.Println("输入的城市信息有误")
					time.Sleep(time.Second * 10)
				}
				Number = append(Number, valu["value"])
			}
		}
	}

	return Number
}

func GetDate(in []string) (string, string) {
	if len(in) != 3 {
		return "", "日期输入错误，请重新输入"
	}
	var date string
	date = in[0] + "-" + in[1] + "-" + in[2]
	layout := "2006-01-02"
	_, err := time.Parse(layout, date)
	if err != nil {
		return "", "月份或者日期超出范围"
	}
	return date, ""
}

func GetScan(in string) []string {
	return strings.Split(in, " ")
}

func GetUserScan() ([]string, []string, string, string, int) {
	var jingdian, Provinces []string
	var Provincesmode = 0

	fmt.Printf("请输入需要查询的关键字(多个关键字用空格隔开）：")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	for _, val := range GetScan(input.Text()) {
		jingdian = append(jingdian, val)
	}

	fmt.Printf("请输入需要查询的省份(多个关键字用空格隔开）：")
	input.Scan()
	for _, val := range GetScan(input.Text()) {
		Provinces = append(Provinces, val)
	}

	fmt.Printf("如果需要查询到城市，请输入Y，否则输入N：")
	input.Scan()
	if input.Text() == "y" || input.Text() == "Y" {
		Provincesmode = 1
	}

LABEL1:
	fmt.Printf("请输入需要查询的开始时间(列如：1995 01 13）：")
	input.Scan()
	var StartDate []string
	for _, val := range GetScan(input.Text()) {
		StartDate = append(StartDate, val)
	}
	startdate, err := GetDate(StartDate)
	if err != "" {
		goto LABEL1
	}

LABEL2:
	fmt.Printf("请输入需要查询的结束时间(列如：1995 01 13）：")
	var EndDate []string
	input.Scan()
	for _, val := range GetScan(input.Text()) {
		EndDate = append(EndDate, val)
	}
	enddate, err := GetDate(EndDate)
	if err != "" {
		goto LABEL2
	}

	return jingdian, Provinces, startdate, enddate, Provincesmode
}

func GetData(dataurl, word, area, startdate, enddate string,cookie *http.Cookie) Bddata {
	var data Bddata
	newarea, _ := strconv.Atoi(area)
	newdataurl := fmt.Sprintf(dataurl, newarea, word, startdate, enddate)
	client := &http.Client{}
	req, err := http.NewRequest("GET", newdataurl, nil)
	if err != nil {
		panic(err)
	}
	req.AddCookie(cookie)
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	val, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(val, &data)
	return data
}


func GetKey(unipidurl string ,data Bddata,cookie *http.Cookie) string {
	var key BdKey
	newunipidurl := fmt.Sprintf(unipidurl, data.Data.Uniqid)
	fmt.Println(newunipidurl)
	client := &http.Client{}
	req, err := http.NewRequest("GET", newunipidurl, nil)
	if err != nil {
		panic(err)
	}
	req.AddCookie(cookie)
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	val, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(val, &key)
	return key.Data
}