package main

import (
	"testing"
)

func TestRequest(t *testing.T) {
	r, err := Request(ApiURL + "深圳")
	if err != nil {
		t.Error(err)
	}
	t.Log(r)
}

func TestPrint(t *testing.T) {
	resp := Response{}
	resp.City = "深圳"
	resp.Status = 200
	resp.Date = "20171108"
	Print("今天", resp)
	Print("昨天", resp)
	Print("预测", resp)
	Print("未知", resp)
}
