package utils

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"
)

func AddMonth(thisTime time.Time, duration int) time.Time {
	year, month := thisTime.Year(), int(thisTime.Month())
	thisTimeStr := thisTime.Format("2006-01-02 15:04:05")
	month += duration
	if month > 13 {
		month -= 12
		year += 1
	}
	newTimeStr := strconv.Itoa(year) + "-" + strconv.Itoa(month) + thisTimeStr[7:]
	newTime,_ := time.Parse("2006-1-2 15:4:5", newTimeStr)
	return newTime
}

func Md5(original string) string{
	data := []byte(original)
	return fmt.Sprintf("%x", md5.Sum(data))
}
