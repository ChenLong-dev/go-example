package timers

import (
	"cron/jobs"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestTimerJob_AddJob(t *testing.T) {
	Convey("AddJob测试\n", t, func() {
		timeInterval := 2 * time.Second
		timer := MustNewSyncTimer()
		start := time.Now()
		err := timer.AddJob(jobs.MustNewSyncJob(timeInterval, "tid", "pid", "ldap", func() {
			t.Logf("[%v] === time:%v", time.Since(start), time.Now())
		}))
		So(err, ShouldBeNil)
		timer.Start()
	})
}

func TestTimerJob_AddJobs(t *testing.T) {
	Convey("AddJobs测试\n", t, func() {
		timeInterval := 2 * time.Second
		timer := MustNewSyncTimer()
		start := time.Now()

		for i := 0; i < 10; i++ {
			tid := fmt.Sprintf("tid-%d", i)
			pid := fmt.Sprintf("pid-%d", i+1000)
			err := timer.AddJob(jobs.MustNewSyncJob(timeInterval, tid, pid, "ldap", func() {
				t.Logf("[%d-%v] === tid:%s, pid:%s", timer.Count(), time.Since(start), tid, pid)
			}))
			So(err, ShouldBeNil)
		}
		So(timer.Count(), ShouldEqual, 10)

		timer.Start()
	})
}

func TestTimerJob_SkipJob(t *testing.T) {
	Convey("SkipJob测试\n", t, func() {
		timeInterval := 2 * time.Second
		timer := MustNewSyncTimer()
		start := time.Now()
		err := timer.AddJob(jobs.MustNewSyncJob(timeInterval, "tid", "pid", "ldap", func() {
			t.Logf("[%v] === time:%v", time.Since(start), time.Now())
			time.Sleep(time.Second * 10)
		}))
		So(err, ShouldBeNil)
		timer.Start()
	})
}

func TestTimerJob_FetchJob(t *testing.T) {
	Convey("FetchJob测试\n", t, func() {
		timeInterval := 2 * time.Second
		timer := MustNewSyncTimer()
		start := time.Now()

		for i := 0; i < 10; i++ {
			tid := fmt.Sprintf("tid-%d", i)
			pid := fmt.Sprintf("pid-%d", i+1000)
			err := timer.AddJob(jobs.MustNewSyncJob(timeInterval, tid, pid, "ldap", func() {
				t.Logf("[%d-%v] === tid:%s, pid:%s", timer.Count(), time.Since(start), tid, pid)
			}))
			So(err, ShouldBeNil)
		}

		job1 := timer.FetchJob("tid-1", "pid-1001")
		So(job1, ShouldNotBeNil)

		job2 := timer.FetchJob("tid-1", "pid-1002")
		So(job2, ShouldBeNil)

		timer.Start()
	})
}

func TestTimerJob_DelJob(t *testing.T) {
	Convey("DelJob测试\n", t, func() {
		timeInterval := 2 * time.Second
		timer := MustNewSyncTimer()
		start := time.Now()

		for i := 0; i < 10; i++ {
			tid := fmt.Sprintf("tid-%d", i)
			pid := fmt.Sprintf("pid-%d", i+1000)
			err := timer.AddJob(jobs.MustNewSyncJob(timeInterval, tid, pid, "ldap", func() {
				t.Logf("[%d-%v] === tid:%s, pid:%s", timer.Count(), time.Since(start), tid, pid)
			}))
			So(err, ShouldBeNil)
		}

		go func() {
			time.Sleep(time.Second * 10)
			for i := 0; i < 5; i++ {
				tid := fmt.Sprintf("tid-%d", i)
				pid := fmt.Sprintf("pid-%d", i+1000)
				timer.RemoveJob(tid, pid)
				t.Logf("[x] [%d-%v] === tid:%s, pid:%s", timer.Count(), time.Since(start), tid, pid)
				time.Sleep(time.Second * 5)
			}
		}()

		timer.Start()
	})
}

func TestTimerJob_ModifyJob(t *testing.T) {
	Convey("DelJob测试\n", t, func() {
		timeInterval := 2 * time.Second
		timer := MustNewSyncTimer()
		start := time.Now()

		for i := 0; i < 10; i++ {
			tid := fmt.Sprintf("tid-%d", i)
			pid := fmt.Sprintf("pid-%d", i+1000)
			err := timer.AddJob(jobs.MustNewSyncJob(timeInterval, tid, pid, "ldap", func() {
				t.Logf("[%d-%v] === tid:%s, pid:%s", timer.Count(), time.Since(start), tid, pid)
			}))
			So(err, ShouldBeNil)
		}

		go func() {
			time.Sleep(time.Second * 10)
			for i := 0; i < 5; i++ {
				tid := fmt.Sprintf("tid-%d", i)
				pid := fmt.Sprintf("pid-%d", i+1000)
				timer.RemoveJob(tid, pid)
				t.Logf("[x] [%d-%v] === tid:%s, pid:%s", timer.Count(), time.Since(start), tid, pid)
				time.Sleep(time.Second * 5)
				tid1 := tid + "-modify"
				pid1 := pid + "-modify"
				err := timer.AddJob(jobs.MustNewSyncJob(timeInterval, tid1, pid1, "ldap", func() {
					t.Logf("[modify][%d-%v] === tid1:%s, pid1:%s", timer.Count(), time.Since(start), tid1, pid1)
				}))
				if err != nil {
					t.Error(err)
				}
				//So(err, ShouldBeNil)
			}
		}()

		timer.Start()
	})
}
