package service

import (
	v1 "micro-snark-server/api/v1"
	"sync"
)

type srvStatus struct {
	Status v1.ServerStatus
	Lock   *sync.RWMutex
}

func NewSrvStatus() *srvStatus {
	return &srvStatus{
		Status: v1.ServerStatus_SERVER_FREE,
		Lock:   new(sync.RWMutex),
	}
}

func (ss *srvStatus) isFree() bool {
	ss.Lock.RLock()
	defer ss.Lock.RUnlock()
	if ss.Status == v1.ServerStatus_SERVER_FREE {
		return true
	} else {
		return false
	}
}

func (ss *srvStatus) setStatus(s v1.ServerStatus) {
	ss.Lock.Lock()
	defer ss.Lock.Unlock()
	ss.Status = s
}

func (ss *srvStatus) getStatus() v1.ServerStatus {
	ss.Lock.RLock()
	defer ss.Lock.RUnlock()
	return ss.Status
}
