package service

import (
	"micro-snark-server/internal/snark-ffi"
	"micro-snark-server/internal/task"
)

type worker struct {
	TaskChan chan *task.Task
	StopChan chan struct{}
}

func (w *worker) run() {
	for {
		select {
		case t := <-w.TaskChan:
			snark_ffi.SnarkPost(t)

		case <-w.StopChan:
			break
		}
	}
}
