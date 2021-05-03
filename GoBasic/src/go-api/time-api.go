package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	date_api()
}

func date_api(){

	now:=time.Now()
	fmt.Println("now=",now)

	var year,month,day=now.Date();
	fmt.Printf("year=%d,month=%d,day=%d\n",year,month,day)

	var hour,minite,second=now.Clock();
	fmt.Printf("hour=%d,minite=%d,second=%d\n",hour,minite,second)

	fmt.Printf("timestamp=%d\n",now.Unix())//时间截，从1970年1月1日0时开始到这个时间的秒数

	//afterAdd:=now.Add(time.Minute)//加时分秒
	afterAdd:=now.AddDate(0,0,-3)//加年月日
	fmt.Printf("afterAdd=%s\n",afterAdd)
	fmt.Println(now.Before(afterAdd))

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())

	var iso_time_format string="2006-01-02 15:04:05" //看文档一定要是2006年01月02日15:04:05来定义格式,而不是%Y或yyyy的格式
	//很容易看错了,是个坑,按1,2,3(下午),4,5,6(年)来记忆
	//date->string
	fmt.Printf(t.Format(iso_time_format))

	//string->date
	str:="2009-11-10 23:00:00"
	beginTime,err := time.Parse(str,iso_time_format);
	if err!=nil {
		fmt.Printf("error=%s\n", err)
	}else {
		fmt.Printf("beginTime=%s\n", beginTime)
	}

	rand.Seed(time.Now().Unix())
	randNum:=10+rand.Intn(10)*3//Intn(10)取0-9
	fmt.Println("10-30 random=%d",randNum)

	//time.Sleep(time.Duration(randNum))
	fmt.Println("3秒后退出")
	time.Sleep(3*time.Second)//暂时3秒


	//定时器,只触发一次
	fmt.Println("定时器,now=",time.Now());
	mytimer:=time.NewTimer(3*time.Second)
	go func(){
		ch1:=mytimer.C //C 是一个 chann
		//ch2:=time.After(3*time.Second) //源码就是 return NewTimer(d).C
		fmt.Println(<-ch1) //会阻塞3秒,显示3秒后的时间
	}()

	time.Sleep(2*time.Second)
	var ok bool =mytimer.Stop();//可以提前停止
	if ok {
		fmt.Println("成功取消定时器");
	}
	//---ticker 多次触发定时
	timer2 := time.NewTicker( time.Duration(time.Second*2))
	defer timer2.Stop()
	for {
		<- timer2.C
		fmt.Println("这用for每隔2秒执行一次")
	}

}