package service

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	v1 "micro-snark-server/api/v1"
)

var (
	ErrorCanNotUsedWithStatus = func(srvId string, status v1.ServerStatus) error {
		return errors.New(fmt.Sprintf("server: %s is cannot be used with status: %s", srvId, status.String()))
	}
)

func (s *SnarkerService) DoSnarkTask(ctx context.Context, in *v1.DoSnarkTaskRequest) (*emptypb.Empty, error) {
	// todo
	return nil, nil
}

func (s *SnarkerService) GetOneFreeServer(ctx context.Context, in *v1.GetOneFreeServerRequest) (*v1.GetOneFreeServerResponse, error) {
	s.correctServerStatus()
	if st := s.Status.getStatus(); st != v1.ServerStatus_FREE {
		err := v1.ErrorReason_SRV_CAN_NOT_USED_NOW
		return &v1.GetOneFreeServerResponse{Ok: false, ErrorReason: &err}, nil
	}
	s.Status.setStatus(v1.ServerStatus_LOCKED)
	s.Task.setId(in.TaskId)
	s.Task.setStatus(v1.TaskStatus_Waiting)
	return &v1.GetOneFreeServerResponse{Ok: true}, nil
}

func (s *SnarkerService) GetServerStatus(ctx context.Context, in *v1.ServerStatusRequest) (*v1.ServerStatusResponse, error) {
	// todo
	return nil, nil
}

func (s *SnarkerService) GetTaskResult(ctx context.Context, in *v1.GetTaskResultRequest) (*v1.GetTaskResultResponse, error) {
	// todo
	return nil, nil
}

func (s *SnarkerService) correctServerStatus() {
	if s.Status.lockExpired(*s.Options.MaxSrvLockedTime) {
		s.Log.Warnf("server locked by task: %s expired,will unlock this server", s.Task.TaskId)
		s.resetServerStatusAndTask()
	}
	if s.Task.resultExpired(*s.Options.MaxResRetTime) {
		s.Log.Warnf("task result of task_id: %s expired,will be drop", s.Task.TaskId)
		s.resetServerStatusAndTask()
	}
}

func (s *SnarkerService) resetServerStatusAndTask() {
	s.Status.setStatus(v1.ServerStatus_FREE)
	s.Task.clean()
}
