package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("time: %v\n", now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
	fmt.Printf("时间戳: %v\n", now.Unix())
	fmt.Printf("纳秒时间戳: %v\n", now.UnixNano())
	fmt.Printf("时间戳转时间: %v\n", time.Unix(now.Unix(),0))
	fmt.Printf("星期：%v\n", now.Weekday().String())
	fmt.Printf("加1小时: %v\n", now.Add(time.Hour))
	fmt.Printf("时间差: %v\n", now.Add(time.Hour).Sub(now))//now.add(-time)
	fmt.Printf("判断时间相等：%v\n", now.Equal(now))//会比较地点和时区
	fmt.Printf("判断是否在另一时间之前: %v\n", now.Add(time.Minute).Before(now))
	fmt.Printf("判断是否在另一时间之后: %v\n", now.Add(time.Minute).After(now))

	//定时器
	// ticker := time.Tick(time.Second)
	// for i := range ticker {
	// 	fmt.Println(i)
	// }

	//格式化，需要使用2006 年 1 月 2 号 15 点 04 分 05 秒。
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))//24小时
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))//12小时
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/02/01"))
	fmt.Println(now.Format("2006-02-01"))

	//解析字符串格式时间
	var t1 string = "2006-01-02 15:04:05"
	var timeStr string = "2019-12-12 15:22:12"
	fmt.Println(time.Parse(t1, timeStr))//缺少时区解释：UTC, 提供了时区偏移量信息时：匹配本地时区
	fmt.Println(time.ParseInLocation(t1, timeStr, time.Local))//缺少时区Location设置：loc, 提供了时区偏移量信息时：匹配loc
}