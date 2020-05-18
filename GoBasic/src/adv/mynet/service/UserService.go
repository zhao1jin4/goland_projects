package service

import (
	"adv/mynet/bean"
	"strconv"
)

type UserService struct {
}

func (s UserService) QueryUser(request bean.MyRequest  ,resp *bean.MyResponse )  error{
	newName:=request.Name +"1"
	newId:=request.Id+1
	var iso_time_format string="2006-01-02 15:04:05"
	newBirthday:=request.Birthday.AddDate(0,0,1).Format(iso_time_format);
	resp.RespCode=200
	resp.RespMsg="成功"
	resp.ExtJson="{userId:"+strconv.Itoa(newId)+",userName:"+newName+",birthday:"+newBirthday+"}"
	return nil
}