package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "micro-snark-server/api/v1"
)

func (s *SnarkerService) DoSnarkTask(ctx context.Context, in *v1.DoSnarkTaskRequest) (*v1.BaseResponse, error) {
	// todo
	return nil, nil
}

func (s *SnarkerService) GetOneFreeServer(ctx context.Context, in *v1.GetOneFreeServerRequest) (*v1.BaseResponse, error) {
	// todo
	return nil, nil
}

func (s *SnarkerService) GetServerStatus(ctx context.Context, in *v1.ServerStatusRequest) (*v1.ServerStatusResponse, error) {
	// todo
	return nil, nil
}

func (s *SnarkerService) GetTaskResult(ctx context.Context, in *v1.GetTaskResultRequest) (*v1.GetTaskResultResponse, error) {
	// todo
	return nil, nil
}
