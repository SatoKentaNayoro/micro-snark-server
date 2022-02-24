package service

import (
	"encoding/json"
	v1 "micro-snark-server/api/v1"
	"sync"
	"time"
)

type task struct {
	TaskId       string
	VanillaProof []byte
	PubIn        []byte
	PostConfig   []byte
	ReplicasLen  uint64
	Result       []byte
	TaskStatus   v1.TaskStatus
	Lock         *sync.RWMutex
	LastUpdateAt time.Time
}

func NewTask() *task {
	return &task{
		TaskStatus:   v1.TaskStatus_None,
		Lock:         new(sync.RWMutex),
		LastUpdateAt: time.Now(),
	}
}

func (t *task) fromBytes(raw []byte) error {
	t.Lock.Lock()
	defer t.Lock.Unlock()
	if err := json.Unmarshal(raw, t); err != nil {
		return err
	}
	return nil
}

func (t *task) fromReq(raw *v1.DoSnarkTaskRequest) {
	t.Lock.Lock()
	defer t.Lock.Unlock()
	t.TaskId = raw.TaskId
	t.VanillaProof = raw.VanillaProof
	t.PubIn = raw.PubIn
	t.PostConfig = raw.PostConfig
	t.ReplicasLen = raw.ReplicasLen
}

func (t *task) setStatus(s v1.TaskStatus) {
	t.Lock.Lock()
	defer t.Lock.Unlock()
	t.TaskStatus = s
	t.LastUpdateAt = time.Now()
}

func (t *task) getStatus() v1.TaskStatus {
	t.Lock.RLock()
	defer t.Lock.RUnlock()
	return t.TaskStatus
}

func (t *task) setId(id string) {
	t.Lock.Lock()
	defer t.Lock.Unlock()
	t.TaskId = id
}

func (t *task) clean() {
	t.Lock.Lock()
	defer t.Lock.Unlock()
	t.TaskId,
		t.VanillaProof,
		t.PubIn,
		t.PostConfig,
		t.ReplicasLen,
		t.Result,
		t.TaskStatus = "", nil, nil, nil, 0, nil, v1.TaskStatus_None
}

func (t *task) resultExpired(expireTime time.Duration) bool {
	if time.Since(t.LastUpdateAt) >= expireTime && t.TaskStatus == v1.TaskStatus_Done {
		return true
	}
	return false
}
