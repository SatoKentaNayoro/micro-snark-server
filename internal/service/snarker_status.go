package service

import (
	v1 "micro-snark-server/api/v1"
	"sync"
	"time"
)

type srvStatus struct {
	Status       v1.ServerStatus
	LastUpdateAt time.Time
	Lock         *sync.RWMutex
}

func NewSrvStatus() *srvStatus {
	return &srvStatus{
		Status:       v1.ServerStatus_FREE,
		Lock:         new(sync.RWMutex),
		LastUpdateAt: time.Now(),
	}
}

func (ss *srvStatus) isFree() bool {
	ss.Lock.RLock()
	defer ss.Lock.RUnlock()
	if ss.Status == v1.ServerStatus_FREE {
		return true
	} else {
		return false
	}
}

func (ss *srvStatus) setStatus(s v1.ServerStatus) {
	ss.Lock.Lock()
	defer ss.Lock.Unlock()
	ss.Status = s
	ss.LastUpdateAt = time.Now()
}

func (ss *srvStatus) getStatus() v1.ServerStatus {
	ss.Lock.RLock()
	defer ss.Lock.RUnlock()
	return ss.Status
}

// check is the server locked by one miner,but not used for long time
func (ss *srvStatus) lockExpired(expireTime time.Duration) bool {
	ss.Lock.RLock()
	defer ss.Lock.RUnlock()
	if time.Since(ss.LastUpdateAt) >= expireTime && ss.Status == v1.ServerStatus_LOCKED {
		return true
	}
	return false
}
