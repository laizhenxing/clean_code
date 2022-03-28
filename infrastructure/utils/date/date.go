package date

import (
	"time"
)

const (
	// GolangBornDateYear 年格式转化用到的常量
	GolangBornDateYear = "2006"
	// GolangBornDateMonth 月格式转化用到的常量
	GolangBornDateMonth = "2006-01"
	// GolangBornDate 日期格式转化用到的常量
	GolangBornDate = "2006-01-02"
	// GolangBornTime 时间转化是用到的常量
	GolangBornTime = "2006-01-02 15:04:05"
	// GolangBornMinute 时间转化是用到的常量
	GolangBornMinute = "2006-01-02 15:04"
	// ZeroDateTime 计算机起始时间字符串
	ZeroDateTime = "1970-01-01 00:00:00"
	// ZeroDate 计算机起始日期字符串
	ZeroDate = "1970-01-01"
)

//TToTime 时间戳转时间格式
func TToTime(t int64, timeTemplate string) string {
	return time.Unix(t, 0).Format(timeTemplate)
}

// FormatDate 时间格式化 目标格式"2006-01-02"
func FormatDate(tt time.Time) string {
	return tt.In(time.Local).Format(GolangBornDate)
}

// FormatDateTime 时间格式化 目标格式"2006-01-02 15:04:05"
func FormatDateTime(tt time.Time) string {
	return tt.In(time.Local).Format(GolangBornTime)
}

// FormatDateMinute 时间格式化 目标格式"2006-01-02 15:04"
func FormatDateMinute(tt time.Time) string {
	return tt.In(time.Local).Format(GolangBornMinute)
}

// GetTimeUnix 转为时间戳->秒数
func GetTimeUnix(t time.Time) int64 {
	return t.Unix()
}

//StrToDateTime 字符串转时间格式(为空时返回1970-01-01 00:00:00 的 time.Time)
func StrToDateTime(tStr string) (time.Time, error) {
	var et time.Time
	var err error
	if tStr != "" {
		et, err = time.ParseInLocation(GolangBornTime, tStr, time.Local)
	} else {
		et, err = time.ParseInLocation(GolangBornTime, ZeroDateTime, time.Local)

	}
	return et, err
}

//StrToDate 字符串转时间格式(为空时返回1970-01-01 的 time.Time)
func StrToDate(tStr string) (time.Time, error) {
	var et time.Time
	var err error
	if tStr != "" {
		et, err = time.ParseInLocation(GolangBornDate, tStr, time.Local)
	} else {
		et, err = time.ParseInLocation(GolangBornDate, ZeroDate, time.Local)

	}
	return et, err
}

//YearStrToDate 年字符串转时间格式(为空时返回1970 的 time.Time)
func YearStrToDate(tStr string) (time.Time, error) {
	var et time.Time
	var err error
	if tStr != "" {
		et, err = time.ParseInLocation(GolangBornDateYear, tStr, time.Local)
	} else {
		et, err = time.ParseInLocation(GolangBornDateYear, ZeroDate, time.Local)

	}
	return et, err
}

//MonthStrToDate 月字符串转时间格式(为空时返回1970-01 的 time.Time)
func MonthStrToDate(tStr string) (time.Time, error) {
	var et time.Time
	var err error
	if tStr != "" {
		et, err = time.ParseInLocation(GolangBornDateMonth, tStr, time.Local)
	} else {
		et, err = time.ParseInLocation(GolangBornDateMonth, ZeroDate, time.Local)

	}
	return et, err
}
