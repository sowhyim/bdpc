package function

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/Luxurioust/excelize"
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
	v, err := ioutil.ReadFile("e:/nihao/src/sowhy/bdpc/bd_index_info.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(v, &info)
	for _, val := range in {
		if info.Provinces[val] == "" {
			fmt.Println("输入的省份信息有误")
			return nil
		}
		Number = append(Number, info.Provinces[val])
	}
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

func GetName(number string) string {
	var info map[string]string
	v, err := ioutil.ReadFile("e:/nihao/src/sowhy/bdpc/bdpc.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(v, &info)
	return info[number]
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

func GetUserScan() ([]string, []string, []string, string, string) {
	var jingdian []string

	fmt.Printf("请输入需要查询的关键字(多个关键字用空格隔开）：")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	for _, val := range GetScan(input.Text()) {
		jingdian = append(jingdian, val)
	}
LABEL:
	var Provinces []string
	var Provincesmode = 0
	fmt.Printf("请输入需要查询的省份(多个关键字用空格隔开，查全国数据直接打入全国）：")
	input.Scan()
	for _, val := range GetScan(input.Text()) {
		Provinces = append(Provinces, val)
	}

	fmt.Printf("如果需要查询到城市，请输入Y，否则输入N(查询全国数据，直接回车）：")
	input.Scan()
	if input.Text() == "y" || input.Text() == "Y" {
		Provincesmode = 1
	}
	area := GetNumber(Provinces, Provincesmode)
	if area == nil {
		goto LABEL
	}

	fmt.Println("---------------------特别注意---------------------------")
	fmt.Println("输入时间时，注意结束时间一定要大于开始时间，否则会报错！")
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

	return jingdian, area, Provinces, startdate, enddate
}

func GetData(dataurl, word, area, startdate, enddate string, cookie *http.Cookie) Bddata {
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

func GetKey(unipidurl string, data Bddata, cookie *http.Cookie) string {
	var key BdKey
	newunipidurl := fmt.Sprintf(unipidurl, data.Data.Uniqid)
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

func Decrypt(key, number string, xlsx *excelize.File, count int, data Bddata) {
	for _, val := range data.Data.UserIndexes {
		var a [1000]uint8
		var alldata []uint8
		for i := 0; i < len(key)/2; i++ {
			a[key[i]] = key[len(key)/2+i]
		}
		for i := 0; i < len(val.All.Data); i++ {
			alldata = append(alldata, a[val.All.Data[i]])
		}
		fmt.Println(GetName(number), string(alldata))
		data := strings.Split(string(alldata), ",")
		Insert(xlsx, count, GetName(number), data)
	}
}

func Insert(xlsx *excelize.File, count int, name string, data []string) {
	a := Axis(count) + "1"
	xlsx.SetCellValue("Sheet1", a, name)
	for i, val := range data {
		axis := Axis(count) + strconv.Itoa(i+2)
		xlsx.SetCellValue("Sheet1", axis, val)
	}
}

func Axis(row int) string {
	var a = [26]string{
		"Z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y",
	}
	var axis string
	var rowslice []int
	if row <= 0 {
		return "err"
	}
	for row > 0 {
		rowslice = append(rowslice, row%26)
		row = (row - 1) / 26
	}
	l := len(rowslice)
	for l > 0 {
		axis = axis + a[rowslice[l-1]]
		l--
	}
	return axis
}

func SetDate(a, b string) []string {
	start, _ := time.Parse("2006-01-02", a)
	end, _ := time.Parse("2006-01-02", b)
	var out []string
	out = append(out, a)
	for start != end {
		start = start.Add(time.Hour * 24)
		out = append(out, start.Format("2006-01-02"))
	}
	return out
}
