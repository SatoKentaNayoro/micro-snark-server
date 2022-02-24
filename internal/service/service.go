package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "micro-snark-server/api/v1"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewSnarkerService)

// SnarkerService is a snark task service
type SnarkerService struct {
	v1.UnimplementedSnarkTaskServer

	Status *srvStatus

	log *log.Helper
}

func NewSnarkerService(logger log.Logger) *SnarkerService {
	return &SnarkerService{Status: NewSrvStatus(), log: log.NewHelper(logger)}
}
