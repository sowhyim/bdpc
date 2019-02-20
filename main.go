package main

import (
	"bdpc/function"
	"fmt"
	"github.com/Luxurioust/excelize"
	"net/http"
	"strconv"
)

var (
	dataurl   = "https://index.baidu.com/api/SearchApi/index?area=%d&word=%s&startDate=%s&endDate=%s"
	unipidurl = "http://index.baidu.com/Interface/api/ptbk?uniqid=%s"
	cookie    = &http.Cookie{Name: "BDUSS", Value: "d6SmVaY3JncXV0bjhwR1hoWkc4MjhWVFlIaThSeTZTaThGZEJ2Yi05QVE5R3hjQVFBQUFBJCQAAAAAAAAAAAEAAABGQ7LLc29Cb3lfc29naXJsAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABBnRVwQZ0VcdF"}
)

func main() {
	for {
		jingdian, area, Provinces, startdate, enddate := function.GetUserScan()

		fmt.Println("输入的数据为", jingdian, Provinces, startdate, enddate)
		for _, jd := range jingdian {
			fmt.Println("开始抓取“", jd, "”的数据")
			xlsx := excelize.NewFile()
			xlsx.SetCellValue("Sheet1", "A1", jd)
			date := function.SetDate(startdate, enddate)
			datecount := 2
			for _, val := range date {
				axit := "A" + strconv.Itoa(datecount)
				xlsx.SetCellValue("Sheet1", axit, val)
				datecount++
			}

			count := 2
			for _, p := range area {
				data := function.GetData(dataurl, jd, p, startdate, enddate, cookie)
				key := function.GetKey(unipidurl, data, cookie)
				function.Decrypt(key, p, xlsx, count, data)
				count++
			}
			xlsx.SaveAs(jd + ".xlsx")
			fmt.Printf("--------------------------------------%s抓取完毕！--------------------------------------\n", jd)
		}
		fmt.Println("------------------------------------------抓取完毕！------------------------------------------")
		fmt.Println("------------------------------------------我是分割线------------------------------------------")
	}
}
