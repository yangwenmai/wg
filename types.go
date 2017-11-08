package main

// Response 接口的响应
type Response struct {
	Date    string `json:"date"`    // 日期
	Message string `json:"message"` // 请求响应消息
	Status  int    `json:"status"`  // 状态码：成功状态为200 ，失败为非200
	City    string `json:"city"`    // 城市
	Count   int32  `json:"count"`   // count
	Data    Data   `json:"data"`    // 数据
}

// Data 数据
type Data struct {
	Shidu     string        `json:"shidu"`     // 湿度
	Pm25      float64       `json:"pm25"`      // pm25
	Pm10      float64       `json:"pm10"`      // pm10
	Quality   string        `json:"quality"`   // 质量
	WenDu     string        `json:"wendu"`     // 温度
	Notice    string        `json:"ganmao"`    // 温馨提示：感冒指数
	Yesterday WeatherInfo   `json:"yesterday"` // 昨天
	Forecast  []WeatherInfo `json:"forecast"`  // 预测
}

// WeatherInfo 天气数据
type WeatherInfo struct {
	Date    string  `json:"date"`    // 日期
	Sunrise string  `json:"sunrise"` // 日出
	High    string  `json:"high"`    // 最高温
	Low     string  `json:"low"`     // 最低温
	Sunset  string  `json:"sunset"`  // 日落
	Aqi     float32 `json:"aqi"`     // AQI
	Fx      string  `json:"fx"`      // 风向
	Fl      string  `json:"fl"`      // 风力
	Type    string  `json:"type"`    // 天气
	Notice  string  `json:"notice"`  // 温馨提示
}
