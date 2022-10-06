package main

import (
	"cron/jobs"
	"cron/timers"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

func test01() {
	timeInterval := 2 * time.Second
	timer := timers.MustNewSyncTimer()
	start := time.Now()
	err := timer.AddJob(jobs.MustNewSyncJob(timeInterval, "tid", "pid", "ldap", func() {
		logx.Infof("[%v] === time:%v", time.Since(start), time.Now())
		time.Sleep(time.Second * 10)
	}))
	if err != nil {
		logx.Error(err)
		return
	}
	timer.Start()
}

func test02() {
	timeInterval := 2 * time.Second
	timer := timers.MustNewSyncTimer()
	start := time.Now()

	for i := 0; i < 10; i++ {
		tid := fmt.Sprintf("tid-%d", i)
		pid := fmt.Sprintf("pid-%d", i+1000)
		err := timer.AddJob(jobs.MustNewSyncJob(timeInterval, tid, pid, "ldap", func() {
			logx.Infof("[%d-%v] === tid:%s, pid:%s", timer.Count(), time.Since(start), tid, pid)
		}))
		if err != nil {
			logx.Error(err)
			return
		}
	}

	timer.Start()
}

func main() {
	logx.Info("cron start ...")
	//test01()
	test02()
}
