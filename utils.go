package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	// APIURL 天气接口API
	APIURL = "http://www.sojson.com/open/api/weather/json.shtml?city="
)

// Print 打印输出
func Print(day string, r Response) {
	fmt.Println("城市:", r.City)
	if day == "今天" {
		fmt.Println("温度:", r.Data.WenDu)
		fmt.Println("空气质量:", r.Data.Quality)
		fmt.Println("湿度:", r.Data.Shidu)
		fmt.Println("PM25:", r.Data.Pm25)
		fmt.Println("PM10:", r.Data.Pm10)
		fmt.Println("温馨提示:", r.Data.Notice)
	} else if day == "昨天" {
		fmt.Println("日期:", r.Data.Yesterday.Date)
		fmt.Println("天气:", r.Data.Yesterday.Type)
		fmt.Println("温度:", r.Data.Yesterday.Low, r.Data.Yesterday.High)
		fmt.Println("AQI:", r.Data.Yesterday.Aqi)
		fmt.Println("风量:", r.Data.Yesterday.Fx, r.Data.Yesterday.Fl)
		fmt.Println("日出:", r.Data.Yesterday.Sunrise)
		fmt.Println("日落:", r.Data.Yesterday.Sunset)
		fmt.Println("温馨提示:", r.Data.Yesterday.Notice)
	} else if day == "预测" {
		fmt.Println("====================================")
		for _, item := range r.Data.Forecast {
			fmt.Println("日期:", item.Date)
			fmt.Println("天气:", item.Type)
			fmt.Println("温度:", item.Low, item.High)
			fmt.Println("AQI:", item.Aqi)
			fmt.Println("风量:", item.Fx, item.Fl)
			fmt.Println("日出:", item.Sunrise)
			fmt.Println("日落:", item.Sunset)
			fmt.Println("温馨提示:", item.Notice)
			fmt.Println("====================================")
		}
	} else {
		fmt.Println("现在暂时无法获取到天气状况/(ㄒoㄒ)/~~")
	}
}

// Request URL 请求
func Request(url string) (string, error) {
	start := time.Now()
	response, err := http.Get(url)
	statRequest(url, start)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body), nil
}

func statRequest(url string, t time.Time) {
	// fmt.Printf("Request URL: %s, cost %v.\n", url, time.Now().Sub(t))
}
