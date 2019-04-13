package service

import (
	"log"
	"time"
)

type Task struct {
	ticker *time.Ticker
	do     func() error
}

func (t *Task) Run() {
	for {
		select {
		case <-t.ticker.C:
			err := t.do()
			if err != nil {
				log.Panic("failed to execute task", err)
			}
		}
	}
}

func NewTask(do func() error, duration time.Duration) *Task {
	task := new(Task)
	task.do = do
	task.ticker = time.NewTicker(duration)
	return task
}
