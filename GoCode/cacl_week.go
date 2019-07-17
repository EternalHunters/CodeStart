//  获取上月最后一周到今天为止的周
package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println(getShowWeek())
}

func getShowWeek()map[string][2]string{
	result:=make(map[string][2]string)
	now:=time.Now()
	now = now.AddDate(0,-2,18)
	if now.Weekday() == 1 {
		now= now.AddDate(0,0,-1)
	}
	monthStr := strconv.Itoa(int(now.Month()))
	// now= now.AddDate(0,-1,0)
	// 本月第一天
	thisMonth := time.Date(now.Year(),now.Month(),1,0,0,0,0,now.Location())
	// 本月第一天是星期几
	monthFirstDay := int(thisMonth.Weekday())
	lastWeekStart := time.Now()
	if monthFirstDay >= 5{
		// 超过周四，本月第一周就是下一周
		lastWeekStart = thisMonth.AddDate(0, 0, 1-monthFirstDay)
	}else{
		lastWeekStart = thisMonth.AddDate(0, 0, -monthFirstDay-6)
	}
	result["上月最后一周"]=[2]string{lastWeekStart.Format("2006-01-02"), lastWeekStart.AddDate(0,0,6).Format("2006-01-02")}
	if lastWeekStart.AddDate(0,0,6).After(now) {
		return result
	}
	result[monthStr+"月第一周"]=[2]string{lastWeekStart.AddDate(0,0,7).Format("2006-01-02"),
		lastWeekStart.AddDate(0,0,13).Format("2006-01-02")}
	if lastWeekStart.AddDate(0,0,13).After(now){
		return result
	}
	result[monthStr+"月第二周"]=[2]string{lastWeekStart.AddDate(0,0,14).Format("2006-01-02"),
		lastWeekStart.AddDate(0,0,20).Format("2006-01-02")}
	if lastWeekStart.AddDate(0,0,20).After(now){
		return result
	}
	result[monthStr+"月第三周"]=[2]string{lastWeekStart.AddDate(0,0,21).Format("2006-01-02"),
		lastWeekStart.AddDate(0,0,27).Format("2006-01-02")}
	if lastWeekStart.AddDate(0,0,27).After(now){
		return result
	}
	result[monthStr+"月第四周"]=[2]string{lastWeekStart.AddDate(0,0,28).Format("2006-01-02"),
		lastWeekStart.AddDate(0,0,34).Format("2006-01-02")}
	if lastWeekStart.AddDate(0,0,34).After(now){
		return result
	}
	result[monthStr+"月第五周"]=[2]string{lastWeekStart.AddDate(0,0,35).Format("2006-01-02"),
		lastWeekStart.AddDate(0,0,41).Format("2006-01-02")}
	if lastWeekStart.AddDate(0,0,41).After(now){
		return result
	}
	return result
}