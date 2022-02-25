package service

import (
	v1 "micro-snark-server/api/v1"
	"micro-snark-server/internal/snark-ffi"
	"micro-snark-server/internal/task"
)

type worker struct {
	TaskChan chan *task.Task
	StopChan chan struct{}
}

func NewWorker() *worker {
	return &worker{
		TaskChan: make(chan *task.Task, 1),
		StopChan: make(chan struct{}, 1),
	}
}

func (w *worker) run() {
	for {
		select {
		case t := <-w.TaskChan:
			res, err := snark_ffi.SnarkPost(t)
			if err != nil {
				t.SetFailedReason(err)
				t.SetStatus(v1.TaskStatus_Failed)
			} else {
				t.SetResult(res)
				t.SetStatus(v1.TaskStatus_Done)
			}
		case <-w.StopChan:
			break
		}
	}
}
