package data

import v1 "micro-snark-server/api/v1"

type TaskRepo struct {
	TaskId       string
	VanillaProof []byte
	PubIn        []byte
	PostConfig   []byte
	ReplicasLen  int64
	Result       []byte
	TaskStatus   v1.TaskStatus
}
