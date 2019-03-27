package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"time"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: false})
	rand.Seed(time.Now().UnixNano())
}

// Job is an interface which represents an idea of performing
// time sensitive form of processing.
type Job interface {
	Perform()
	GetID() int
}

// ShortJob is a struct, which main responsibility is to implement
// Job's interface behavior.
type ShortJob struct {
	ID   int
	Name string
}

// Perform is a *ShortJob receiver method, which simulates less time
// sensitive form of processing. It is a part of Job's interface contract.
func (sj *ShortJob) Perform() { time.Sleep(10 * time.Second) }

// GetID is a *ShortJob receiver method, which is a getter for short
// job ID attribute
func (sj *ShortJob) GetID() int { return sj.ID }

// LongJob is a struct, which main responsibility is to implement
// Job's interface behavior.
type LongJob struct {
	ID   int
	Name string
}

// Perform is a *LongJob receiver method, which fulfill Job's
// interface contract. It simulates more time sensitive form of
// processing.
func (lj *LongJob) Perform() { time.Sleep(15 * time.Second) }

// GetID is a *LongJob receiver method, which is a getter for long
// job ID attribute
func (lj *LongJob) GetID() int { return lj.ID }

func main() {
	log.Infoln("Starts program...")

	// Returns a new *WorkerPool and ensures that neither of
	// errors have occurred during initialization.
	pool, err := NewWorkerPool(3)

	if err != nil {
		log.Fatalf("Err: %s", err.Error())
		os.Exit(1)
	}

	// Triggers all of the workers within the pool. It listens
	// for jobs concurrently, possibly in parallel.
	pool.Run()

	// Creates simple closed loop to begin with job scheduling
	// in slightly random fashion (manner)
	for j := 0; j < 6; j++ {
		switch r := rand.Intn(2); r {
		case 0:
			if !pool.Schedule(&ShortJob{ID: j, Name: "Short Job"}) {
				fmt.Printf("Job (ID = %d) cannot be scheduled. Queue is full", j)
				continue
			}
		case 1:
			if !pool.Schedule(&LongJob{ID: j, Name: "Long Job"}) {
				fmt.Printf("Job (ID = %d) cannot be scheduled. Queue is full", j)
				continue
			}
		}
	}

	// Halts program and waits until all scheduled jobs will be done.
	pool.CloseAndWait()

	log.Infoln("Finishes the program ...")
	os.Exit(1)

}

// WorkerPool is a struct, which main responsibility is to represent
// grouping mechanism (abstraction | container) for workers and jobs
// channel. It orchestrates workload distribution between all available
// workers.
type WorkerPool struct {
	Workers []*Worker
	Queue   chan Job
}

// NewWorkerPool is a constructor function. It returns fresh instance
// of the WorkerPool struct, which main responsibility is to keep track
// over all resources necessary for future background workload distribution.
func NewWorkerPool(n int) (*WorkerPool, error) {
	workers := make([]*Worker, n)
	worker := new(Worker)
	queue := make(chan Job, 100)

	var id uuid.UUID
	var err error
	var fl *FileLogger

	for i := 0; i < n; i++ {

		worker = NewWorker()

		if id, err = uuid.NewV4(); err != nil {
			return nil, err
		}

		if fl, err = NewFileLogger(fmt.Sprintf("worker_%d.log", i)); err != nil {
			return nil, err
		}

		worker.ID = id
		worker.Queue = queue
		worker.Done = make(chan bool, 1)
		worker.Fl = fl

		workers[i] = worker
	}

	return &WorkerPool{Workers: workers, Queue: queue}, nil
}

// Schedule is a *WorkerPool receiver method, which main responsibility
// is to append new jobs into *WorkerPool.Queue. It's impossible to schedule
// new jobs, at least until queue won't become out of capacity.
func (wp *WorkerPool) Schedule(job Job) bool {
	select {
	case wp.Queue <- job:
		return true
	default:
		return false
	}
}

// Run is a *WorkerPool receiver method, which main responsibility
// is to delegate Run() command to each specific worker.
func (wp *WorkerPool) Run() {
	for _, worker := range wp.Workers {
		worker.Run()
	}
}

// CloseAndWait is a *WorkerPool receiver method, which main responsibility
// is to close *WorkerPool.Queue as well as wait until all previously scheduled
// jobs will be completed.
func (wp *WorkerPool) CloseAndWait() {
	close(wp.Queue)

	for _, worker := range wp.Workers {
		<-worker.Done
	}
}

// Worker is a struct, which main responsibility is to orchestrate
// all resources necessary for running background jobs in parallel.
type Worker struct {
	ID    uuid.UUID
	Queue chan Job
	Done  chan bool
	Fl    *FileLogger
}

// NewWorker is a constructor function. It returns fresh instance
// of the Worker struct.
func NewWorker() *Worker {
	return &Worker{}
}

// FileLogger is a struct, which main responsibility is to represent
// layer of abstraction over logging service where all logs will become
// redirected to specific file instead of standard output.
type FileLogger struct {
	logger *log.Logger
	file   *os.File
}

// NewFileLogger is a constructor function. It returns fresh instance
// of the FileLogger struct and creates a new, separate .log file for each
// of the workers.
func NewFileLogger(fileName string) (*FileLogger, error) {

	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("Failed to create log file for worker. Err: %s", err.Error())
		return nil, err
	}

	l := log.New()
	l.Out = f

	return &FileLogger{logger: l, file: f}, nil
}

// Run is a *Worker receiver method, which main responsibility is to
// fire up Worker's background processing. It spawns a new goroutine to
// make it run independently, in a non-blocking fashion.
func (w *Worker) Run() {
	go func() {
		defer func() { w.Done <- true }()

		w.Fl.logger.WithFields(log.Fields{"WorkerID": w.ID}).Info("Starting worker")

		for job := range w.Queue {
			w.Fl.logger.WithFields(log.Fields{"WorkerID": w.ID, "JobID": job.GetID()}).Info("Job received")
			job.Perform()
			w.Fl.logger.WithFields(log.Fields{"WorkerID": w.ID, "JobID": job.GetID()}).Info("Job done")
		}
	}()
}
