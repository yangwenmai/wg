## wg 
[![Build Status](https://travis-ci.org/yangwenmai/wg.svg?branch=master)](https://travis-ci.org/yangwenmai/wg) [![Go Report Card](https://goreportcard.com/badge/github.com/yangwenmai/wg)](https://goreportcard.com/report/github.com/yangwenmai/wg)  [![Documentation](https://godoc.org/github.com/yangwenmai/wg?status.svg)](http://godoc.org/github.com/yangwenmai/wg) [![Coverage Status](https://coveralls.io/repos/github/yangwenmai/wg/badge.svg?branch=master)](https://coveralls.io/github/yangwenmai/wg?branch=master) [![GitHub issues](https://img.shields.io/github/issues/yangwenmai/wg.svg)](https://github.com/yangwenmai/wg/issues) [![license](https://img.shields.io/github/license/yangwenmai/wg.svg?maxAge=2592000)](https://github.com/yangwenmai/wg/LICENSE) [![Release](https://img.shields.io/github/release/yangwenmai/wg.svg?label=Release)](https://github.com/yangwenmai/wg/releases)

用 Go 写的一个简单的命令行天气预报小助手。

## Usage ##

1. clone 到本地 `git clone https://github.com/yangwenmai/wg.git`
2. 当前目录编译 `go build -o wg wg.go types.go utils.go`
3. `./wg`

或者

`go build -o wg wg.go types.go utils.go && ./wg`

你也可以直接 `go install`，将 wg 安装到你的GOBIN目录下，你可以直接在命令行使用 wg 命令。

## 参考资料 ##

1. 天气API接口 [http://www.sojson.com/api/weather.html](http://www.sojson.com/api/weather.html)

## 赞助我

![微信支付](./docs/wxpay.jpg)
