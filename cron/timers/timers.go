package timers

import (
	"cron/jobs"
	"cron/loggerx"
	"github.com/robfig/cron/v3"
	"sync"
)

type (
	jobMeta struct {
		jid cron.EntryID
		job *jobs.SyncJob
	}

	SyncTimer struct {
		c      *cron.Cron
		jobMap sync.Map
	}
)

func MustNewSyncTimer() *SyncTimer {
	return &SyncTimer{
		c: cron.New(),
	}
}

func (t *SyncTimer) AddJob(job *jobs.SyncJob) error {
	jid, err := t.c.AddJob(job.TimeSpec, cron.NewChain(cron.SkipIfStillRunning(loggerx.DefaultLogger),
		cron.Recover(loggerx.DefaultLogger)).Then(job))
	if err != nil {
		return err
	}

	t.jobMap.Store(job.Identifier(), jobMeta{
		jid: jid,
		job: job,
	})
	return nil
}

func (t *SyncTimer) RemoveJob(tid, pid string) {
	id := jobs.NewSyncIdentifier(tid, pid)

	jm, ok := t.jobMap.Load(id)
	if !ok {
		return
	}
	t.c.Remove(jm.(jobMeta).jid)
	t.jobMap.Delete(id)
}

func (t *SyncTimer) FetchJob(tid, pid string) *jobs.SyncJob {
	id := jobs.NewSyncIdentifier(tid, pid)

	jm, ok := t.jobMap.Load(id)
	if !ok {
		return nil
	}

	return jm.(jobMeta).job
}

func (t *SyncTimer) Count() int {
	return len(t.c.Entries())
}

func (t *SyncTimer) Start() {
	t.c.Start()

	<-make(chan struct{}, 1)
}
