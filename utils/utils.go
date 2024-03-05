package utils

import (
    "crypto/md5"
    "encoding/json"
    "fmt"
    "strconv"
    "time"
    
    "goweb/errors"
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
    newTime, _ := time.Parse("2006-1-2 15:4:5", newTimeStr)
    return newTime
}

func Md5(original string) string {
    data := []byte(original)
    return fmt.Sprintf("%x", md5.Sum(data))
}

// 解析参数
func InterfaceToStruct(data, res interface{}) error {
    marshal, err := json.Marshal(data)
    if err != nil {
        return errors.New("解析参数失败", 400)
    }
    
    err = json.Unmarshal(marshal, res)
    if err != nil {
        return errors.New("解析参数失败", 400)
    }
    return nil
}
