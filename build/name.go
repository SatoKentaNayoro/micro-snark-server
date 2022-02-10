package build

import "os"

type AppName struct {
	Name string
	ID   string
}

func GetAppName() AppName {
	hostname, _ := os.Hostname()
	return AppName{
		Name: "micro-snark-server",
		ID:   hostname,
	}
}
