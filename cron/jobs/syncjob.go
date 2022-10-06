package jobs

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
	"time"

	"golang.org/x/exp/rand"
)

var (
	SyncDelayInterval    = 10 * time.Minute
	SyncResponseInterval = 2 * time.Minute

	//MinMaxOffset  = 600 * time.Second
	//HourMaxOffset = 3 * MinMaxOffset
	MinMaxOffset  = 10 * time.Second
	HourMaxOffset = 3 * MinMaxOffset

	PREFIX = "{usersync-cron}"
	DELIM  = "@"
)

type SyncJob struct {
	TimeSpec, Tid, Pid, Type string
	Cmd                      func()
}

func MustNewSyncJob(interval time.Duration, tid, pid, syncType string, f func()) *SyncJob {
	// 浮动下时间，缓解并发，间隔小于 1 小时的浮动范围为 MinMaxOffset ，否则为 HourMaxOffset
	offset := MinMaxOffset
	if interval >= time.Hour {
		offset = HourMaxOffset
	}
	randSeed := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
	interval = interval + time.Duration(randSeed.Int63n(int64(offset)))
	timeSpec := fmt.Sprintf("@every %vs", interval.Seconds())

	logx.Infof("the timer job (tid: %v, pid: %v), interval: %v after rand: %s", tid, pid, interval, timeSpec)

	return &SyncJob{
		TimeSpec: timeSpec,
		Tid:      tid,
		Pid:      pid,
		Type:     syncType,
		Cmd:      f,
	}
}

func (j *SyncJob) Run() {
	//start := time.Now()
	//logx.Info("====== start Run() ======")
	j.Cmd()
	//logx.Infof("====== end Run() [%v] ======\n", time.Since(start))
}

func (j *SyncJob) Identifier() string {
	return NewSyncIdentifier(j.Tid, j.Pid)
}

func NewSyncIdentifier(tid, pid string) string {
	return NewUniqueKey(tid, pid)
}

func NewUniqueKey(tid, pid string) string {
	return PREFIX + tid + DELIM + pid
}

func ParseUniqueKey(key string) (string, string) {
	ss := strings.Split(strings.Split(key, PREFIX)[1], DELIM)
	return ss[0], ss[1]
}
