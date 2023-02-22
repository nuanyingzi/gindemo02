package models

import "time"

const (
	January = iota + 1
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

// UnixToTime 时间戳转换为日期
func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

// GetUnix 获取当前时间戳
func GetUnix() int64 {
	return int64(time.Now().Unix())
}

// GetDay 获取当前日期 例子：20230222
func GetDay() (date string) {

	year := time.Now().Format("2006")
	month := time.Now().Format("01")
	day := time.Now().Format("02")
	date = year + month + day
	return date
}
