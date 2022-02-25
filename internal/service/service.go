package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "micro-snark-server/api/v1"
	"micro-snark-server/internal/conf"
	"micro-snark-server/internal/task"
	"time"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewSnarkerService)

// SnarkerService is a snark task service
type SnarkerService struct {
	v1.UnimplementedSnarkTaskServer
	SrvID   string
	Status  *srvStatus
	Task    *task.Task
	Worker  *worker
	Log     *log.Helper
	Options struct {
		// If a server is locked,it will be unlocked automatically within the MaxSrvLockedTime limit
		// MaxSrvLockedTime will be parsed to second,default 10 seconds
		MaxSrvLockedTime *time.Duration
		// if miner do not get result back right now when task done,the result will be dropped after MaxResRetTime.
		// and within the MaxResRetTime,the server cannot be used by other miner,default 60 seconds
		MaxResRetTime *time.Duration
		// When server received an exit signal,if miner hasn't got result back,server will wait MaxWaitExitTime.
		MaxWaitExitTime *time.Duration
	}
}

func NewSnarkerService(conf conf.Server, srvId string, logger log.Logger) *SnarkerService {
	return &SnarkerService{SrvID: srvId, Status: NewSrvStatus(), Task: task.NewTask(), Worker: NewWorker(), Log: log.NewHelper(logger), Options: struct {
		MaxSrvLockedTime *time.Duration
		MaxResRetTime    *time.Duration
		MaxWaitExitTime  *time.Duration
	}{MaxSrvLockedTime: conf.Options.MaxSrvLockedTime, MaxResRetTime: conf.Options.MaxResRetTime, MaxWaitExitTime: conf.Options.MaxWaitExitTime}}
}
