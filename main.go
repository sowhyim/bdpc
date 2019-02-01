package main

import (
	"bdpc/function"
	"fmt"
	"net/http"
)

var (
	dataurl   = "https://index.baidu.com/api/SearchApi/index?area=%d&word=%s&startDate=%s&endDate=%s"
	unipidurl = "http://index.baidu.com/Interface/api/ptbk?uniqid=%s"
	cookie    = &http.Cookie{Name: "BDUSS", Value: "d6SmVaY3JncXV0bjhwR1hoWkc4MjhWVFlIaThSeTZTaThGZEJ2Yi05QVE5R3hjQVFBQUFBJCQAAAAAAAAAAAEAAABGQ7LLc29Cb3lfc29naXJsAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABBnRVwQZ0VcdF"}
)

func Decrypt(key string,data function.Bddata){

	for _,val := range data.Data.UserIndexes{
		var a []uint8
		var alldata []uint8
		for i:=0;i<len(key)/2;i++{
			a[key[i]]= key[len(key)/2+i]
		}
		for i:=0;i<len(val.All.Data);i++{
			alldata=append(alldata,a[val.All.Data[i]])
		}
		fmt.Println(alldata)
	}
}

func main() {
	for {
		jingdian, Provinces, startdate, enddate, Provincesmode := function.GetUserScan()

		var area []string
		area = function.GetNumber(Provinces, Provincesmode)
		fmt.Println(jingdian, Provinces, startdate, enddate, Provincesmode, area)
		for _, jd := range jingdian {
			for _, p := range area {
				data := function.GetData(dataurl, jd, p, startdate, enddate, cookie)
				key:=function.GetKey(unipidurl,data,cookie)
				Decrypt(key,data)
			}
		}
	}
}
