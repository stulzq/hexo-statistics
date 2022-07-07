package job

import (
	"github.com/stulzq/hexo-statistics/common/globalwait"
	"github.com/stulzq/hexo-statistics/logger"
	"github.com/stulzq/hexo-statistics/util"
	"time"
)

func RunAt(jobName string, timeFormat string, exec func()) {

	globalwait.Add(1)
	defer func() {
		logger.Infof("[Job][%s] job exit", jobName)
		globalwait.Done()
	}()

	interrupt := globalwait.NewInterrupt()
	nextTime := util.GetNextDay(timeFormat, false)
	timer := time.NewTimer(nextTime.Sub(time.Now()))
	logger.Infof("[Job][%s] job %s will running at %s", jobName, jobName, util.FormatTime(nextTime))

LOOP:
	for true {
		select {
		case <-timer.C:
			exec()
			nextTime = util.GetNextDay(timeFormat, true)
			timer.Reset(nextTime.Sub(time.Now()))
			logger.Infof("[Job][%s] job %s next running at %s", jobName, jobName, util.FormatTime(nextTime))
		case <-interrupt:
			break LOOP
		}
	}
}
