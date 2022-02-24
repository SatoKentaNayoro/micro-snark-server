package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "micro-snark-server/api/v1"
)

// SnarkerService is a snark task service
type SnarkerService struct {
	v1.UnimplementedSnarkTaskServer

	Status v1.ServerStatus

	log *log.Helper
}

func NewSnarkerService(logger log.Logger) *SnarkerService {
	return &SnarkerService{Status: v1.ServerStatus_SERVER_FREE, log: log.NewHelper(logger)}
}

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
