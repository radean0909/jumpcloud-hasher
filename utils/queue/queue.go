package queue

import (
	"os"
	"sync"

	"github.com/radean0909/jumpcloud-hasher/utils/database"
)

type Queue struct {
	Channel      chan Job
	WorkerCount  int
	RunningState bool
	Process      func(*Job) *Job
	Completed    int64
	wg           sync.WaitGroup
}

type Job struct {
	ID          int64
	DB          *database.DBTable
	In          string
	Out         string
	ProcessTime int64
}

func NewQueue(size, workers int, process func(*Job) *Job) (*Queue, error) {
	return &Queue{
		Channel:      make(chan Job, size),
		WorkerCount:  workers,
		RunningState: true,
		Process:      process,
		Completed:    0,
	}, nil
}

func NewJob(id int64, db *database.DBTable, in string) *Job {
	return &Job{
		ID:          id,
		DB:          db,
		In:          in,
		Out:         "",
		ProcessTime: 0,
	}
}

// Add tries to add a job, but fails if the queue is full or the server is shutting down
func (q *Queue) Add(j *Job) bool {
	if q.RunningState == false {
		return false
	}
	select {
	case q.Channel <- *j:
		return true
	default:
		return false
	}
}

func (q *Queue) Worker() {
	defer q.wg.Done()
	for j := range q.Channel {
		q.Process(&j)
		q.Completed++
		// Update the database with the hash
		j.DB.Update(j.ID, j.Out)

		// Update the processing times using cumulative moving average
		j.DB.Meta.Hashtime = (j.DB.Meta.Hashtime*float32(q.Completed-1) + float32(j.ProcessTime)) / float32(q.Completed)
	}
}

func (q *Queue) Start() {
	q.wg.Add(q.WorkerCount)
	for i := 0; i < q.WorkerCount; i++ {
		go q.Worker()
	}
}

func (q *Queue) Shutdown() {
	close(q.Channel)
	q.wg.Wait()
	os.Exit(1)
}
