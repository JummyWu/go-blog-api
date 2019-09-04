package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/gofrs/uuid"
)

/*
Message 返回格式
*/
type Message struct {
	Code   int         `json:"code"`
	Result string      `json:"result"`
	Data   interface{} `json:"data"`
}

/*
UID 随机生成id
*/
func UID() string {
	uid, err := uuid.NewV4()
	if err != nil {
		logs.Info(err)
	}
	return uid.String()
}
