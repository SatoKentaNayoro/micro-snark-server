package conf

import "time"

type Server struct {
	Http struct {
		// Network default tcp
		Network string
		// Addr likes 0.0.0.0:8000
		Addr string
		// Default 1s
		Timeout *time.Duration
	}

	Grpc struct {
		// Network default tcp
		Network string
		// Addr likes 0.0.0.0:9000
		Addr string
		// Default 1s
		Timeout *time.Duration
	}

	// If a server is locked,it will be unlocked automatically within the MaxSrvLockedTime limit
	// MaxSrvLockedTime will be parsed to second,default 10 seconds
	MaxSrvLockedTime int
	// if miner do not get result back right now when task done,the result will be dropped after MaxResRetTime.
	// and within the MaxResRetTime,the server cannot be used by other miner,default 60 seconds
	MaxResRetTime int
	// When server received an exit signal,if miner hasn't got result back,server will wait MaxWaitExitTime.
	MaxWaitExitTime int
}
