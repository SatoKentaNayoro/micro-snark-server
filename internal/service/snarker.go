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
	ErrorTaskIdError = func(localID, givenID string) error {
		return errors.New(fmt.Sprintf("task id error,locak id: %s but %s", localID, givenID))
	}
	ErrorTaskWrongStatus = errors.New("task on wrong status,this should never happen")
)

func (s *SnarkerService) DoSnarkTask(ctx context.Context, in *v1.DoSnarkTaskRequest) (*v1.DoSnarkTaskResponse, error) {
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
	s.Task.SetId(in.TaskId)
	s.Task.SetStatus(v1.TaskStatus_Waiting)
	return &v1.GetOneFreeServerResponse{Ok: true}, nil
}

func (s *SnarkerService) GetServerStatus(ctx context.Context, in *emptypb.Empty) (*v1.ServerStatusResponse, error) {
	return &v1.ServerStatusResponse{ServerStatus: s.Status.getStatus()}, nil
}

func (s *SnarkerService) GetTaskResult(ctx context.Context, in *v1.GetTaskResultRequest) (*v1.GetTaskResultResponse, error) {
	id, same := s.checkTaskId(in.TaskId)
	if !same {
		return nil, ErrorTaskIdError(id, in.TaskId)
	}
	var err v1.ErrorReason
	switch s.Task.GetStatus() {
	case v1.TaskStatus_Running:
		err = v1.ErrorReason_TASK_NOT_DONE
		return &v1.GetTaskResultResponse{Ok: false, ErrorReason: &err, Result: nil}, nil
	case v1.TaskStatus_Failed:
		err = v1.ErrorReason_TASK_FAILED
		msg := s.Task.GetFailedReason().Error()
		return &v1.GetTaskResultResponse{Ok: false, ErrorReason: &err, ErrorMsg: &msg, Result: nil}, nil
	case v1.TaskStatus_Done:
		return &v1.GetTaskResultResponse{Ok: true, Result: s.Task.GetResult()}, nil
	default:
		return nil, ErrorTaskWrongStatus
	}
}

func (s *SnarkerService) correctServerStatus() {
	if s.Status.lockExpired(*s.Options.MaxSrvLockedTime) {
		s.Log.Warnf("server locked by task: %s expired,will unlock this server", s.Task.TaskId)
		s.resetServerStatusAndTask()
	}
	if s.Task.ResultExpired(*s.Options.MaxResRetTime) {
		s.Log.Warnf("task result of task_id: %s expired,will be drop", s.Task.TaskId)
		s.resetServerStatusAndTask()
	}
}

func (s *SnarkerService) resetServerStatusAndTask() {
	s.Status.setStatus(v1.ServerStatus_FREE)
	s.Task.Clean()
}

func (s *SnarkerService) checkTaskId(givenId string) (string, bool) {
	if id := s.Task.GetId(); id != givenId {
		return id, false
	}
	return givenId, true
}
