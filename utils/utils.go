package utils

import (
    "crypto/md5"
    "encoding/json"
    "fmt"
    "math"
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

// float64保留N为小数
func RetainDecimalNum(num float64, precisionList ...int) float64 {
    precision := 4
    if len(precisionList) != 0 {
        precision = precisionList[0]
    }
    return math.Round(num*math.Pow10(precision)) / math.Pow10(precision)
}

// SetBit 设置第N位为1,从右往左数
// 34, 3 => 38(100010, 100110)
// 34, 2 => 34(100010, 100010)
func SetBit(num, n int) int {
    return num | (1 << (n - 1))
}

// IsOne 判断第N位是否为1, 从右往左数
// 34, 3 => false 100010
// 35, 1 => true  100011
func IsOne(num int, n int) bool {
    return GetBit(num, n) == 1
}

// IsOne 判断第N位是否为0, 从右往左数
// 34, 3 => true 100010
// 35, 1 => false  100011
func IsZero(num int, n int) bool {
    return GetBit(num, n) == 0
}

// GetBit 获取第n位的值
// 34, 3 => 0 100010
// 35, 1 => 1 100011
func GetBit(num int, n int) int {
    return (num >> (n - 1)) & 1
}
