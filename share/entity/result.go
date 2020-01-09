package entity

import (
	"encoding/json"
	"fmt"
	"main/share"
)

type result struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Content interface{} `json:"content"`
}

type Content interface{}

func (r *result) SetCode(code int) *result {
	r.Code = code
	return r
}

func (r *result) SetMsg(msg string) *result {
	r.Msg = msg
	return r
}

func (r *result) SetContent(content Content) *result {
	r.Content = content
	return r
}

func (r *result) Result() string {
	j, err := json.Marshal(r)
	if err != nil {
		share.Log.Error(err)
	}
	fmt.Println(r)
	return string(j)
}

func Result(code int, msg string, content Content) string {
	data := result{
		Code: code,
		Msg:  msg,
	}
	if content != nil {
		data.Content = content
	}
	j, err := json.Marshal(data)
	if err != nil {
		share.Log.Error(err)
	}
	fmt.Println(data)
	return string(j)
}
