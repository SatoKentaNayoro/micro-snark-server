package task

import (
	"encoding/json"
	v1 "micro-snark-server/api/v1"
	"sync"
	"time"
)

type Task struct {
	TaskId       string
	VanillaProof []byte
	PubIn        []byte
	PostConfig   []byte
	ReplicasLen  uint64
	Result       []byte
	TaskStatus   v1.TaskStatus
	FailedReason error
	Lock         *sync.RWMutex
	LastUpdateAt time.Time
}

func NewTask() *Task {
	return &Task{
		TaskStatus:   v1.TaskStatus_None,
		Lock:         new(sync.RWMutex),
		LastUpdateAt: time.Now(),
	}
}

func (t *Task) FromBytes(raw []byte) error {
	t.Lock.Lock()
	defer t.Lock.Unlock()
	if err := json.Unmarshal(raw, t); err != nil {
		return err
	}
	return nil
}

func (t *Task) FromReq(raw *v1.DoSnarkTaskRequest) {
	t.Lock.Lock()
	defer t.Lock.Unlock()
	t.TaskId = raw.TaskId
	t.VanillaProof = raw.VanillaProof
	t.PubIn = raw.PubIn
	t.PostConfig = raw.PostConfig
	t.ReplicasLen = raw.ReplicasLen
}

func (t *Task) SetStatus(s v1.TaskStatus) {
	t.Lock.Lock()
	defer t.Lock.Unlock()
	t.TaskStatus = s
	t.LastUpdateAt = time.Now()
}

func (t *Task) GetStatus() v1.TaskStatus {
	t.Lock.RLock()
	defer t.Lock.RUnlock()
	return t.TaskStatus
}

func (t *Task) GetFailedReason() error {
	t.Lock.RLock()
	defer t.Lock.RUnlock()
	return t.FailedReason
}

func (t *Task) GetResult() []byte {
	t.Lock.RLock()
	defer t.Lock.RUnlock()
	return t.Result
}

func (t *Task) SetId(id string) {
	t.Lock.Lock()
	defer t.Lock.Unlock()
	t.TaskId = id
}

func (t *Task) GetId() string {
	t.Lock.Lock()
	defer t.Lock.Unlock()
	return t.TaskId
}

func (t *Task) Clean() {
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

func (t *Task) SetFailedReason(err error) {
	t.Lock.Lock()
	defer t.Lock.Unlock()
	t.FailedReason = err
}

func (t *Task) ResultExpired(expireTime time.Duration) bool {
	if time.Since(t.LastUpdateAt) >= expireTime && t.TaskStatus == v1.TaskStatus_Done {
		return true
	}
	return false
}
