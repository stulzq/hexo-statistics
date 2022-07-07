package util

import (
	"fmt"
	"github.com/pkg/errors"
	"time"
)

// GetNextDay 获取下一次运行时间（day）
func GetNextDay(timeFormat string, forceNextDay bool) time.Time {
	// 当前时间
	currentTime := time.Now()
	// 下一次运行时间
	nextTime := time.Now()

	// 获取当前日期(day)
	currentTimeDate := currentTime.Format("2006-01-02")

	// 根据 timeFormat 计算出今天运行时间
	todayRunTime, initErr := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("%s %s", currentTimeDate, timeFormat), time.Local)
	if initErr != nil {
		panic(errors.Wrapf(initErr, "timeFormat %s parse to time error", timeFormat))
	}

	if forceNextDay {
		//明天运行
		nextTime = todayRunTime.AddDate(0, 0, 1)
	} else if currentTime.After(todayRunTime) || currentTime.Equal(todayRunTime) {
		//如果今天没运行，且大于等于目标时间点，也放到明天运行
		nextTime = todayRunTime.AddDate(0, 0, 1)
	} else {
		return todayRunTime
	}

	return nextTime
}

func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
