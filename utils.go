package main

import (
	"fmt"
	"github.com/modood/table"
	"io/ioutil"
	"net/http"

	"time"
)

const (
	// APIURL 天气接口API
	APIURL = "http://www.sojson.com/open/api/weather/json.shtml?city="
)

type WeatherToday struct {
	Date    string
	City    string
	Wendu   string
	Quality string
	Shidu   string
	Pm25    float64
	Pm10    float64
	Notice  string
}
type WeatherNotToday struct {
	Date      string
	City      string
	Tianqi    string
	Wendu     string
	Aqi       string
	Fengliang string
	Sunrise   string
	Sunset    string
	Notice    string
}

// Print 打印输出
func Print(day string, r Response) {
	// fmt.Println("城市:", r.City)
	if day == "今天" {
		// fmt.Println("温度:", r.Data.WenDu)
		// fmt.Println("空气质量:", r.Data.Quality)
		// fmt.Println("湿度:", r.Data.Shidu)
		// fmt.Println("PM25:", r.Data.Pm25)
		// fmt.Println("PM10:", r.Data.Pm10)
		// fmt.Println("温馨提示:", r.Data.Notice)
		w := WeatherToday{}
		w.Date = day
		w.City = r.City
		w.Wendu = r.Data.WenDu
		w.Quality = r.Data.Quality
		w.Shidu = r.Data.Shidu
		w.Pm25 = r.Data.Pm25
		w.Pm10 = r.Data.Pm10
		w.Notice = r.Data.Notice
		ws := []WeatherToday{w}
		table.Output(ws)
	} else if day == "昨天" {
		w := WeatherNotToday{}
		w.Date = r.Data.Yesterday.Date
		w.City = r.City
		w.Tianqi = r.Data.Yesterday.Type
		w.Wendu = r.Data.Yesterday.Low + r.Data.Yesterday.High
		w.Aqi = fmt.Sprintf("%v", r.Data.Yesterday.Aqi)
		w.Fengliang = r.Data.Yesterday.Fx + r.Data.Yesterday.Fl
		w.Sunrise = r.Data.Yesterday.Sunrise
		w.Sunset = r.Data.Yesterday.Sunset
		w.Notice = r.Data.Notice
		ws := []WeatherNotToday{w}
		table.Output(ws)
	} else if day == "预测" {
		fmt.Println("====================================")
		ws := []WeatherNotToday{}
		for _, item := range r.Data.Forecast {
			w := WeatherNotToday{}
			w.Date = item.Date
			w.City = r.City
			w.Tianqi = item.Type
			w.Wendu = item.Low + item.High
			w.Aqi = fmt.Sprintf("%v", item.Aqi)
			w.Fengliang = item.Fx + item.Fl
			w.Sunrise = item.Sunrise
			w.Sunset = item.Sunset
			w.Notice = item.Notice
			ws = append(ws, w)
		}
		table.Output(ws)
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
